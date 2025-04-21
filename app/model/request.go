package model

import (
	"io"
	"net/http"
)

type ParamsMap = map[string][]string
type HeadersMap = map[string]string

const (
	GET = iota
	POST
	PUT
	PATCH
	DELETE
	HEAD
	OPTIONS
)

var (
	Methods = []string{
		http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodHead, http.MethodOptions,
	}
	BodyOptions = []string{"none", "form-data", "x-www-form-urlencoded", "json", "xml"}
)

type Request struct {
	Method    int
	Url       string
	Attribute *Attribute
}

type Attribute struct {
	Params  ParamsMap
	Headers HeadersMap
	Body    io.Reader
}
