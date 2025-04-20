package request

import (
	"fmt"
	"net/http"
	"net/url"
	"postgirl/app/lib"
	"postgirl/app/model"
)

func NewRequest(r *model.Request) (*http.Response, *Url, error) {

	parsed, err := parseUrl(r.Url)
	if err != nil {
		return nil, nil, err

	}

	lib.Tview.Stop()

	fmt.Println(*parsed)

	return nil, parsed, nil

}

type Url struct {
	params   map[string][]string
	cleanUrl string
}

func parseUrl(u string) (*Url, error) {
	res := Url{}

	parse, err := url.Parse(u)
	if err != nil {
		return nil, err
	}

	qparams, err := url.ParseQuery(parse.RawQuery)
	if err != nil {
		return nil, err
	}

	res.cleanUrl = fmt.Sprintf("%s://%s%s", parse.Scheme, parse.Host, parse.Path)
	res.params = qparams

	return &res, nil
}
