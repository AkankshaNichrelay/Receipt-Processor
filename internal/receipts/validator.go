package receipts

import (
	"fmt"
	"regexp"

	structTagValidator "github.com/go-playground/validator/v10"
)

const (
	RegexFloats      = "^\\d+\\.\\d{2}$"
	RegexDescription = "^[\\w\\s\\-]+$"
	RegexStrings     = "^\\S+$"
)

func (r *Receipts) validateReceipt(receipt *Receipt) error {
	if err := r.validateRequiredFields(receipt); err != nil {
		return err
	}
	// TODO: Can be added if regex needs to be enforced.
	// if err := r.validateRegex(receipt); err != nil {
	// 	return err
	// }
	return nil
}

func (r *Receipts) validateRequiredFields(receipt *Receipt) error {
	validate := structTagValidator.New()

	err := validate.Struct(receipt)
	if err != nil {
		validationErrs, ok := err.(structTagValidator.ValidationErrors)
		if ok {
			for _, validationErr := range validationErrs {
				if validationErr.Tag() == "required" {
					return fmt.Errorf("field missing %s expected type %s", validationErr.Field(), validationErr.Type())
				}
				if validationErr.Tag() == "min" {
					return validationErr
				}
			}
		}
	}
	return nil
}

// validateRegex to match the api.yml specification. Not checking for now due to conflicting example in ReadME
func (r *Receipts) validateRegex(receipt *Receipt) error {
	reFloats := regexp.MustCompile(RegexFloats)
	reDesc := regexp.MustCompile(RegexDescription)
	reStrings := regexp.MustCompile(RegexStrings)

	if !reStrings.MatchString(receipt.Retailer) {
		return fmt.Errorf("did not match required regex pattern: Field: Retailer, Pattern %s", reStrings)
	}
	if !reFloats.MatchString(receipt.Total) {
		return fmt.Errorf("did not match required regex pattern: Field: Total, Pattern %s", reFloats)
	}
	for _, item := range receipt.Items {
		if !reDesc.MatchString(item.ShortDescription) {
			return fmt.Errorf("did not match required regex pattern: Field: ShortDescription, Pattern %s", reDesc)
		}
		if !reFloats.MatchString(item.Price) {
			return fmt.Errorf("did not match required regex pattern. Field: Price, Pattern %s", reFloats)
		}
	}
	return nil
}
