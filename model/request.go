package model

import (
	"io"
	"net/http"
)

type (
	ParamsMap  = map[string][]string
	HeadersMap = map[string]string
	BodyMap    = map[string]any
)

const (
	GET = iota
	POST
	PUT
	PATCH
	DELETE
	HEAD
	OPTIONS
)

const (
	ENCODE_GZIP    = "gzip"
	ENCODE_BR      = "br"
	ENCODE_DEFLATE = "deflate"
)

const (
	NONE                  = "none"
	FORM_DATA             = "form-data"
	X_WWW_FORM_URLENCODED = "x-www-form-urlencoded"
	JSON                  = "json"
	XML                   = "xml"
)

const (
	CONTENT_TYPE_X_WWW_FORM_URLENCODED = "application/x-www-form-urlencoded"
	CONTENT_TYPE_JSON                  = "application/json"
	CONTENT_TYPE_XML                   = "application/xml"
)

var (
	Methods = []string{
		http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodHead, http.MethodOptions,
	}
	BodyOptions = []string{
		NONE, FORM_DATA, X_WWW_FORM_URLENCODED, JSON, XML,
	}
)

type Request struct {
	Method    int        `json:"method"`
	Url       string     `json:"url"`
	Attribute *Attribute `json:"attribute"`
}

type Attribute struct {
	Params     ParamsMap  `json:"params_map"`
	Headers    HeadersMap `json:"headers_map"`
	BodyType   string     `json:"body_type"`
	BodyString string     `json:"body_string"`
	Body       io.Reader  `json:"-"`
}
