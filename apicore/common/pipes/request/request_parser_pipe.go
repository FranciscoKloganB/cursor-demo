package request

import "encore.dev/beta/errs"

// ParseRequest validates a struct using go-playground/validator/v10 and returns an Encore error errs.Error
func ParseRequest(s interface{}) error {
	if err := ValidateStruct(s); err != nil {
		return ConvertValidatorError(err)
	}

	if err := TransformStruct(s); err != nil {
		return errs.WrapCode(err, errs.Internal, "transform_failed")
	}

	if err := ScrubStruct(s); err != nil {
		return errs.WrapCode(err, errs.Internal, "scrub_failed")
	}

	return nil
}
