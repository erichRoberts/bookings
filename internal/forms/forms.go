package forms

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

// Form creates a custom form struct and embeds url values
type Form struct {
	url.Values
	Errors errors
}

// New initialises a Form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Required checks that fields listed as required as not blank
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

// MinLength returns an error if the field is shorter than required
func (f *Form) MinLength(field string, length int) bool {

	x := f.Get(field)

	if len(x) < length {
		f.Errors.Add(field, fmt.Sprintf("Must be at least %d characters long", length))
		return false
	}
	return true
}

// Has checks if field is in form Post and not empty
func (f *Form) Has(field string) bool {
	x := f.Get(field)
	if x == "" {

		return false
	}

	return true
}

// Valid return true if there are no errors on the form, otherwise false
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// IsEmail checks for valid email address format
func (f *Form) IsEmail(field string) bool {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "This must be valid email")
		return false
	}

	return true
}
