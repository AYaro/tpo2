package model

type Dish struct {
	Name  string   `json:"name"`
	Price float64  `json:"price"`
	Users []string `json:"users"`
}
