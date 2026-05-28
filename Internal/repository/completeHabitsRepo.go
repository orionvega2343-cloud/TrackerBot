package repository

import (
	"TrackerBot/Internal/models"

	"github.com/jmoiron/sqlx"
)

type CompleteRepo struct {
	db *sqlx.DB
}

func NewCompleteRepo(db *sqlx.DB) *CompleteRepo {
	return &CompleteRepo{db: db}
}

func (c *CompleteRepo) CreateCompleteHabits(m *models.CompleteHabits) error {
	_, err := c.db.Exec(`INSERT INTO complete_habits(habit_id,date,is_complete) VALUES($1,$2,$3)`, m.HabitId, m.Date, m.IsComplete)
	if err != nil {
		return err
	}
	return nil
}

func (c *CompleteRepo) GetCompleteHabits(userId int64) ([]*models.CompleteHabits, error) {
	var habits []*models.CompleteHabits

	query := `SELECT ch.id, ch.habit_id, ch.date, ch.is_complete
		FROM complete_habits ch
		JOIN habits h ON h.id = ch.habit_id
		WHERE ch.is_complete = true AND h.user_id = $1
		ORDER BY ch.date DESC`
	err := c.db.Select(&habits, query, userId)
	if err != nil {
		return nil, err
	}
	return habits, nil
}
