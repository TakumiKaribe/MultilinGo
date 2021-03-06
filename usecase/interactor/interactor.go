package interactor

import (
	"multilingo/infrastructure/request/paiza"
	"multilingo/logger"
	"multilingo/usecase/interactor/buildMessage"
	"multilingo/usecase/interactor/parsetext"
	"multilingo/usecase/interfaces"
)

// Interactor -
type Interactor struct {
	presenter interfaces.Presenter
}

// NewInteractor -
func NewInteractor(presenter interfaces.Presenter) *Interactor {
	return &Interactor{presenter: presenter}
}

// ExecProgram -
func (i *Interactor) ExecProgram(language string, text string) error {
	// parse program
	input, program, err := parsetext.Parse(text)
	if err != nil {
		i.presenter.ShowError(err)
		logger.Log.Warn(err)
		return err
	}

	client := paiza.NewClient()
	result, err := client.Request(input, language, program)
	if err != nil {
		i.presenter.ShowError(err)
		return err
	}

	i.presenter.ShowResult(buildMessage.MakeMessage(result))

	return nil
}
