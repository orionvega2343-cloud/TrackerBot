package models

import "time"

type CompleteHabits struct {
	Id         int       `db:"id"`
	HabitId    int       `db:"habit_id"`
	Date       time.Time `db:"date"`
	IsComplete bool      `db:"is_complete"`
}
