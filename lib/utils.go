package lib

import (
	"errors"
	"strconv"
)

func ParseRowNum(rowNum string) (int, error) {
	if rowNum == "" {
		return 0, errors.New("invalid row number")
	}

	return strconv.Atoi(rowNum)
}
