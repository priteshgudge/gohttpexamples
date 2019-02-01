package domain

import "errors"

var DomainErrorNotFound = errors.New("Object Not Found")

var DomainErrorProcessingDB = errors.New("Error in Processing DB query")
