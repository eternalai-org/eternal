package utils

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Err struct {
	err          error
	SystemError  string `json:"error"`
	ErrorCode    int32  `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	HttpCode     int    `json:"-"`
}

func (e *Err) Error() string {
	return fmt.Sprintf("%v ,code=%d, message=%s", e.err, e.ErrorCode, e.ErrorMessage)
}

func FromErr(err error) *Err {
	if e, ok := err.(*Err); ok {
		return e
	}
	er := &Err{
		err:         err,
		ErrorCode:   50_000,
		SystemError: err.Error(),
		HttpCode:    fiber.StatusInternalServerError,
	}
	if e, ok := err.(*fiber.Error); ok {
		er.HttpCode = e.Code
		er.ErrorMessage = e.Message
		return er
	}

	if s, ok := status.FromError(err); ok {
		switch s.Code() {
		case codes.NotFound:
			er.HttpCode = fiber.StatusNotFound
		case codes.InvalidArgument:
			er.HttpCode = fiber.StatusBadRequest
		case codes.DeadlineExceeded:
			er.HttpCode = fiber.StatusRequestTimeout
		case codes.FailedPrecondition:
			er.HttpCode = fiber.StatusPreconditionFailed
		case codes.AlreadyExists:
			er.HttpCode = fiber.StatusConflict
		}
		er.ErrorMessage = s.Message()
	}

	if er.ErrorMessage == "" {
		er.ErrorMessage = er.SystemError
	}
	return er
}
