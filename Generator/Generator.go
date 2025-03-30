package Generator

import "encoding/json"

type Generator interface {
	GetCombinations() []string
}

func SliceToString(slice []any) (string, error) {
	// Marshal the slice into JSON bytes.
	b, err := json.Marshal(slice)
	if err != nil {
		return "", err
	}
	// Convert the bytes to a string.
	return string(b), nil
}
