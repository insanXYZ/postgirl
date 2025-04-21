package util

import (
	"fmt"
	"net/http"
	"net/url"
	"postgirl/app/model"
)

func NewRequest(r *model.Request) (*http.Response, error) {
	req, err := http.NewRequest(model.Methods[r.Method], r.Url, r.Attribute.Body)
	if err != nil {
		return nil, err
	}

	for i, v := range r.Attribute.Headers {
		req.Header.Add(i, v)
	}
	client := http.Client{}
	return client.Do(req)
}

type Url struct {
	Params   map[string][]string
	CleanUrl string
}

func ParseUrl(u string) (*Url, error) {
	res := Url{}

	parse, err := url.Parse(u)
	if err != nil {
		return nil, err
	}

	qparams, err := url.ParseQuery(parse.RawQuery)
	if err != nil {
		return nil, err
	}

	res.CleanUrl = fmt.Sprintf("%s://%s%s", parse.Scheme, parse.Host, parse.Path)
	res.Params = qparams

	return &res, nil
}
