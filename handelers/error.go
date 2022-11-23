package handelers

import (
	"errors"
	// "gorm.io/gorm/logger"

)

var (
	// ErrRecordNotFound record not found error
    ErrRecordNotFound = errors.New("record not found")
    // ErrRecordExists record already exists error
    ErrRecordExists = errors.New("record exists")
    // ErrPrimaryKeyRequired primary keys required
	ErrPrimaryKeyRequired = errors.New("primary key required")
)