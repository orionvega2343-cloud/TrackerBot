package handlers

import (
	"TrackerBot/Internal/models"
	"TrackerBot/Internal/service"
	"fmt"
	"strconv"
	"time"

	tele "gopkg.in/telebot.v3"
)

type CompleteHabits struct {
	s   *service.CompleteService
	bot *tele.Bot
}

func NewCompleteHabitsHandler(s *service.CompleteService, bot *tele.Bot) *CompleteHabits {
	return &CompleteHabits{s: s, bot: bot}
}

func (h *CompleteHabits) Reg() {
	h.bot.Handle(tele.OnCallback, h.CompleteHabitHandler)
	h.bot.Handle(btn3, h.StreakHandler)
}

func (h *CompleteHabits) CompleteHabitHandler(c tele.Context) error {
	data := c.Data()
	parsed, err := strconv.Atoi(data)
	if err != nil {
		return err
	}
	var m models.CompleteHabits
	m.IsComplete = true
	m.HabitId = parsed
	m.Date = time.Now()
	err = h.s.Complete(&m)
	if err != nil {
		return err
	}
	c.Respond()
	return nil
}

func (h *CompleteHabits) StreakHandler(c tele.Context) error {
	uId := c.Sender().ID
	res, err := h.s.Streak(int(uId))
	if err != nil {
		return err
	}
	return c.Send(fmt.Sprintf("Ваша серия длится %d дней", res))
}
