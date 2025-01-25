package pointer

import (
	"testing"

	"github.com/matryer/is"
)

func TestResolveWithValue(t *testing.T) {
	i := is.New(t)
	val := 42
	ptr := &val
	result := Resolve(ptr)
	i.Equal(result, val)
}

func TestResolveOrDefaultWithValue(t *testing.T) {
	i := is.New(t)
	val := 42
	result := ResolveOrDefault(&val)
	i.Equal(result, val)
}

func TestResolveOrDefaultWithNil(t *testing.T) {
	i := is.New(t)
	var nilPtr *int
	result := ResolveOrDefault(nilPtr)
	i.Equal(result, 0)
}

func TestResolveOrWithValue(t *testing.T) {
	i := is.New(t)
	val := 42
	defaultVal := 100
	result := ResolveOr(&val, defaultVal)
	i.Equal(result, val)
}

func TestResolveOrWithNil(t *testing.T) {
	i := is.New(t)
	var nilPtr *int
	defaultVal := 100
	result := ResolveOr(nilPtr, defaultVal)
	i.Equal(result, defaultVal)
}
