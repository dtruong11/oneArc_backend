package model

import "errors"

var ErrorBadRequest = errors.New("Bad request")
var ErrorNotAllowed = errors.New("Not allowed")
var ErrorInternalServer = errors.New("Internal Server Error")
var ErrorNotFound = errors.New("Not Found")
