package signedpay_api

import (
	"errors"
)

type PayError interface {
	error
}

var (
	ErrGeneralDecline = errors.New("general decline")
	ErrAuthorization  = errors.New("authorization failed")
	ErrInvalidData    = errors.New("invalid data")
	ErrUnknown        = errors.New("unknown error")
)

func apiError(a Error) PayError {
	switch a.GetCode() {
	case "":
		return nil
	case "0.01":
		return ErrGeneralDecline
	case "1.01":
		return ErrAuthorization
	case "2.01":
		return ErrInvalidData
	default:
		return ErrUnknown
	}
}
