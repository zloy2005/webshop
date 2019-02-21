package handler

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/zloy2005/webshop/internal/service/payment"
	"github.com/zloy2005/webshop/internal/storage"
)

type OrderHandler struct {
	storage *storage.Storage
	payment *payment.Payment
}

func NewOrderHandler(storage *storage.Storage, payment *payment.Payment) *OrderHandler {
	return &OrderHandler{storage: storage, payment: payment}
}

func (h *OrderHandler) Orders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	orders, err := h.storage.PaymentOrders()
	if err != nil {
		logrus.Error(err)
		return
	}
	tmpl, err := template.ParseFiles("web/templates/dashboard.html", "web/templates/orders.html")
	if err != nil {
		logrus.Error(err)
		return
	}
	if err := tmpl.Execute(w, orders); err != nil {
		logrus.Error(err)
		return
	}
}

func (h *OrderHandler) OrderRefund(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	orderID := r.FormValue("order")
	id, _ := strconv.Atoi(orderID)
	order, err := h.storage.PaymentOrder(id)
	if err != nil {
		logrus.Error(err)
		return
	}

	if _, err := h.payment.Refund(order); err != nil {
		logrus.Error(err)
	}

	http.Redirect(w, r, "/orders", 302)

}
