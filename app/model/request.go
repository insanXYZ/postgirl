package model

import (
	"net/http"
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

var (
	Methods = []string{
		http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodHead, http.MethodOptions,
	}
	BodyOptions = []string{"none", "form-data", "x-www-form-urlencoded", "json", "xml"}
)

type Request struct {
	Method    int
	Url       string
	Attribute Attribute
}

type Attribute struct {
	Params  map[string][]string
	Headers map[string]string
	Body    any
}
