package utils

import (
	"encoding/json"
	"regexp"
	"text/template"
)

type Page struct {
	Title string
	Body  *ResponseJson
}

type ResponseJson struct {
	Appnews struct {
		Appid     int `json:"appid"`
		Newsitems []struct {
			Gid           string   `json:"gid"`
			Title         string   `json:"title"`
			URL           string   `json:"url"`
			IsExternalURL bool     `json:"is_external_url"`
			Author        string   `json:"author"`
			Contents      string   `json:"contents"`
			Feedlabel     string   `json:"feedlabel"`
			Date          int      `json:"date"`
			Feedname      string   `json:"feedname"`
			FeedType      int      `json:"feed_type"`
			Appid         int      `json:"appid"`
			Tags          []string `json:"tags,omitempty"`
		} `json:"newsitems"`
		Count int `json:"count"`
	} `json:"appnews"`
}

const (
	TemplateDir = "../../tmpl"
	DataDir     = "data"

	Scheme     = "https"
	ApiUrl     = "api.steampowered.com"
	ApiUrlPath = "/ISteamNews/GetNewsForApp/v2/"
	// The ID of Counter Strike
	AppId = "730"

	PageTitle = "The Steam API response"
)

var (
	Templates = template.Must(template.ParseFiles(TemplateDir + "/display.html"))
	ValidPath = regexp.MustCompile("^/$")
)

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
