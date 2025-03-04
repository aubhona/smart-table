package app

import (
	"github.com/google/uuid"
)

type UserSingUpCommandHandlerResult struct {
	OrderUUID uuid.UUID
}

type UserSingUpCommandHandler struct {
	// orderRepository    domain.UserRepository
	// customerRepository domain.CustomerRepository
	// uuidGenerator      *domain_services.UUIDGenerator
	// roomCodeService    *app.RoomCodeService
}

// func NewOrderCreateCommandHandler(
// 	orderRepository domain.OrderRepository,
// 	customerRepository domain.CustomerRepository,
// 	uuidGenerator *domain_services.UUIDGenerator,
// 	roomCodeService *app.RoomCodeService,
// ) *OrderCreateCommandHandler {
// 	return &OrderCreateCommandHandler{
// 		orderRepository,
// 		customerRepository,
// 		uuidGenerator,
// 		roomCodeService,
// 	}
// }
