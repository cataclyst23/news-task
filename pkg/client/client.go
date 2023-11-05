package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/cataclyst23/homework/pkg/utils"
)

func renderTemplate(w http.ResponseWriter, tmpl string, p *utils.Page) {
	err := utils.Templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func MakeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := utils.ValidPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r)
	}
}

func ResponseHandler(w http.ResponseWriter, r *http.Request) {

	responseObject := &utils.ResponseJson{}

	if err := json.Unmarshal(makeRequest(), responseObject); err != nil {
		fmt.Printf("Error occured during JSON marshaling: %s\n", err.Error())
	}

	fmt.Println(utils.PrettyPrint(responseObject))

	p := &utils.Page{
		Title: utils.PageTitle,
		Body:  responseObject,
	}

	renderTemplate(w, "display", p)
}

func makeRequest() []byte {
	url := url.URL{}

	url.Scheme = utils.Scheme
	url.Host = utils.ApiUrl
	url.Path = utils.ApiUrlPath

	q := url.Query()
	q.Add("appid", utils.AppId)

	url.RawQuery = q.Encode()

	req, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		fmt.Printf("Error building the HTTP request: %s\n", err.Error())
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error making the HTTP request: %s", err.Error())
	}

	if resp.StatusCode <= 200 && resp.StatusCode < 400 {
		fmt.Printf("Response code: %d\n", resp.StatusCode)

	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Could not read response body: %s\n", err)
	}

	return data
}