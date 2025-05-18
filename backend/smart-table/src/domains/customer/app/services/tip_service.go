package app

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/smart-table/src/domains/customer/domain"
	"github.com/smart-table/src/utils"
)

type ItemTipInfo struct {
	Count       int
	DishUUID    uuid.UUID
	Name        string
	Price       decimal.Decimal
	ResultPrice decimal.Decimal
}

type CustomerTipInfo struct {
	UUID           uuid.UUID
	TgLogin        string
	TgID           string
	TotalPrice     decimal.Decimal
	ItemTipInfoMap map[string]ItemTipInfo
}

type TipService struct {
}

func NewTipService() *TipService {
	return &TipService{}
}

func (t *TipService) GetItemsGroupKey(item utils.SharedRef[domain.Item]) string {
	return fmt.Sprintf(
		"{%s}_{%s}_{%s}",
		item.Get().GetDishUUID(),
		item.Get().GetStatus(),
		item.Get().GetComment().ValueOr(""),
	)
}

func (t *TipService) groupCustomerItems(order utils.SharedRef[domain.Order]) map[uuid.UUID]CustomerTipInfo {
	customerTipInfoMap := make(map[uuid.UUID]CustomerTipInfo)

	for _, customer := range order.Get().GetCustomers() {
		customerTipInfoMap[customer.Get().GetUUID()] = CustomerTipInfo{
			UUID:           customer.Get().GetUUID(),
			TgLogin:        customer.Get().GetTgLogin(),
			TgID:           customer.Get().GetTgID(),
			TotalPrice:     decimal.Zero,
			ItemTipInfoMap: make(map[string]ItemTipInfo),
		}
	}

	for _, item := range order.Get().GetItems() {
		if item.Get().GetIsDraft() {
			continue
		}

		key := t.GetItemsGroupKey(item)

		customerUUID := item.Get().GetCustomer().Get().GetUUID()
		customerTipInfo := customerTipInfoMap[customerUUID]

		customerTipInfo.TotalPrice = customerTipInfo.TotalPrice.Add(item.Get().GetPrice())

		itemTipInfo, isExist := customerTipInfo.ItemTipInfoMap[key]

		if !isExist {
			customerTipInfo.ItemTipInfoMap[key] = ItemTipInfo{
				Count:       1,
				DishUUID:    item.Get().GetDishUUID(),
				Name:        item.Get().GetName(),
				Price:       item.Get().GetPrice(),
				ResultPrice: item.Get().GetPrice(),
			}
		} else {
			itemTipInfo.Count++
			itemTipInfo.ResultPrice = itemTipInfo.ResultPrice.Add(itemTipInfo.Price)

			customerTipInfo.ItemTipInfoMap[key] = itemTipInfo
		}

		customerTipInfoMap[customerUUID] = customerTipInfo
	}

	return customerTipInfoMap
}

func (t *TipService) CreateStringTip(order utils.SharedRef[domain.Order]) string {
	customerTipInfoMap := t.groupCustomerItems(order)

	var result strings.Builder

	result.WriteString("🧾 *Чек по заказу*\n\n")

	for i, customer := range order.Get().GetCustomers() {
		info := customerTipInfoMap[customer.Get().GetUUID()]
		result.WriteString(fmt.Sprintf("*Гость:* @%s\n", utils.EscapeMarkdown(info.TgLogin)))
		result.WriteString(fmt.Sprintf("*Сумма:* `%s`\n", info.TotalPrice.StringFixed(2)))

		if len(info.ItemTipInfoMap) > 0 {
			result.WriteString("_Блюда:_\n")

			for _, item := range info.ItemTipInfoMap {
				result.WriteString(fmt.Sprintf("%s %s×%d = %s\n",
					utils.EscapeMarkdown(item.Name),
					item.Price.StringFixed(2),
					item.Count,
					item.ResultPrice.StringFixed(2),
				))
			}
		}

		if i != len(order.Get().GetCustomers())-1 {
			result.WriteString("\n---\n")
		}
	}

	return result.String()
}
