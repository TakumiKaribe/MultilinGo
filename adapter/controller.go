package adapter

import (
	"github.com/TakumiKaribe/multilingo/entity"
	"github.com/TakumiKaribe/multilingo/entity/config"
	"github.com/TakumiKaribe/multilingo/usecase/interactor"
	"github.com/TakumiKaribe/multilingo/usecase/interfaces"
	log "github.com/sirupsen/logrus"
)

// Controller is controller by clean architecture
type Controller struct {
	body    *entity.APIGateWayRequestBody
	useCase interfaces.UseCase
}

// NewController is initialize Controller
func NewController(requestBody *entity.APIGateWayRequestBody) (*Controller, error) {
	presenter, err := NewPresenter(requestBody)
	if err != nil {
		// TODO: use multilingo error
		log.Warn(err.Error())
		return nil, err
	}
	return &Controller{body: requestBody, useCase: interactor.NewInteractor(presenter)}, nil
}

// ExecProgram is to exec program by post to paiza
func (c *Controller) ExecProgram() error {
	// init bot_info
	bot, err := config.SharedConfig.NewBotInfo(c.body.APIAppID)
	if err != nil {
		log.Warn(err.Error())
		return err
	}

	return c.useCase.ExecProgram(bot.Language, c.body.Event.Text)
}

// Challenge is to validate token
func (c *Controller) Challenge() {
	c.useCase.Challenge()
}

// Kick is to kick corresponding to bot
func (c *Controller) Kick() {
	c.useCase.Kick()
}