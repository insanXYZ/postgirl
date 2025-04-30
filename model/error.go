package model

const (
	ErrStartApp = "error running app, you can report it to github issue on https://github.com/insanXYZ/postgirl"

	// Format
	ErrInvalidFormatParams   = "invalid params format"
	ErrInvalidFormatHeaders  = "invalid headers format"
	ErrInvalidFormatUrl      = "invalid url format"
	ErrInvalidFormatBody     = "invalid body format"
	ErrInvalidFormatFileBody = "invalid file body format"

	// InputUrlBar
	ErrUrlRequired      = "url required"
	ErrProtocolRequired = "protocol required"

	// Attribute
	ErrReadResponseBody   = "error read response body"
	ErrReadResponseHeader = "error read response header"
	ErrReadDirectory      = "error read directory"

	//Cache
	ErrSaveCache = "error save request to cache"
	ErrReadCache = "error read cache"

	// Sidebar
	ErrCreateRequest     = "error create request"
	ErrNameRequired      = "name required"
	ErrDuplicateName     = "duplicate name"
	ErrFieldnameRequired = "fieldname required"

	// form-data
	ErrInvalidFieldnameFile = "invalid fieldname file"
	ErrReadFileFormData     = "error read file"
	ErrCreateFormDataBody   = "error create form-data body"
)
