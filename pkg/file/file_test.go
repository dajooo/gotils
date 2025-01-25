package file

import (
	"os"
	"testing"

	"github.com/matryer/is"
)

func TestWriteAndRead(t *testing.T) {
	i := is.New(t)
	tempFile := "test.txt"
	defer os.Remove(tempFile)

	testData := []byte("test data")
	err := Write(tempFile, testData)
	i.NoErr(err)

	data, err := Read(tempFile)
	i.NoErr(err)
	i.Equal(string(data), string(testData))
}

func TestWriteAndReadString(t *testing.T) {
	i := is.New(t)
	tempFile := "test_string.txt"
	defer os.Remove(tempFile)

	testStr := "test string"
	err := WriteString(tempFile, testStr)
	i.NoErr(err)

	result, err := ReadString(tempFile)
	i.NoErr(err)
	i.Equal(result, testStr)
}

func TestWriteAndReadLines(t *testing.T) {
	i := is.New(t)
	tempFile := "test_lines.txt"
	defer os.Remove(tempFile)

	lines := []string{"line1", "line2", "line3"}
	err := WriteLines(tempFile, lines)
	i.NoErr(err)

	result, err := ReadLines(tempFile)
	i.NoErr(err)
	i.Equal(len(result), len(lines))
	i.Equal(result, lines)
}

func TestWriteAndReadJson(t *testing.T) {
	i := is.New(t)
	tempFile := "test.json"
	defer os.Remove(tempFile)

	data := struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{
		Name: "test",
		Age:  25,
	}

	err := WriteJson(tempFile, data, "  ")
	i.NoErr(err)

	result, err := ReadJson[struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}](tempFile)
	i.NoErr(err)
	i.Equal(result.Name, data.Name)
	i.Equal(result.Age, data.Age)
}
