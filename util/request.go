package util

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"postgirl/model"
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

	splitUrl := strings.Split(u, ":")
	if len(splitUrl) == 0 {
		return nil, errors.New(model.ErrUrlRequired)
	}

	if splitUrl[0] != "http" && splitUrl[0] != "https" {
		return nil, errors.New(model.ErrMissingProtocol)
	}

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

		if strings.Contains(i, ":file") {
			splits := strings.Split(i, ":")
			if len(splits) == 0 {
				return nil, "", errors.New(model.ErrInvalidFormatFileFormData)
			}

			fieldName := splits[0]

			files, ok := v.([]string)
			if !ok {
				return nil, "", errors.New(model.ErrInvalidFormatFileFormData)
			}

			for _, filePath := range files {
				fileName := filepath.Base(filePath)

				fileWriter, err := multipart.CreateFormFile(fieldName, fileName)
				if err != nil {
					return nil, "", errors.New(model.ErrCreateFormDataBody)
				}

				file, err := os.Open(filePath)
				if err != nil {
					return nil, "", fmt.Errorf("%s %s", model.ErrReadFileFormData, filePath)
				}

				_, err = io.Copy(fileWriter, file)
				if err != nil {
					return nil, "", errors.New(model.ErrCreateFormDataBody)
				}
			}

		} else {
			writer, err := multipart.CreateFormField(i)
			if err != nil {
				return nil, "", err
			}

			_, err = writer.Write([]byte(v.(string)))
			if err != nil {
				return nil, "", err
			}
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

func CreateReaderXmlType(body model.BodyMap) (io.Reader, error) {
	m := mxj.Map(body)
	b, err := m.Xml()
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(b), nil

}
