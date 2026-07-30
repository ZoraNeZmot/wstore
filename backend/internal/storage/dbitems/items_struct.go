package dbitems

import (
	"time"

	"github.com/google/uuid"
)

type ItemStatus string

const (
	ItemStatusDraft    ItemStatus = "draft"
	ItemStatusActive   ItemStatus = "active"
	ItemStatusInactive ItemStatus = "inactive"
	ItemStatusDeleted  ItemStatus = "deleted"
	ItemStatusSold     ItemStatus = "sold"
)

const DefaultAmount = 10

func (i *Item) validateAmount() {
	if i.amount <= 0 {
		i.amount = DefaultAmount
	}
}

type Item struct {
	ID          uuid.UUID  `json:"id,omitempty"`
	SellerID    uuid.UUID  `json:"seller_id,omitempty"`
	Title       string     `json:"title,omitempty"`
	Description string     `json:"description,omitempty"`
	Price       float64    `json:"price,omitempty"`
	Quantity    int        `json:"quantity,omitempty"`
	Status      ItemStatus `json:"status,omitempty"`
	CreatedAt   time.Time  `json:"created_at,omitzero"`
	UpdatedAt   time.Time  `json:"updated_at,omitzero"`
	// Functinal fields
	amount int
}

func NewItem(sellerID uuid.UUID, title string, description string, price float64, quantity int, status ItemStatus) *Item {
	return &Item{
		SellerID:    sellerID,
		Title:       title,
		Description: description,
		Price:       price,
		Quantity:    quantity,
		Status:      status,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func NewWithAmount(amount int) *Item {
	return &Item{amount: amount}
}

func NewWithSellerIDAndAmount(sellerID uuid.UUID, amount int) *Item {
	return &Item{SellerID: sellerID, amount: amount}
}

func NewWithSellerID(sellerID uuid.UUID) *Item {
	return &Item{SellerID: sellerID}
}

func NewWithID(id uuid.UUID) *Item {
	return &Item{ID: id}
}

func (i *Item) GetSQLQuery(procName string) (string, []any) {
	i.validateAmount()
	switch procName {
	case "GetItemsBySeller":
		return GetItemsBySeller, []any{i.SellerID, i.amount}
	case "GetItems":
		return GetItems, []any{i.amount}
	case "GetItem":
		return GetItem, []any{i.ID}
	default:
		return "", []any{}
	}
}
