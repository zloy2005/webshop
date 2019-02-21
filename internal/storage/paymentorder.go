package storage

import (
	"github.com/zloy2005/webshop/internal/model"
)

func (s *Storage) PaymentOrder(id int) (*model.PaymentOrder, error) {
	order := &model.PaymentOrder{}
	if result := s.db.Find(order, id); result.Error != nil {
		return nil, result.Error
	}
	return order, nil
}

func (s *Storage) NewPaymentOrder(order *model.PaymentOrder) (*model.PaymentOrder, error) {
	if result := s.db.Create(order); result.Error != nil {
		return nil, result.Error
	}
	return order, nil
}

func (s *Storage) UpdatePaymentOrder(order *model.PaymentOrder) (*model.PaymentOrder, error) {
	if result := s.db.Save(order); result.Error != nil {
		return nil, result.Error
	}
	return order, nil
}

func (s *Storage) PaymentOrders() ([]*model.PaymentOrder, error) {
	var orders []*model.PaymentOrder
	if result := s.db.Find(&orders); result.Error != nil {
		return nil, result.Error
	}
	return orders, nil
}

func (s *Storage) PaymentOrdersInProccess() ([]*model.PaymentOrder, error) {
	var orders []*model.PaymentOrder

	if result := s.db.Where("status in (?)", []string{"created", "processing"}).Find(&orders); result.Error != nil {
		return nil, result.Error
	}
	return orders, nil
}
