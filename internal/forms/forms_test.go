package forms

import (
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/something", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("Got false when should have got true")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/something", nil)
	form := New(r.PostForm)

	form.Required("a", "n", "d")
	if form.Valid() {
		t.Error("form shows valid when required fields missing")
	}

	postData := url.Values{}

	postData.Add("a", "a")
	postData.Add("b", "a")
	postData.Add("c", "a")

	// r = httptest.NewRequest("POST", "/something", nil)
	// r.PostForm = postData

	form = New(postData)

	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("form shows invalid when required fields present")
	}

}

func TestForm_MinLength(t *testing.T) {

	postData := url.Values{}
	form := New(postData)
	form.MinLength("a", 3)
	if form.Valid() {
		t.Error("Valid even though field not in form")
	}

	postData.Add("a", "123")

	form = New(postData)

	if !form.MinLength("a", 3) {
		t.Error("Min length test failed when correct length given")
	}

	if !form.Valid() {
		t.Error("Should not have Min length test error when correct length given")
	}

	form.MinLength("a", 4)
	if form.Valid() {
		t.Error("Passes min length even though short field given")
	}
}

func TestForm_Has(t *testing.T) {
	postData := url.Values{}
	form := New(postData)

	if form.Has("a") {
		t.Error("Reports has field but it does not")
	}

	postData.Add("b", "asd")
	if !form.Has("b") {
		t.Error("Reports does not have field but it does")
	}
}

func TestForm_IsEmail(t *testing.T) {
	postData := url.Values{}
	postData.Add("a", "b@c.com")
	postData.Add("b", "b")
	form := New(postData)

	if !form.IsEmail("a") {
		t.Error("say email invalid but it is")
	}

	if !form.Valid() {
		t.Error("gives error message on valid email")
	}

	if form.IsEmail("b") {
		t.Error("says email valid but it is not")
	}

	if form.Valid() {
		t.Error("does not give email faulty message")
	}

}
