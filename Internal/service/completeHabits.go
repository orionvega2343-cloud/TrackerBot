package service

import (
	"TrackerBot/Internal/models"
	"TrackerBot/Internal/repository"
	"time"
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

func (s *CompleteService) Streak(id int64) (int, error) {
	res, err := s.c.GetCompleteHabits(id)
	if err != nil {
		return 0, err
	}
	if len(res) == 0 {
		return 0, nil
	}

	days := make(map[string]bool)
	for _, v := range res {
		days[v.Date.Format("2006-01-02")] = true
	}

	streak := 0
	day := time.Now()
	for {
		if !days[day.Format("2006-01-02")] {
			break
		}
		streak++
		day = day.AddDate(0, 0, -1)
	}
	return streak, nil
}
