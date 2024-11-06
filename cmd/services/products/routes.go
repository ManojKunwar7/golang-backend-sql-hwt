package products

import (
	"fmt"
	"net/http"
	"test-project/helper"
	"test-project/types"

	"github.com/gorilla/mux"
)

type Handler struct {
	ProductStore types.ProductStore
}

func NewHandler(ProductStore types.ProductStore) *Handler {
	return &Handler{ProductStore: ProductStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/products", h.handleListProducts).Methods("GET")
}

func (h *Handler) handleListProducts(w http.ResponseWriter, r *http.Request) {
	p, err := h.ProductStore.GetProducts()

	if err != nil {
		helper.WriteError(w, http.StatusInternalServerError, fmt.Errorf("something went wrong"))
	}

	helper.WriteJson(w, http.StatusOK, map[string]any{"status": true, "c_msg": "Data Found!", "alert_status": "true", "data": p})
}
