package utils

import "github.com/go-errors/errors"

var (
	ErrNotSupportedDayOfTheWeek = errors.New("`day_of_the_field` must be between 1-7")
)
