package handlers

import "TrackerBot/Internal/service"

type HabitsHandler struct {
	s *service.HabitsService
}
func NewHabitsHandler(s *service.HabitsService) *HabitsHandler {
	return &HabitsHandler{s: s}
}

