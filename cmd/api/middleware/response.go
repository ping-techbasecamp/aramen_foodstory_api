package middleware

import "aramen-api/pkg/errs"

func unauthorizedResponse(err error, infoKey string) errs.Error {
	return errs.Error{
		Code:      errs.ErrCodeUnauthorized,
		Message:   err.Error(),
		Info:      map[string]interface{}{infoKey: err.Error()},
		Timestamp: errs.TimestampNow(),
	}
}
