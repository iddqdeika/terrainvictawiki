package project

import (
	"encoding/json"
	"net/http"
)

func NewHTTPHandler(repo *Repository) *Handler {
	return &Handler{repo: repo}
}

type Handler struct {
	repo *Repository
}

func (h *Handler) HandleGetAll(writer http.ResponseWriter, request *http.Request) {
	all := h.repo.GetAll()
	data, err := json.MarshalIndent(all, "", "\t")
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.Write(data)
}
