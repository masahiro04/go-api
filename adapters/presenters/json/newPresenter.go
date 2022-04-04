package json

import (
	"go-api/adapters/presenters"
	"go-api/domains/models"
	"go-api/usecases"
)

const dateLayout = "2006-01-02T15:04:05.999Z"

type ResponsePresenter struct {
	Presenter *presenters.GinPresenter
}

func NewPresenter(presenter *presenters.GinPresenter) usecases.PresenterRepository {
	return &ResponsePresenter{Presenter: presenter}
}

func (presenter ResponsePresenter) Raise(errorKind models.ErrorKinds, err error) {
	switch errorKind {
	case models.BadRequest:
		presenter.Presenter.BadRequest(err)
	case models.UnprocessableEntity:
		presenter.Presenter.UnprocessableEntity(err)
	case models.NotFound:
		presenter.Presenter.NotFound(err)
	case models.InternalServerError:
		presenter.Presenter.InternalServerError(err)
	case models.Unauthorized:
		presenter.Presenter.Unauthorized(err)
	default:
		presenter.Presenter.BadRequest(err)
	}
}

func (presenter ResponsePresenter) Present() error {
	return presenter.Presenter.Err()
}
