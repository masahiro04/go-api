package formatter

import (
	presenter "clean_architecture/golang/adapters/json.presenter"
	"clean_architecture/golang/domain"
	"clean_architecture/golang/usecases"
)

const dateLayout = "2006-01-02T15:04:05.999Z"

type ResponsePresenter struct {
	Presenter *presenter.GinPresenter
}

func NewPresenter(presenter *presenter.GinPresenter) uc.Presenter {
	return &ResponsePresenter{Presenter: presenter}
}

func (presenter ResponsePresenter) Raise(errorKind domain.ErrorKinds, err error) {
	switch errorKind {
	case domain.BadRequest:
		presenter.Presenter.BadRequest(err)
	case domain.UnprocessableEntity:
		presenter.Presenter.UnprocessableEntity(err)
	case domain.NotFound:
		presenter.Presenter.NotFound(err)
	case domain.InternalServerError:
		presenter.Presenter.InternalServerError(err)
	case domain.Unauthorized:
		presenter.Presenter.Unauthorized(err)
	default:
		presenter.Presenter.BadRequest(err)
	}
}

func (presenter ResponsePresenter) Present() error {
	return presenter.Presenter.Err()
}
