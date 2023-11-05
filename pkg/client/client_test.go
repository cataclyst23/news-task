package client

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"

	"github.com/cataclyst23/homework/pkg/utils"
	"golang.org/x/net/html"
)

const templateDir = "../../tmpl"

func init() {
	utils.Templates = template.Must(template.ParseFiles(templateDir + "/display.html"))
}

func TestResponseHandler(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(ResponseHandler))

	resp, err := http.Get(server.URL)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 but got %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	println(string(body))

	// Test the returned HTML page generated by the template package
	// The contents could also be checked
	_, err = html.Parse(resp.Body)
	if err != nil {
		t.Errorf("Returned HTML page is not valid: %s", err.Error())
	}
}
