package service

import (
	"TrackerBot/Internal/models"
	"TrackerBot/Internal/repository"
)

type HabitsService struct {
	h *repository.HabitsRepo
}

func NewHabitsService(h *repository.HabitsRepo) *HabitsService {
	return &HabitsService{h: h}
}

func (c *HabitsService) CreateHabits(m *models.Habits) error {
	err := c.h.CreateHabits(m)
	if err != nil {
		return err
	}
	return nil
}

func (c *HabitsService) GetHabits(id int64) ([]*models.Habits, error) {
	res, err := c.h.GetHabits(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
