package storage

import (
	"context"
	"wstore/internal/storage/dbitems"
)

func (d *Database) GetItems(ctx context.Context, item *dbitems.Item) ([]*dbitems.Item, error) {
	items := make([]*dbitems.Item, 0)
	rows, err := d.Query(ctx, "GetItems", item)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		item := &dbitems.Item{}
		err := rows.Scan(&item.ID, &item.SellerID, &item.Title, &item.Price, &item.Status)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (d *Database) GetItemsBySellerID(ctx context.Context, item *dbitems.Item) ([]*dbitems.Item, error) {
	items := make([]*dbitems.Item, 0)
	sellerID := item.SellerID
	rows, err := d.Query(ctx, "GetItemsBySeller", item)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		item := &dbitems.Item{}
		err := rows.Scan(&item.ID, &item.Title, &item.Price, &item.Status)
		if err != nil {
			return nil, err
		}
		item.SellerID = sellerID
		items = append(items, item)
	}
	return items, nil
}

func (d *Database) GetItem(ctx context.Context, item *dbitems.Item) (*dbitems.Item, error) {
	row := d.QueryRow(ctx, "GetItem", item)
	err := row.Scan(&item.SellerID, &item.Title, &item.Description, &item.Price, &item.Quantity, &item.Status)
	if err != nil {
		return nil, err
	}
	return item, nil
}
