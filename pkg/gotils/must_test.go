package gotils

import (
	"errors"
	"testing"

	"github.com/matryer/is"
)

func TestMustWithNoError(t *testing.T) {
	i := is.New(t)
	value := "success"
	result := Must(value, nil)
	i.Equal(result, value)
}

func TestMustWithError(t *testing.T) {
	i := is.New(t)
	err := errors.New("test error")

	defer func() {
		r := recover()
		i.True(r != nil)
		i.Equal(r, err)
	}()

	Must("", err)
}

func TestMustOkWithValidValue(t *testing.T) {
	i := is.New(t)
	value := 42
	result := MustOk(value, true)
	i.Equal(result, value)
}

func TestMustOkWithInvalidValue(t *testing.T) {
	i := is.New(t)

	defer func() {
		r := recover()
		i.True(r != nil)
	}()

	MustOk(0, false)
}

func TestMustWithInt(t *testing.T) {
	i := is.New(t)
	result := Must(42, nil)
	i.Equal(result, 42)
}

func TestMustWithStruct(t *testing.T) {
	i := is.New(t)
	type testStruct struct {
		value string
	}
	expected := testStruct{value: "test"}
	result := Must(expected, nil)
	i.Equal(result, expected)
}

func TestMustWithSlice(t *testing.T) {
	i := is.New(t)
	expected := []int{1, 2, 3}
	result := Must(expected, nil)
	i.Equal(result, expected)
}

func TestMustOkWithString(t *testing.T) {
	i := is.New(t)
	result := MustOk("test", true)
	i.Equal(result, "test")
}

func TestMustOkWithMap(t *testing.T) {
	i := is.New(t)
	expected := map[string]int{"test": 42}
	result := MustOk(expected, true)
	i.Equal(result, expected)
}
