package file

import (
	"encoding/json"
	"os"
	"strings"
)

func Write(path string, data []byte) error {
	return os.WriteFile(path, data, 0644)
}

func WriteString(path, data string) error {
	return Write(path, []byte(data))
}

func WriteLines(path string, lines []string) error {
	return WriteString(path, strings.Join(lines, "\n"))
}

func WriteJson[T any](path string, data T, indent ...string) error {
	var bytes []byte
	var err error

	if len(indent) > 0 {
		bytes, err = json.MarshalIndent(data, "", indent[0])
	} else {
		bytes, err = json.Marshal(data)
	}
	if err != nil {
		return err
	}

	return Write(path, bytes)
}
