package web

import (
	"fmt"
	"math/rand"
	"tpo2/model"
)

const (
	FAIR   = "Честный"
	RANDOM = "Рандомный"
	EQUAL  = "Поровну"
)

type result struct {
	Name string `json:"name"`
	Sum  string `json:"sum"`
}

func calculateResult(order model.Order) []result {
	resultCalc := make(map[string]float64)

	switch order.Mode {
	case RANDOM:
		randUserNum := rand.Intn(len(order.Users))
		username := order.Users[randUserNum]
		resultCalc[username] = float64(calculateDishSum(order.Dishes))
	case EQUAL:
		val := calculateDishSum(order.Dishes) / float64(len(order.Users))
		for _, user := range order.Users {
			resultCalc[user] = float64(val)
		}
	default:
		for _, dish := range order.Dishes {
			val := dish.Price / float64(len(dish.Users))
			for _, user := range dish.Users {
				if _, ok := resultCalc[user]; ok {
					resultCalc[user] += float64(val)
					continue
				}
				resultCalc[user] = float64(val)
			}
		}
	}
	resultArr := make([]result, len(resultCalc))
	i := 0
	for name, sum := range resultCalc {
		resultArr[i] = result{
			Name: name,
			Sum:  fmt.Sprintf("%.2f", float64(sum)),
		}
		i++
	}
	return resultArr
}

func calculateDishSum(dishes []model.Dish) float64 {
	sum := 0.0
	for _, dish := range dishes {
		sum += dish.Price
	}
	return sum
}
