package saleshandlers

import (
	"io"
	"net/http"
	"os"
	"path"
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
	mux.HandleFunc("GET /sales/items", s.GetItems)                // multiple items(all or by seller)
	mux.HandleFunc("GET /sales/items/{id}", s.GetItem)            // single item full by ID
	mux.HandleFunc("GET /sales/items/search", s.GetItemsBySearch) // search items by query in query string
	mux.Handle("/sales/items/images/{id}",
		handlershelp.AuthMiddleware(
			s.log,
			handlershelp.ItemsAuthMiddleware(
				s.log,
				s.db,
				http.HandlerFunc(s.CreateItemImage),
			),
		),
	) // create item image for item by ID
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
		items, err = s.db.GetItemsBySellerID(
			r.Context(),
			dbitems.NewItem(dbitems.WithSellerID(sellerID), dbitems.WithAmount(amount)),
		)
	} else {
		items, err = s.db.GetItems(r.Context(), dbitems.NewItem(dbitems.WithAmount(amount)))
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
	item, err := s.db.GetItem(r.Context(), dbitems.NewItem(dbitems.WithID(itemID)))
	if err != nil {
		s.log.WithError(err).Error("Failed to get item by ID")
		response.Error(w, http.StatusInternalServerError, "Server error", "Failed to get item by ID")
		return
	}
	response.Success(w, http.StatusOK, item)
}

// @Summary		Get items by search
// @Description	Get items by search
// @Tags			sales
// @Produce		json
// @Param			query	query		string					true	"Query"
// @Param			amount	query		int						false	"Amount of items to return"	default(10)
// @Success		200		{object}	response.ItemsListDTO	"Items list"
// @Failure		400		{object}	response.ErrorDTO		"Bad Request"
// @Failure		500		{object}	response.ErrorDTO		"Server Error"
// @Router			/sales/items/search [get]
func (s *SalesHandlers) GetItemsBySearch(w http.ResponseWriter, r *http.Request) {
	query, err := handlershelp.GetStringFromQuery(r, "query")
	if err != nil {
		s.log.WithError(err).Debug("Failed to get query from query")
		response.Error(w, http.StatusBadRequest, "Bad Request", "Failed to get query from query")
		return
	}
	amount, err := handlershelp.GetIntFromQuery(r, "amount", dbitems.DefaultAmount)
	if err != nil {
		s.log.WithError(err).Debug("Failed to get amount from query")
		response.Error(w, http.StatusBadRequest, "Bad Request", "Failed to get amount from query")
		return
	}
	items, err := s.db.GetItemsBySearch(
		r.Context(),
		dbitems.NewItem(dbitems.WithQuery(query), dbitems.WithAmount(amount)),
	)
	if err != nil {
		s.log.WithError(err).Error("Failed to get items by search")
		response.Error(w, http.StatusInternalServerError, "Server error", "Failed to get items by search")
		return
	}
	response.Success(w, http.StatusOK, items)
}

// @Summary		Create item image
// @Description	Create item image for item by ID
// @Tags			sales
// @Produce		json
// @Param			id		path		string						true	"Item ID"
// @Param			image	formData	file						true	"Image"
// @Success		200		{object}	response.ImageIDResponse	"Image ID"
// @Failure		400		{object}	response.ErrorDTO			"Bad Request"
// @Failure		500		{object}	response.ErrorDTO			"Server Error"
// @Router			/sales/items/images/{id} [post]
// @Security		Bearer
func (s *SalesHandlers) CreateItemImage(w http.ResponseWriter, r *http.Request) {
	itemID, err := handlershelp.GetUUIDFromPath(r, "id")
	if err != nil {
		s.log.WithError(err).Error("Failed to get item ID from path")
		response.Error(w, http.StatusBadRequest, "Bad Request", "Failed to get item ID from path")
		return
	}
	claims, err := handlershelp.GetClaimsFromContext(r)
	if err != nil {
		s.log.WithError(err).Error("Failed to get claims from context")
		response.Error(w, http.StatusBadRequest, "Bad Request", "Failed to get claims from context")
		return
	}
	err = r.ParseMultipartForm(10 << 20) // 10MB
	if err != nil {
		s.log.WithError(err).Error("Failed to parse multipart form")
		response.Error(w, http.StatusBadRequest, "Bad Request", "Failed to parse multipart form")
		return
	}
	image, header, err := r.FormFile("image")
	if err != nil {
		s.log.WithError(err).Error("Failed to get image from form")
		response.Error(w, http.StatusBadRequest, "Bad Request", "Failed to get image from form")
		return
	}
	defer image.Close()
	buffer := make([]byte, 512)
	_, err = image.Read(buffer)
	if err != nil {
		s.log.WithError(err).Error("Failed to read image")
		response.Error(w, http.StatusBadRequest, "Bad Request", "Failed to read image")
		return
	}
	imageType := http.DetectContentType(buffer)
	if imageType != "image/jpeg" && imageType != "image/png" &&
		imageType != "image/gif" && imageType != "image/webp" {
		s.log.Error("Invalid image type")
		response.Error(w, http.StatusBadRequest, "Bad Request", "Invalid image type")
		return
	}
	_, err = image.Seek(0, io.SeekStart)
	if err != nil {
		s.log.WithError(err).Error("Failed to seek image")
		response.Error(w, http.StatusBadRequest, "Bad Request", "Failed to seek image")
		return
	}

	imageUUID := uuid.New()
	imagePath := path.Join(
		s.config.DatabaseConfig.ItemsImagesPath,
		claims.UserID.String(),
		itemID.String(),
		imageUUID.String()+"."+header.Filename,
	)
	os.MkdirAll(path.Dir(imagePath), 0755)

	imageFile, err := os.OpenFile(imagePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		s.log.WithError(err).Error("Failed to create image file")
		response.Error(w, http.StatusInternalServerError, "Server error", "Failed to create image file")
		return
	}
	defer imageFile.Close()

	io.Copy(imageFile, image)
	response.Success(w, http.StatusOK, response.ImageIDResponse{
		ID:   imageUUID,
		Path: imagePath,
		Type: imageType,
		Size: header.Size,
	})
}
