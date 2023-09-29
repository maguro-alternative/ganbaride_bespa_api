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
	const sqlStr = `SELECT * FROM rider`
	rows, err := s.db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var riders []*model.Rider
	for rows.Next() {
		var rider model.Rider
		if err := rows.Scan(
			&rider.Number,
			&rider.Name,
			&rider.Rarity,
			&rider.Category,
			&rider.Attribute,
			&rider.Attack,
			&rider.Defence,
			&rider.HitPoint,
			&rider.FinishingAttack,
			&rider.SkillName,
			&rider.AttackTeki,
			&rider.DefenceTeki,
			&rider.HitPointTeki,
			&rider.FinishingAttackTeki,
		); err != nil {
			return nil, err
		}
		riders = append(riders, &rider)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return riders, nil
}
