package handlers

import (
	"TrackerBot/Internal/models"
	"TrackerBot/Internal/service"
	"fmt"
	"strconv"

	tele "gopkg.in/telebot.v3"
)

type HabitsHandler struct {
	s   *service.HabitsService
	bot *tele.Bot
}

func NewHabitsHandler(s *service.HabitsService, bot *tele.Bot) *HabitsHandler {
	return &HabitsHandler{s: s, bot: bot}
}

var uState = map[int64]string{}

var (
	markup = &tele.ReplyMarkup{
		ResizeKeyboard: true, // подстроить размер
	}
	btn1 = markup.Text("➕Добавить привычку")
	btn2 = markup.Text("👁Посмотреть список привычек")
	btn3 = markup.Text("👁Показать серию")
)

func (h *HabitsHandler) Reg() {
	h.bot.Handle("/start", func(c tele.Context) error {
		markup.Reply(
			markup.Row(btn1, btn2, btn3),
		)
		return c.Send("Привет! Я бот трекер привычек, помогу тебе бросить вредные привычки или приобрести полезные,для начала работы выбери действие ниже!", markup)
	})
	h.bot.Handle(btn1, h.AddHabitsHandler)
	h.bot.Handle(btn2, h.GetHabitsHandler)
	h.bot.Handle(tele.OnText, h.TextHandler)

}

func (h *HabitsHandler) AddHabitsHandler(c tele.Context) error {
	uId := c.Sender().ID
	uState[uId] = "adding_habit"
	return c.Send("Введите название привычки:")
}

func (h *HabitsHandler) TextHandler(c tele.Context) error {
	uId := c.Sender().ID
	var m models.Habits
	msg := c.Message().Text
	m.Title = msg
	m.UserId = int(uId)
	if uState[c.Sender().ID] == "adding_habit" {
		err := h.s.CreateHabits(&m)
		if err != nil {
			return err
		}
		uState[c.Sender().ID] = ""
		return c.Send("Привычка успешно добавлена")
	}
	return nil
}

func (h *HabitsHandler) GetHabitsHandler(c tele.Context) error {

	uId := c.Sender().ID
	err := c.Send("Список привычек")
	if err != nil {
		return err
	}
	res, err := h.s.GetHabits(int(uId))
	if err != nil {
		return err
	}
	for _, v := range res {
		parsed := strconv.Itoa(v.Id)
		mp := h.bot.NewMarkup()
		completeBtn := mp.Data("Отметить как выполненное", "habits_id", parsed)
		mp.Inline(
			mp.Row(completeBtn),
		)
		err = c.Send(fmt.Sprintf("Ваши привычки: %s", v.Title), mp)
		if err != nil {
			return err
		}
	}
	return nil
}
