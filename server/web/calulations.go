package web

import (
	"math/rand"
	"tpo2/model"
)

const (
	FAIR   = "Честный"
	RANDOM = "Рандомный"
	EQUAL  = "Поровну"
)

func calculateResult(order model.Order) map[string]float64 {
	result := make(map[string]float64)

	switch order.Mode {
	case RANDOM:
		randUserNum := rand.Intn(len(order.Users))
		username := order.Users[randUserNum]
		result[username] = calculateDishSum(order.Dishes)
	case EQUAL:
		val := calculateDishSum(order.Dishes) / float64(len(order.Users))
		for _, user := range order.Users {
			result[user] = val
		}
	default:
		for _, dish := range order.Dishes {
			val := dish.Price / float64(len(dish.Users))
			for _, user := range dish.Users {
				if _, ok := result[user]; ok {
					result[user] += val
					continue
				}
				result[user] = val
			}
		}
	}

	return result
}

func calculateDishSum(dishes []model.Dish) float64 {
	sum := 0.0
	for _, dish := range dishes {
		sum += dish.Price
	}
	return sum
}
