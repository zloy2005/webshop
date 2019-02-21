package handler

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/pborman/uuid"
	"github.com/sirupsen/logrus"

	"github.com/zloy2005/webshop/internal/model"
	"github.com/zloy2005/webshop/internal/service/payment"
	"github.com/zloy2005/webshop/internal/storage"
)

type PaymentHandler struct {
	storage *storage.Storage
	payment *payment.Payment
}

type NewPaymentParam struct {
	Product         *model.Product
	NotHasSavedCard bool
	SavedCards      []*model.SaveCard
}

func NewPaymentHandler(storage *storage.Storage, payment *payment.Payment) *PaymentHandler {
	return &PaymentHandler{storage: storage, payment: payment}
}

func (h *PaymentHandler) NewPayment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	productID := r.FormValue("product")
	pID, _ := strconv.Atoi(productID)

	logrus.Info(productID)
	product, err := h.storage.Product(pID)
	if err != nil {
		logrus.Error(err)
		return
	}
	savedCards, err := h.storage.SavedCards()
	if err != nil {
		logrus.Error(err)
		return
	}
	data := NewPaymentParam{
		Product:         product,
		NotHasSavedCard: len(savedCards) == 0,
		SavedCards:      savedCards,
	}
	tmpl, err := template.ParseFiles("web/templates/dashboard.html", "web/templates/new_payment.html")
	if err != nil {
		logrus.Error(err)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		logrus.Error(err)
		return
	}
}

func (h *PaymentHandler) NewPaymentOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	productID := r.FormValue("product")
	pID, _ := strconv.Atoi(productID)
	product, err := h.storage.Product(pID)
	if err != nil {
		logrus.Error(err)
		return
	}
	order := &model.PaymentOrder{No: uuid.New(), Description: product.Description, Amount: product.Price}
	if r.FormValue("method") == "new" {
		cardNumber := r.FormValue("card-number")
		cvc := r.FormValue("cvc")
		expMonth := r.FormValue("exp-month")
		expYear := r.FormValue("exp-year")

		_, err := h.payment.PayByNewCard(order, &model.Card{Number: cardNumber, ExpMonth: expMonth, ExpYear: expYear, Cvv: cvc})
		if err != nil {
			logrus.Error(err)
			return
		}
	} else {
		savedCardID := r.FormValue("saved-card")
		id, _ := strconv.Atoi(savedCardID)
		savedCard, err := h.storage.SavedCard(id)
		if err != nil {
			logrus.Error(err)
			return
		}
		_, err = h.payment.PayBySavedCard(order, savedCard.Token)
		if err != nil {
			logrus.Error(err)
			return
		}
	}
	http.Redirect(w, r, "/orders", 302)
}
