package handler

import (
	"html/template"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/zloy2005/webshop/internal/storage"
)

type HomeHandler struct {
	storage *storage.Storage
}

func NewHomeHandler(storage *storage.Storage) *HomeHandler {
	return &HomeHandler{storage: storage}
}

func (h *HomeHandler) GetHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	products, err := h.storage.Products()
	if err != nil {
		logrus.Error(err)
		return
	}
	tmpl, err := template.ParseFiles("web/templates/dashboard.html", "web/templates/home.html")
	if err != nil {
		logrus.Error(err)
		return
	}

	if err := tmpl.Execute(w, products); err != nil {
		logrus.Error(err)
		return
	}
}
