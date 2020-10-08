package model

type Order struct {
	Currency string `json:"currency"`
	Mode     string `json:"mode"`
	Users    []string `json:"users"`
	Dishes   []Dish `json:"dishes"`
	Finished bool `json:"finished"`
}
