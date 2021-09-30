package domain

type ErrorKinds int

const (
	BadRequest ErrorKinds = iota
	UnprocessableEntity
	NotFound
	InternalServerError
	Unauthorized
)
