package file

import (
	"os"
	"testing"
)

func TestReadWrite(t *testing.T) {
	tempFile := "test.txt"
	defer os.Remove(tempFile)

	// Test Write
	testData := []byte("test data")
	err := Write(tempFile, testData)
	if err != nil {
		t.Fatalf("Write failed: %v", err)
	}

	// Test Read
	data, err := Read(tempFile)
	if err != nil {
		t.Fatalf("Read failed: %v", err)
	}
	if string(data) != string(testData) {
		t.Errorf("Read got %s, want %s", data, testData)
	}
}

func TestReadWriteString(t *testing.T) {
	tempFile := "test_string.txt"
	defer os.Remove(tempFile)

	testStr := "test string"
	err := WriteString(tempFile, testStr)
	if err != nil {
		t.Fatalf("WriteString failed: %v", err)
	}

	result, err := ReadString(tempFile)
	if err != nil {
		t.Fatalf("ReadString failed: %v", err)
	}
	if result != testStr {
		t.Errorf("ReadString got %s, want %s", result, testStr)
	}
}

func TestReadWriteLines(t *testing.T) {
	tempFile := "test_lines.txt"
	defer os.Remove(tempFile)

	lines := []string{"line1", "line2", "line3"}
	err := WriteLines(tempFile, lines)
	if err != nil {
		t.Fatalf("WriteLines failed: %v", err)
	}

	result, err := ReadLines(tempFile)
	if err != nil {
		t.Fatalf("ReadLines failed: %v", err)
	}
	if len(result) != len(lines) {
		t.Errorf("ReadLines got %d lines, want %d", len(result), len(lines))
	}
}

func TestReadWriteJson(t *testing.T) {
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
	if err != nil {
		t.Fatalf("WriteJson failed: %v", err)
	}

	var result struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	result, err = ReadJson[struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}](tempFile)
	if err != nil {
		t.Fatalf("ReadJson failed: %v", err)
	}
	if result.Name != data.Name || result.Age != data.Age {
		t.Errorf("ReadJson got %+v, want %+v", result, data)
	}
}
