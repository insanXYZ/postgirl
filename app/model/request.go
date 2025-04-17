package model

type Request struct {
	Label     string
	Method    string
	Url       string
	Attribute Attribute
}

type Attribute struct {
	Params  map[string]string
	Headers map[string]string
	Body    any
}
