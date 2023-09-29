package service

import (
	"database/sql"

	"github.com/maguro-alternative/ganbaride_bespa_api/model"
)

type HeroCardService struct {
	db *sql.DB
}

// NewTODOService returns new TODOService.
func NewHeroCardService(db *sql.DB) *HeroCardService {
	return &HeroCardService{
		db: db,
	}
}

func (s *HeroCardService) GetHeroCardList() ([]*model.Rider, error) {
	return nil,nil
}
