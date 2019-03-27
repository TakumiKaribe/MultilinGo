package interactor

import (
	"github.com/TakumiKaribe/multilingo/infrastructure/request/paiza"
	"github.com/TakumiKaribe/multilingo/usecase/interactor/buildMessage"
	"github.com/TakumiKaribe/multilingo/usecase/interactor/parsetext"
	"github.com/TakumiKaribe/multilingo/usecase/interfaces"
)

type Interactor struct {
	presenter interfaces.Presenter
}

func NewInteractor(presenter interfaces.Presenter) *Interactor {
	return &Interactor{presenter: presenter}
}

func (i *Interactor) ExecProgram(language string, text string) error {
	// parse program
	program, err := parsetext.Parse(text)
	if err != nil {
		i.presenter.ShowError(err)
		return err
	}

	client := paiza.NewClient()
	result, err := client.Request(language, program)
	if err != nil {
		i.presenter.ShowError(err)
		return err
	}

	i.presenter.ShowResult(buildMessage.MakeMessage(result))

	return nil
}

func (i *Interactor) Challenge() {
}

func (i *Interactor) Kick() {
}