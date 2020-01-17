package errors

import (
	"errors"
)

var NoRecord = errors.New("record not found")
var DuplicateRecord = errors.New("record already exists")
