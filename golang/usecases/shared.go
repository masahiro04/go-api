package uc

import (
	"errors"
)

var (
	//ErrAlreadyInUse = errors.New("this username is already in use")
	//ErrUserEmailAlreadyInUsed = errors.New("this email address is already in use")
	errWrongUser    = errors.New("woops, wrong user")
	errWrongCompany = errors.New("woops, wrong company")
	errNotFound     = errors.New("NotFound")
	errNotAuthorize = errors.New("InvalidAuthorization")
)
