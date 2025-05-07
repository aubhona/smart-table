package errors

import (
	"fmt"
)

type RestaurantNameAlreadyExists struct {
	Name string
}

func (e RestaurantNameAlreadyExists) Error() string {
	return fmt.Sprintf("restaurant with name '%s' already exists", e.Name)
}
