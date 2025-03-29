package errors

import (
	"fmt"
)

type RestaurantNameExists struct {
	Name string
}

func (e RestaurantNameExists) Error() string {
	return fmt.Sprintf("restaurant with name '%s' already exists", e.Name)
}
