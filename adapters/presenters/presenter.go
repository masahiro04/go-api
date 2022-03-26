package presenters

import (
	"net/http"

	"github.com/hashicorp/go-multierror"

	"github.com/gin-gonic/gin"
)

type Presenter interface {
	Present() error
}

type GinPresenter struct {
	ctx *gin.Context
	err error
}

func (presenter *GinPresenter) setErrors(err error) {
	if merr, ok := err.(*multierror.Error); ok {
		for _, e := range merr.Errors {
			presenter.err = presenter.ctx.Error(e)
		}
		return
	}
	presenter.err = presenter.ctx.Error(err)
}

func New(ctx *gin.Context) *GinPresenter {
	return &GinPresenter{ctx, ctx.Err()}
}

func (presenter *GinPresenter) Unauthorized(err error) {
	presenter.ctx.Status(http.StatusUnauthorized)
	presenter.setErrors(err)
}

func (presenter *GinPresenter) UnprocessableEntity(err error) {
	presenter.ctx.Status(http.StatusUnprocessableEntity)
	presenter.setErrors(err)
}

func (presenter *GinPresenter) NotFound(err error) {
	presenter.ctx.Status(http.StatusNotFound)
	presenter.setErrors(err)
}

func (presenter *GinPresenter) BadRequest(err error) {
	presenter.ctx.Status(http.StatusBadRequest)
	presenter.setErrors(err)
}

func (presenter *GinPresenter) InternalServerError(err error) {
	presenter.ctx.Status(http.StatusInternalServerError)
	presenter.setErrors(err)
}

func (presenter *GinPresenter) Created(value interface{}) {
	presenter.ctx.JSON(http.StatusCreated, value)
}

func (presenter *GinPresenter) StatusOK(value interface{}) {
	presenter.ctx.JSON(http.StatusOK, value)
}

func (presenter *GinPresenter) Err() error {
	return presenter.err
}
