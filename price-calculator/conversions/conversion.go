package conversion

import (
	"errors"
	"strconv" //provides utility functions for converting text/string to other types
)

func StringsToFloats(inputStrings []string) ([]float64, error) {
	floats := make([]float64, len(inputStrings))
	for stringIndex, stringVal := range inputStrings {
		number, err := strconv.ParseFloat(stringVal, 64)
		if err != nil {
			return nil, errors.New("someting went wrong whilst converting string to float")
		}
		floats[stringIndex] = number
	}
	return floats, nil
}
