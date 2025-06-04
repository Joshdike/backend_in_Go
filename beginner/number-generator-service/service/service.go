package service

import (
	"errors"
	"math/rand"
)

func RandomNumbers(min, max, quantity int) ([]int, error) {
	switch {
	case min < 0:
		return []int{}, errors.New("min cannot be negative")
	case max <= 0:
		return []int{}, errors.New("max must be greater than zero")
	case min > max:
		return []int{}, errors.New("min cannot be greater than max")
	case quantity < 1:
		return []int{}, errors.New("quantity must be atleast 1")
	}

	numbers := make([]int, quantity)
	for i := 0; i < quantity; i++ {
		numbers[i] = rand.Intn(max-min+1) + min
	}
	return numbers, nil

}
