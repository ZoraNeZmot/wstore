package saleshandlers

import (
	"net/http"
	"wstore/internal/config"
	handlershelp "wstore/internal/handlers/help"
	"wstore/internal/handlers/help/response"
	"wstore/internal/storage"
	"wstore/internal/storage/dbitems"

	"github.com/google/uuid"

	"github.com/sirupsen/logrus"
)

type SalesHandlers struct {
	db     *storage.Database
	log    *logrus.Entry
	config *config.Config
}

func NewSalesHandlers() *SalesHandlers {
	config := config.GetConfig()
	return &SalesHandlers{
		db:     storage.NewDatabase(config.DatabaseConfig.GetDB()),
		log:    logrus.WithField("Handler Type", "Sales"),
		config: config,
	}
}

func (s *SalesHandlers) RegisterRoutes(mux *http.ServeMux) error {
	mux.HandleFunc("GET /sales/items", s.GetItems)
	mux.HandleFunc("GET /sales/items/{id}", s.GetItem)
	return nil
}

// @Summary		Get items
// @Description	Get items
// @Tags			sales
// @Produce		json
// @Param			amount		query		int						false	"Amount of items to return"	default(10)
// @Param			seller_id	query		string					false	"Seller ID"
// @Success		200			{object}	response.ItemsListDTO	"Items list"
// @Failure		400			{object}	response.ErrorDTO		"Bad Request"
// @Failure		500			{object}	response.ErrorDTO		"Server Error"
// @Router			/sales/items [get]
func (s *SalesHandlers) GetItems(w http.ResponseWriter, r *http.Request) {
	amount, err := handlershelp.GetIntFromQuery(r, "amount", dbitems.DefaultAmount)
	if err != nil {
		s.log.WithError(err).Error("Failed to get amount from query")
		response.Error(w, http.StatusBadRequest, "Bad Request", "Failed to get amount from query")
		return
	}
	sellerID, err := handlershelp.GetUUIDFromQuery(r, "seller_id")
	if !handlershelp.IfEmptyResult(err) {
		s.log.WithError(err).Error("Failed to get seller ID from query")
		response.Error(w, http.StatusBadRequest, "Bad Request", "Failed to get seller ID from query")
		return
	}
	var items []*dbitems.Item
	if sellerID != uuid.Nil {
		items, err = s.db.GetItemsBySellerID(r.Context(), dbitems.NewWithSellerIDAndAmount(sellerID, amount))
	} else {
		items, err = s.db.GetItems(r.Context(), dbitems.NewWithAmount(amount))
	}
	if err != nil {
		s.log.WithError(err).Error("Failed to get items")
		response.Error(w, http.StatusInternalServerError, "Server error", "Failed to get items")
		return
	}

	response.Success(w, http.StatusOK, items)
}

// @Summary		Get item
// @Description	Get item by ID
// @Tags			sales
// @Produce		json
// @Param			id	path		string				true	"Item ID"
// @Success		200	{object}	response.ItemDTO	"Item"
// @Failure		400	{object}	response.ErrorDTO	"Bad Request"
// @Failure		500	{object}	response.ErrorDTO	"Server Error"
// @Router			/sales/items/{id} [get]
func (s *SalesHandlers) GetItem(w http.ResponseWriter, r *http.Request) {
	itemID, err := handlershelp.GetUUIDFromPath(r, "id")
	if err != nil {
		s.log.WithError(err).Error("Failed to get item ID from path")
		response.Error(w, http.StatusBadRequest, "Bad Request", "Failed to get item ID from path")
		return
	}
	item, err := s.db.GetItem(r.Context(), dbitems.NewWithID(itemID))
	if err != nil {
		s.log.WithError(err).Error("Failed to get item by ID")
		response.Error(w, http.StatusInternalServerError, "Server error", "Failed to get item by ID")
		return
	}
	response.Success(w, http.StatusOK, item)
}
