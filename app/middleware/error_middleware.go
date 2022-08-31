package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/dchest/uniuri"
	"github.com/go-errors/errors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/rizface/golang-api-template/app/entity/responseentity"
	"github.com/rizface/golang-api-template/app/errorgroup"
	"github.com/rizface/golang-api-template/system/logger"
	"github.com/sirupsen/logrus"
)

func ErrorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := logger.CreateErrorLogger()
		defer func() {
			err := recover()
			if err != nil {
				var errStruct = responseentity.Error{
					Id: uniuri.New(),
				}

				if group, ok := err.(errorgroup.Error); ok {
					errStruct.Code = group.Code
					errStruct.Message = group.Message
				} else if validatorError, ok := err.(validation.Errors); ok {
					errStruct.Code = errorgroup.BAD_REQUEST.Code
					errStruct.Message = validatorError.Error()
				} else {
					errStruct.Code = errorgroup.InternalServerError.Code
					errStruct.Message = errorgroup.InternalServerError.Message
				}

				var convertedErrorMessage string
				stackTrace := errors.Wrap(err, 1).ErrorStack()

				if convertedError, ok := err.(error); ok {
					convertedErrorMessage = convertedError.Error()
				} else {
					convertedErrorMessage = err.(string)
				}

				log.WithFields(logrus.Fields{
					"id":    errStruct.Id,
					"error": convertedErrorMessage,
					"trace": stackTrace,
				}).Error(errStruct.Message)

				w.WriteHeader(errStruct.Code)
				json.NewEncoder(w).Encode(errStruct)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
