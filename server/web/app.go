package web

import (
	"encoding/json"
	"errors"
	"flag"
	"log"
	"net/http"
	"tpo2/model"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type App struct {
	handlers map[string]http.HandlerFunc
	cache    map[string]model.Order
}

func NewApp(cors bool) App {
	app := App{
		handlers: make(map[string]http.HandlerFunc),
		cache:    make(map[string]model.Order),
	}

	app.handlers["/"] = http.FileServer(http.Dir("/webapp")).ServeHTTP
	return app
}

func (a *App) Serve() error {
	router := mux.NewRouter().StrictSlash(true)

	var dir string
	flag.StringVar(&dir, "dir", "./resources/static", "the directory to serve static files")
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))
	apiSubrouter := router.PathPrefix("/api").Subrouter()

	apiSubrouter.
		Path("/order").
		Methods("POST").
		Handler(makeHandler(a, StartOrder))

	apiSubrouter.
		Path("/order/{id}").
		Methods("GET").
		Handler(makeHandler(a, GetOrder))

	apiSubrouter.
		Path("/order/{id}/addDish").
		Methods("POST").
		Handler(makeHandler(a, AddDish))

	apiSubrouter.
		Path("/order/{id}/finishOrder").
		Methods("POST").
		Handler(makeHandler(a, FinishOrder))

	http.Handle("/", router)

	var handler http.Handler
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})
	exposedHeaders := handlers.ExposedHeaders([]string{"Location"})

	handler = handlers.CORS(originsOk, headersOk, methodsOk, exposedHeaders)(router)

	log.Println("Web server is available on port 8080")
	return http.ListenAndServe(":8080", handler)
}

func StartOrder(w http.ResponseWriter, r *http.Request, a *App) {
	w.Header().Set("Content-Type", "application/json")
	var order model.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id := String(6)
	a.cache[id] = order
	w.Write([]byte(id))
}

func AddDish(w http.ResponseWriter, r *http.Request, a *App) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	if _, ok := a.cache[id]; !ok {
		sendErr(w, http.StatusBadRequest, errors.New("not found").Error())
		return
	}
	check := a.cache[id]
	var dish model.Dish
	err := json.NewDecoder(r.Body).Decode(&dish)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	check.Dishes = append(check.Dishes, dish)
	a.cache[id] = check
	jsonCheck, err := json.Marshal(check)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.Write(jsonCheck)
}

func GetOrder(w http.ResponseWriter, r *http.Request, a *App) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	if _, ok := a.cache[id]; !ok {
		sendErr(w, http.StatusBadRequest, errors.New("not found").Error())
		return
	}
	check := a.cache[id]
	jsonCheck, err := json.Marshal(check)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.Write(jsonCheck)
}

func FinishOrder(w http.ResponseWriter, r *http.Request, a *App) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	if _, ok := a.cache[id]; !ok {
		sendErr(w, http.StatusBadRequest, errors.New("not found").Error())
		return
	}
	var isFinished bool
	err := json.NewDecoder(r.Body).Decode(&isFinished)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	order := a.cache[id]
	order.Finished = isFinished
	a.cache[id] = order

	calculations := calculateResult(order)
	results := struct {
		Currency string   `json:"currency"`
		Results  []result `json:"results"`
	}{
		Currency: order.Currency,
		Results:  calculations,
	}

	jsonResults, err := json.Marshal(results)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.Write(jsonResults)
}

func sendErr(w http.ResponseWriter, code int, message string) {
	resp, _ := json.Marshal(map[string]string{"error": message})
	http.Error(w, string(resp), code)
}

// Needed in order to disable CORS for local development
func disableCors(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		h(w, r)
	}
}
