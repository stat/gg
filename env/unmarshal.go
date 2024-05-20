package env

import (
	"encoding/json"
	"strconv"
)

//
// Unmarshal
//

func Unmarshal[T any](data []byte) (*T, error) {
	var v T

	// ensure strings are strings

	switch any(v).(type) {
	case string:
		runes := []rune(string(data))

		// check for quotes

		if runes[0] != '"' {
			quoted := strconv.Quote(string(data))
			data = []byte(quoted)
		}
	}

	// unmarshal

	err := json.Unmarshal(data, &v)

	if err != nil {
		return nil, err
	}

	return &v, nil
}
