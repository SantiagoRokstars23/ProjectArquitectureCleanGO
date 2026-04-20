package handlers

import (
	"GO_PROJECT_ARQUITECTURE/internal/usecase"
	"encoding/json"
	"net/http"
)

type Handler struct {
	listCardsUC *usecase.ListCardsUseCase
}

func NewHandler(uc *usecase.ListCardsUseCase) *Handler {
	return &Handler{listCardsUC: uc}
}

func (h *Handler) ListCards(w http.ResponseWriter, r *http.Request) {
	cards, err := h.listCardsUC.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(cards)

}
