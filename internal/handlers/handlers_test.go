package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postdata struct {
	key   string
	value string
}

var theTests = []struct {
	name               string
	url                string
	method             string
	params             []postdata
	expectedStatusCode int
}{
	{"home", "/", "GET", []postdata{}, http.StatusOK},
	{"about", "/about", "GET", []postdata{}, http.StatusOK},
	{"bigroom", "/bigroom", "GET", []postdata{}, http.StatusOK},
	{"littleroom", "/littleroom", "GET", []postdata{}, http.StatusOK},
	{"sa", "/search-availability", "GET", []postdata{}, http.StatusOK},
	{"contact", "/contact", "GET", []postdata{}, http.StatusOK},
	{"mr", "/make-reservation", "GET", []postdata{}, http.StatusOK},
	{"about", "/about", "GET", []postdata{}, http.StatusOK},
	{"post-search-avail", "/search-availability", "POST", []postdata{
		{key: "start", value: "2020-01-02"},
		{key: "end", value: "2020-01-02"},
	}, http.StatusOK},
	{"post-search-avail-json", "/search-availability-json", "POST", []postdata{
		{key: "start", value: "2020-01-02"},
		{key: "end", value: "2020-01-02"},
	}, http.StatusOK},
	{"make-reservation", "/make-reservation", "POST", []postdata{
		{key: "first_name", value: "John"},
		{key: "last_name", value: "Smith"},
		{key: "email", value: "mc@here.com"},
		{key: "phone", value: "123456"},
	}, http.StatusOK},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, e := range theTests {
		if e.method == "GET" {
			resp, err := ts.Client().Get(ts.URL + e.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
		} else {
			values := url.Values{}
			for _, x := range e.params {
				values.Add(x.key, x.value)
			}
			resp, err := ts.Client().PostForm(ts.URL+e.url, values)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
		}
	}
}
