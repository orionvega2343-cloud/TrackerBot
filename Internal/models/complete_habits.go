package models

import "time"

type CompleteHabits struct {
	Id         int
	HabitId    int
	Date       time.Time
	IsComplete bool
}
