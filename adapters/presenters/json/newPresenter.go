package json

import (
	"go-api/adapters/presenters"
	"go-api/domains"
	uc "go-api/usecases"
)

const dateLayout = "2006-01-02T15:04:05.999Z"

type ResponsePresenter struct {
	Presenter *presenters.GinPresenter
}

func NewPresenter(presenter *presenters.GinPresenter) uc.Presenter {
	return &ResponsePresenter{Presenter: presenter}
}

func (presenter ResponsePresenter) Raise(errorKind domains.ErrorKinds, err error) {
	switch errorKind {
	case domains.BadRequest:
		presenter.Presenter.BadRequest(err)
	case domains.UnprocessableEntity:
		presenter.Presenter.UnprocessableEntity(err)
	case domains.NotFound:
		presenter.Presenter.NotFound(err)
	case domains.InternalServerError:
		presenter.Presenter.InternalServerError(err)
	case domains.Unauthorized:
		presenter.Presenter.Unauthorized(err)
	default:
		presenter.Presenter.BadRequest(err)
	}
}

func (presenter ResponsePresenter) Present() error {
	return presenter.Presenter.Err()
}
