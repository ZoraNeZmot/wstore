package dbitems

import (
	"time"

	"github.com/google/uuid"
)

type ItemStatus string
type ItemOpt func(*Item)

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
	query  string
}

func WithAmount(amount int) ItemOpt {
	return func(i *Item) {
		i.amount = amount
	}
}

func WithQuery(query string) ItemOpt {
	return func(i *Item) {
		i.query = query
	}
}

func WithSellerID(sellerID uuid.UUID) ItemOpt {
	return func(i *Item) {
		i.SellerID = sellerID
	}
}

func WithID(id uuid.UUID) ItemOpt {
	return func(i *Item) {
		i.ID = id
	}
}

func WithDescription(description string) ItemOpt {
	return func(i *Item) {
		i.Description = description
	}
}

func WithTitle(title string) ItemOpt {
	return func(i *Item) {
		i.Title = title
	}
}

func WithPrice(price float64) ItemOpt {
	return func(i *Item) {
		i.Price = price
	}
}

func WithQuantity(quantity int) ItemOpt {
	return func(i *Item) {
		i.Quantity = quantity
	}
}

func WithStatus(status ItemStatus) ItemOpt {
	return func(i *Item) {
		i.Status = status
	}
}

func NewItem(opts ...ItemOpt) *Item {
	item := &Item{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	for _, opt := range opts {
		opt(item)
	}
	return item
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
	case "GetItemsBySearch":
		return GetItemsBySearch, []any{i.query, i.amount}
	case "CheckItemOwnership":
		return CheckItemOwnership, []any{i.ID, i.SellerID}
	default:
		return "", []any{}
	}
}
