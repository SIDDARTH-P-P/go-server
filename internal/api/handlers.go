package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"go-server/internal/store"
)

type Handler struct {
	Store *store.Store
}

type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
	Message string      `json:"message,omitempty"`
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	items := h.Store.List()
	message := ""
	if len(items) == 0 {
		message = "no items found"
	}

	writeJSON(w, http.StatusOK, APIResponse{
		Success: true,
		Data:    items,
		Message: message,
	})
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Title string `json:"title"`
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	payload.Title = strings.TrimSpace(payload.Title)
	if payload.Title == "" {
		writeError(w, http.StatusBadRequest, "title is required")
		return
	}

	item, err := h.Store.Create(payload.Title)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to create item")
		return
	}

	writeJSON(w, http.StatusCreated, APIResponse{
		Success: true,
		Data:    item,
		Message: "item created successfully",
	})
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}

	item, err := h.Store.Get(id)
	if err != nil {
		if errors.Is(err, store.ErrItemNotFound) {
			writeError(w, http.StatusNotFound, "item not found")
			return
		}
		writeError(w, http.StatusInternalServerError, "failed to get item")
		return
	}

	writeJSON(w, http.StatusOK, APIResponse{
		Success: true,
		Data:    item,
	})
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}

	var payload struct {
		Title string `json:"title"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	payload.Title = strings.TrimSpace(payload.Title)
	if payload.Title == "" {
		writeError(w, http.StatusBadRequest, "title is required")
		return
	}

	item, err := h.Store.Update(id, payload.Title)
	if err != nil {
		if errors.Is(err, store.ErrItemNotFound) {
			writeError(w, http.StatusNotFound, "item not found")
			return
		}
		writeError(w, http.StatusInternalServerError, "failed to update item")
		return
	}

	writeJSON(w, http.StatusOK, APIResponse{
		Success: true,
		Data:    item,
		Message: "item updated successfully",
	})
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}

	if err := h.Store.Delete(id); err != nil {
		if errors.Is(err, store.ErrItemNotFound) {
			writeError(w, http.StatusNotFound, "item not found")
			return
		}
		writeError(w, http.StatusInternalServerError, "failed to delete item")
		return
	}

	writeJSON(w, http.StatusOK, APIResponse{
		Success: true,
		Message: "item deleted successfully",
	})
}

func writeJSON(w http.ResponseWriter, status int, payload APIResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, APIResponse{
		Success: false,
		Error:   message,
	})
}
