package utils

import (
	"strconv"
)

func ConvertStringToInt(s string) (uint64, error) {
	return strconv.ParseUint(s, 10, 32)
}
