package model

import (
	"io"
	"net/http"
)

type ParamsMap = map[string][]string
type HeadersMap = map[string]string
type BodyMap = map[string]any

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
	NONE                  = "none"
	FORM_DATA             = "form-data"
	X_WWW_FORM_URLENCODED = "application/x-www-form-urlencoded"
	JSON                  = "application/json"
	XML                   = "application/xml"
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
	Method    int
	Url       string
	Attribute *Attribute
}

type Attribute struct {
	Params   ParamsMap
	Headers  HeadersMap
	BodyType string
	Body     io.Reader
}
