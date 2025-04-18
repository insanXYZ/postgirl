package model

import "net/http"

var Methods = []string{
	http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodHead, http.MethodOptions,
}

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
