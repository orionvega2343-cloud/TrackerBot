package repository

import (
	"TrackerBot/Internal/models"

	"github.com/jmoiron/sqlx"
)

type HabitsRepo struct {
	db *sqlx.DB
}

func NewHabitsRepo(db *sqlx.DB) *HabitsRepo {
	return &HabitsRepo{db: db}
}

func (r *HabitsRepo) CreateHabits(m *models.Habits) error {
	_, err := r.db.Exec(`INSERT INTO habits(user_id,title) VALUES($1,$2,$3)`, m.Id, m.UserId, m.Title)
	if err != nil {
		return err
	}
	return nil
}

func (r *HabitsRepo) GetHabits(id int) ([]*models.Habits, error) {
	var habits []*models.Habits
	query := `SELECT id,user_id,title FROM habits WHERE user_id=$1`

	err := r.db.Select(&habits, query, id)
	if err != nil {
		return nil, err
	}
	return habits, nil
}
