package gokore

import "fmt"

const (
	ServerError = iota
	ClientError = iota
)

type gokoreError struct {
	error
	ErrorType int
}

type GokoreError *gokoreError

func NewServerError(message string) GokoreError {
	return &gokoreError{
		error:     fmt.Errorf(message),
		ErrorType: ServerError,
	}
}

func NewClientError(message string) GokoreError {
	return &gokoreError{
		error:     fmt.Errorf(message),
		ErrorType: ClientError,
	}
}
