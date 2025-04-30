package util

import (
	"bytes"
	"compress/flate"
	"compress/gzip"
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

	"github.com/andybalholm/brotli"
	"github.com/clbanning/mxj/v2"
)

func NewRequest(r *model.Request) (*http.Response, io.Reader, error) {
	var body io.Reader
	var err error

	req, err := http.NewRequest(model.Methods[r.Method], r.Url, r.Attribute.Body)
	if err != nil {
		return nil, nil, err
	}

	for i, v := range r.Attribute.Headers {
		req.Header.Add(i, v)
	}

	if r.Attribute.BodyType != model.NONE {
		req.Header.Set("Content-Type", r.Attribute.BodyType)
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}

	defer res.Body.Close()

	encode := res.Header.Get("Accept-Encoding")

	switch encode {
	case model.ENCODE_GZIP:
		body, err = gzip.NewReader(res.Body)
	case model.ENCODE_BR:
		body = brotli.NewReader(res.Body)
	case model.ENCODE_DEFLATE:
		body = flate.NewReader(res.Body)
	}

	if err != nil {
		return nil, nil, err
	}

	return res, body, nil

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
		return nil, errors.New(model.ErrProtocolRequired)
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
			if len(splits) < 2 {
				return nil, "", errors.New(model.ErrInvalidFieldnameFile)
			}

			fieldName := splits[0]

			files, ok := v.([]any)
			if !ok {
				return nil, "", errors.New(model.ErrInvalidFormatFileBody)
			}

			for _, filePath := range files {
				filePathString, ok := filePath.(string)
				if !ok {
					return nil, "", errors.New(model.ErrInvalidFormatFileBody)
				}

				fileName := filepath.Base(filePathString)

				fileWriter, err := multipart.CreateFormFile(fieldName, fileName)
				if err != nil {
					return nil, "", errors.New(model.ErrCreateFormDataBody)
				}

				file, err := os.Open(filePathString)
				if err != nil {
					return nil, "", fmt.Errorf("%s %s", model.ErrReadFileFormData, filePathString)
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
