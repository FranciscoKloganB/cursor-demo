package request

import (
	"sync"

	"encore.dev/beta/errs"
	"encore.dev/rlog"
	"github.com/go-playground/validator/v10"
)

var (
	validate      *validator.Validate
	validatorOnce sync.Once
)

// ValidationErrorDetail represents an api core request validation error
type ValidationErrorDetail struct {
	Error string `json:"error"`
	Field string `json:"field"`
	Path  string `json:"path_segment"`
	Query string `json:"query_parameter"`
}

// ValidationErrors implements errs.ErrDetails interface
type ValidationErrors []ValidationErrorDetail

// ErrDetails is a marker interface required by Encore so that it knows the type for reporting error details.
func (d ValidationErrors) ErrDetails() {}

// GetValidator returns a singleton instance of validator.Validate
func GetValidator() *validator.Validate {
	validatorOnce.Do(func() {
		validate = validator.New()
	})

	return validate
}

// ValidateStruct validates fields from a struct using go-playground/validator/v10
func ValidateStruct(s interface{}) error {
	return GetValidator().Struct(s)
}

// ConvertValidatorError converts go-playground/validator/v10 ValidationError to an Encore error
func ConvertValidatorError(err error) error {
	if err == nil {
		rlog.Debug("Skipping request validation no error received")

		return nil
	}

	// err is a go-playground/validator err and needs to be transformed
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		errDetails := make(ValidationErrors, 0, len(validationErrors))
		for _, e := range validationErrors {
			errDetails = append(errDetails, ValidationErrorDetail{
				Field: e.Field(),
				Error: e.Tag(),
			})
		}

		return &errs.Error{
			Code:    errs.InvalidArgument,
			Message: "validation_failed",
			Details: errDetails,
		}
	}

	// err is already Encore errs.Error however this is not expected to happen
	if _, ok := err.(*errs.Error); ok {
		return errs.WrapCode(err, errs.Unknown, "internal")
	}

	// err is neither an Encore error errs.Error nor a go-playground/validator error, set 500
	return errs.WrapCode(err, errs.Internal, "internal")
}
