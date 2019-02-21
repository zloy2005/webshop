package storage

import (
	"github.com/jinzhu/gorm"
	"github.com/zloy2005/webshop/internal/model"
)

type Storage struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Storage {
	return &Storage{db: db}
}

func (s *Storage) Migrate() {
	s.db.AutoMigrate(&model.Product{})
	s.db.AutoMigrate(&model.PaymentOrder{})
	s.db.AutoMigrate(&model.SaveCard{})

	s.db.Unscoped().Delete(&model.Product{})
	s.db.Create(&model.Product{Code: "ITM-001", Description: " Item #001", Image: "images/items/1.jpg", Price: 100})
	s.db.Create(&model.Product{Code: "ITM-002", Description: " Item #002", Image: "images/items/2.jpg", Price: 150})
	s.db.Create(&model.Product{Code: "ITM-003", Description: " Item #003", Image: "images/items/3.jpg", Price: 75})
}
