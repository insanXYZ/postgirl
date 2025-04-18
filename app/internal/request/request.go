package request

import (
	"fmt"
	"io"
	"net/url"
)

type RequestConfig struct {
	Method, Url string
	Body        io.Reader
	Headers     map[string]string
	Params      map[string][]string
}

func NewRequest(r *RequestConfig) error {

	_, err := parseUrl(r.Url)
	if err != nil {
		return err
	}

	return nil

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
