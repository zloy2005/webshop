package storage

import (
	"github.com/zloy2005/webshop/internal/model"
)

func (s *Storage) Products() ([]*model.Product, error) {
	var products []*model.Product
	if result := s.db.Find(&products); result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

func (s *Storage) Product(id int) (*model.Product, error) {
	product := &model.Product{}
	if result := s.db.Find(product, id); result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}
