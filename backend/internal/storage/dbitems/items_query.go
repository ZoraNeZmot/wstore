package dbitems

const (
	GetItemsBySeller = "SELECT * FROM sales.usf_get_items_by_seller($1, $2)"
	GetItems         = "SELECT * FROM sales.usf_get_items($1)"
	GetItem          = "SELECT * FROM sales.usf_get_item($1)"
)
