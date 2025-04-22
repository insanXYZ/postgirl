package util

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"postgirl/app/model"
	"strings"

	"github.com/clbanning/mxj/v2"
)

func NewRequest(r *model.Request) (*http.Response, error) {
	req, err := http.NewRequest(model.Methods[r.Method], r.Url, r.Attribute.Body)
	if err != nil {
		return nil, err
	}

	for i, v := range r.Attribute.Headers {
		req.Header.Add(i, v)
	}

	if r.Attribute.BodyType != model.NONE {
		req.Header.Set("Content-Type", r.Attribute.BodyType)
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

func CreateReaderFormDataType(body model.BodyMap) (io.Reader, string, error) {
	buf := &bytes.Buffer{}
	multipart := multipart.NewWriter(buf)

	for i, v := range body {
		writer, err := multipart.CreateFormField(i)
		if err != nil {
			return nil, "", err
		}

		_, err = writer.Write([]byte(v.(string)))
		if err != nil {
			return nil, "", err
		}
	}

	err := multipart.Close()
	if err != nil {
		return nil, "", err
	}

	return buf, multipart.FormDataContentType(), nil
}

func CreateReaderXWWWFormUrlencodedType(body model.BodyMap) io.Reader {
	data := url.Values{}

	for i, v := range body {
		data.Set(i, v.(string))
	}

	return strings.NewReader(data.Encode())
}

func CreateReaderJsonType(body model.BodyMap) (io.Reader, error) {
	b, err := JsonMarshal(body)
	return bytes.NewReader(b), err
}

func CreateReaderXmlType(body string) (io.Reader, error) {
	m, err := mxj.NewMapXml([]byte(body))
	if err != nil {
		return nil, err
	}

	b, err := m.Xml()
	return bytes.NewReader(b), err
}
