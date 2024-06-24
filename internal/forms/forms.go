package forms

import (
	"net/http"
	"net/url"
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

// Has checks if field is in form Post and not empty
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x == "" {
		f.Errors.Add(field, "This field cannot be blank")
		return false
	}
	return true
}

// Valid return true if there are no errors on the form, otherwise false
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}
