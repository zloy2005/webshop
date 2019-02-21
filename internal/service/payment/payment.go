package payment

import (
	"github.com/sirupsen/logrus"
	"github.com/zloy2005/signedpay-api"

	"github.com/zloy2005/webshop/internal/model"
	"github.com/zloy2005/webshop/internal/storage"
)

const (
	platform      = "WEB"
	ip            = "8.8.8.8"
	customerEmail = "test@test.com"
)

type Payment struct {
	client  *signedpay_api.Api
	storage *storage.Storage
}

func New(storage *storage.Storage, config Config) *Payment {
	client := signedpay_api.New(config.Merchant, config.Key, true)
	return &Payment{client: client, storage: storage}
}

func (p *Payment) PayByNewCard(order *model.PaymentOrder, card *model.Card) (*model.PaymentOrder, error) {
	resp, err := p.client.Charge(&signedpay_api.ChargeRequest{
		OrderID:          order.No,
		Amount:           order.Amount,
		OrderDescription: order.Description,
		Currency:         "USD",
		CustomerEmail:    customerEmail,
		IPAddress:        ip,
		Platform:         platform,
		CardCvv:          card.Cvv,
		CardExpMonth:     card.ExpMonth,
		CardExpYear:      card.ExpYear,
		CardNumber:       card.Number,
	})
	if err != nil {
		if e, ok := err.(signedpay_api.PayError); ok {
			order.Status = e.Error()
			return p.storage.NewPaymentOrder(order)
		}
		return nil, err
	}
	order.Status = resp.Order.Status
	return p.storage.NewPaymentOrder(order)
}

func (p *Payment) PayBySavedCard(order *model.PaymentOrder, cardToken string) (*model.PaymentOrder, error) {
	resp, err := p.client.Recurring(&signedpay_api.RecurringRequest{
		OrderID:          order.No,
		Amount:           order.Amount,
		OrderDescription: order.Description,
		RecurringToken:   cardToken,
		Currency:         "USD",
		CustomerEmail:    customerEmail,
		IPAddress:        ip,
		Platform:         platform,
	})
	if err != nil {
		if e, ok := err.(signedpay_api.PayError); ok {
			order.Status = e.Error()
			return p.storage.NewPaymentOrder(order)

		}
		return nil, err
	}
	order.Status = resp.Order.Status
	return p.storage.NewPaymentOrder(order)
}

func (p *Payment) Refund(order *model.PaymentOrder) (*model.PaymentOrder, error) {
	resp, err := p.client.Refund(&signedpay_api.RefundRequest{
		OrderID: order.No,
		Amount:  order.Amount,
	})
	if err != nil {
		if e, ok := err.(signedpay_api.PayError); ok {
			order.Status = e.Error()
			return p.storage.NewPaymentOrder(order)
		}
		return nil, err
	}
	order.RefundedAmount = resp.Order.RefundedAmount
	order.Status = resp.Order.Status
	return p.storage.UpdatePaymentOrder(order)
}

func (p *Payment) Status(order *model.PaymentOrder) (*model.PaymentOrder, error) {
	resp, err := p.client.Status(&signedpay_api.StatusRequest{
		OrderID: order.No,
	})
	if err != nil {
		if e, ok := err.(signedpay_api.PayError); ok {
			order.Status = e.Error()
			return p.storage.NewPaymentOrder(order)

		}
		return nil, err
	}
	order.Status = resp.Order.Status
	if order.Status == "approved" {
		for _, t := range resp.Transactions {
			if t.Card.CardToken.Token != "" && t.Card.Number != "" {
				if err := p.storage.SaveCard(t.Card.Number, t.Card.CardToken.Token); err != nil {
					logrus.Error(err)
				}
			}
		}

	}
	return p.storage.UpdatePaymentOrder(order)
}
