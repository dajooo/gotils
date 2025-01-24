package encoding

import (
	"testing"
)

type testStruct struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestUnmarshalJSON(t *testing.T) {
	data := []byte(`{"name":"test","age":25}`)

	result, err := UnmarshalJSON[testStruct](data)
	if err != nil {
		t.Errorf("UnmarshalJSON failed: %v", err)
	}
	if result.Name != "test" || result.Age != 25 {
		t.Errorf("UnmarshalJSON got %+v, want {Name:test Age:25}", result)
	}
}

func TestMarshalJSON(t *testing.T) {
	data := testStruct{Name: "test", Age: 25}

	result, err := MarshalJSON(data)
	if err != nil {
		t.Errorf("MarshalJSON failed: %v", err)
	}

	expected := `{"name":"test","age":25}`
	if string(result) != expected {
		t.Errorf("MarshalJSON got %s, want %s", result, expected)
	}
}

func TestMustUnmarshalJSON(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for invalid JSON")
		}
	}()

	MustUnmarshalJSON[testStruct]([]byte(`invalid json`))
}

func TestMustMarshalJSON(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for invalid value")
		}
	}()

	// Create a value that can't be marshaled
	ch := make(chan int)
	MustMarshalJSON(ch)
}
