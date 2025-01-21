package file

import (
	"encoding/json"
	"os"
	"strings"
)

func Read(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func ReadString(path string) (string, error) {
	bytes, err := Read(path)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func ReadLines(path string) ([]string, error) {
	content, err := ReadString(path)

	if err != nil {
		return nil, err
	}
	return strings.Split(content, "\n"), nil
}

func ReadJson[T any](path string, unmarshaler ...json.Unmarshaler) (T, error) {
	var result T

	data, err := Read(path)
	if err != nil {
		return result, err
	}

	if len(unmarshaler) == 0 {
		err = json.Unmarshal(data, &result)
	} else {
		err = unmarshaler[0].UnmarshalJSON(data)
		result = unmarshaler[0].(T)
	}

	return result, err
}
