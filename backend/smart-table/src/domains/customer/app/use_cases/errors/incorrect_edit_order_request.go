package apperrors

type IncorrectEditOrderRequest struct{}

func (e IncorrectEditOrderRequest) Error() string {
	return "incorrect edit order request, order_status, item_status and item_uuid_group are empty"
}
