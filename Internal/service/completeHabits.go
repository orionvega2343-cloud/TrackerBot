package service

import (
	"TrackerBot/Internal/models"
	"TrackerBot/Internal/repository"
)

type CompleteService struct {
	c *repository.CompleteRepo
}

func NewCompleteService(c *repository.CompleteRepo) *CompleteService {
	return &CompleteService{c: c}
}

func (s *CompleteService) Complete(m *models.CompleteHabits) error {
	err := s.c.CreateCompleteHabits(m)
	if err != nil {
		return err
	}
	return nil
}

func (s *CompleteService) GetHabits(id int) ([]*models.CompleteHabits, error) {
	res, err := s.c.GetCompleteHabits(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *CompleteService) Streak(id int) (int, error) {
	var streak int
	res, err := s.c.GetCompleteHabits(id)
	if err != nil {
		return streak, err
	}
	for _, v := range res {
		if v.IsComplete == true {
			streak++
		} else if v.IsComplete == false {
			break
		}
	}
	return streak, nil
}
