package service

import (
	"context"
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

func (s *HeroCardService) GetHeroCardList(ctx context.Context) ([]*model.Rider, error) {
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

func (s *HeroCardService) GetHeroCard(number string) (*model.Rider, error) {
	const sqlStr = `SELECT * FROM rider WHERE number = ?`
	row := s.db.QueryRow(sqlStr, number)

	var rider model.Rider
	if err := row.Scan(
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
	return &rider, nil
}

func (s *HeroCardService) GetHeroCardListByLot(ctx context.Context,lot string) ([]*model.Rider, error) {
	var sqlStr = `SELECT * FROM rider WHERE number = "?-___"`
	rows, err := s.db.QueryContext(ctx, sqlStr, lot)
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