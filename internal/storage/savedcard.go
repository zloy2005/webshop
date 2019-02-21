package storage

import (
	"github.com/zloy2005/webshop/internal/model"
)

func (s *Storage) SavedCards() ([]*model.SaveCard, error) {
	var savedCard []*model.SaveCard
	if result := s.db.Find(&savedCard); result.Error != nil {
		return nil, result.Error
	}
	return savedCard, nil
}

func (s *Storage) SavedCard(id int) (*model.SaveCard, error) {
	savedCard := &model.SaveCard{}
	if result := s.db.Find(savedCard, id); result.Error != nil {
		return nil, result.Error
	}
	return savedCard, nil
}

func (s *Storage) SaveCard(cardMask, token string) error {
	var savedCard []*model.SaveCard
	if result := s.db.Where("card_mask = ? AND token= ?", cardMask, token).Find(&savedCard); result.Error != nil {
		return result.Error
	}
	if len(savedCard) == 0 {
		if result := s.db.Create(&model.SaveCard{CardMask: cardMask, Token: token}); result.Error != nil {
			return result.Error
		}
	}
	return nil
}
