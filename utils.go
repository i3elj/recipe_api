package main

import (
	"golang.org/x/exp/constraints"
)

// If certain ingredient is found in a list of ingredients, true is returned. False otherwise
func has_ing(ings []Ingredient, str string) bool {
	for _, v := range ings {
		if v.What == str {
			return true
		}
	}
	return false
}

func find[T constraints.Ordered](arr []T, value T) bool {
	for i := 0; i < len(arr); i++ {
		if value == arr[i] {
			return true
		}
	}
	return false
}
