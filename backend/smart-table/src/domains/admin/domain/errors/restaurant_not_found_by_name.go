package errors

import (
	"fmt"
)

type RestaurantNotFoundByName struct {
	Name string
}

func (e RestaurantNotFoundByName) Error() string {
	return fmt.Sprintf("Restaurant not found by name: %s", e.Name)
}
