package model

const (
	ErrStartApp                   = "error running app, you can report it to github issue on https://github.com/insanXYZ/postgirl"
	ErrInvalidFormatParams        = "invalid params format"
	ErrInvalidFormatHeaders       = "invalid headers format"
	ErrInvalidFormatUrl           = "invalid url format"
	ErrInvalidFormatBody          = "invalid body format"
	ErrUrlRequired                = "url required"
	ErrMissingProtocol            = "procol required"
	ErrReadResponseBody           = "error read response body"
	ErrReadHeader                 = "error read response header"
	ErrSaveCache                  = "error save request to cache"
	ErrReadDir                    = "error read directory"
	ErrInvalidFormatFileFormData  = "invalid file body format"
	ErrReadFileFormData           = "error read file"
	ErrCreateFormDataBody         = "error create form-data body"
	ErrCreateRequest              = "error create request"
	ErrCreateRequestNameRequired  = "name required"
	ErrCreateRequestDuplicateName = "duplicate name"
)
