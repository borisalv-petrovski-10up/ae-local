package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	
	"github.com/borisalv-petrovski-10up/ae-local/repository"
)

// Handler is a handler for the application.
type Handler struct {
	repository repository.Repository
}

// NewHandler creates a new handler.
func NewHandler(repository repository.Repository) Handler {
	return Handler{
		repository: repository,
	}
}

// CreateArticle creates an article.
func (h Handler) CreateArticle(w http.ResponseWriter, r *http.Request) {
	var article repository.Article
	err := json.NewDecoder(r.Body).Decode(&article)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	err = h.repository.CreateArticle(r.Context(), &article)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	_, _ = fmt.Fprintf(w, "%v", article)
}

// GetArticle gets an article by title.
func (h Handler) GetArticle(w http.ResponseWriter, r *http.Request) {
	title := strings.TrimPrefix(r.URL.Path, "/get-article/")
	article, err := h.repository.GetArticle(r.Context(), title)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	_, _ = fmt.Fprintf(w, "%v", article)
}
