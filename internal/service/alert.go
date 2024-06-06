package service

import (
	"fmt"

	"rutube-task/internal/config"
	"rutube-task/internal/repository"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"
)

type Alert struct {
	cfg  *config.Config
	cron *cron.Cron
	rep  *repository.Repository
	bot  *tgbotapi.BotAPI
}

func NewAlert(cfg *config.Config, rep *repository.Repository) *Alert {
	bot, err := tgbotapi.NewBotAPI(cfg.BotAPI)
	if err != nil {
		logrus.Error("error creating telegram bot:", err)
		return nil
	}
	return &Alert{
		cfg:  cfg,
		cron: cron.New(),
		bot:  bot,
		rep:  rep,
	}
}

func (a *Alert) Scheduler() {
	a.cron.AddFunc("*/5 * * * * *", func() {
		a.getTodayBirthday()
	})
	a.cron.Start()
	defer a.cron.Stop()
	select {}
}

func (a *Alert) getTodayBirthday() {
	employeeList, err := a.rep.GetEmployeeBirthdayDB()
	if err != nil {
		logrus.Error(err)
	}

	for _, employee := range employeeList {
		msg := fmt.Sprintf("Сегодня день рождения у %s", employee.Name)
		err := a.sendMessage(msg)
		if err != nil {
			logrus.Error("getTodayBirthday: ", err)
		}
	}
}

func (a *Alert) sendMessage(msg string) error {
	ms := tgbotapi.NewMessageToChannel(a.cfg.ChatID, msg)
	_, err := a.bot.Send(ms)
	if err != nil {
		return fmt.Errorf("error sending message: %w", err)
	}

	return nil
}
