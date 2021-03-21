package auth

import (
	"regexp"

	errors "github.com/joaosoft/errors"
	"github.com/joaosoft/validator"
)

const (
	regexForMissingParms = `%\+?[a-z]`
)

func initValidator() {
	validator.
		SetValidateAll(true).
		SetErrorCodeHandler(ErrorHandler)
}

var errs = map[string]*errors.Error{
	"InvalidBodyParameter": ErrorInvalidBodyParameter,
}

var ErrorHandler = func(context *validator.ValidatorContext, validationData *validator.ValidationData) error {
	if err, ok := errs[validationData.ErrorData.Code]; ok {
		var regx = regexp.MustCompile(regexForMissingParms)
		matches := regx.FindAllStringIndex(err.Message, -1)

		if len(matches) > 0 {

			if len(validationData.ErrorData.Arguments) < len(matches) {
				validationData.ErrorData.Arguments = append(validationData.ErrorData.Arguments, validationData.Name)
			}

			err = err.Format(validationData.ErrorData.Arguments...)
		}

		return err
	}
	return nil
}
