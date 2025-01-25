package pointer

import (
	"testing"

	"github.com/matryer/is"
)

func TestOfWithInt(t *testing.T) {
	i := is.New(t)
	val := 42
	ptr := Of(val)
	i.True(ptr != nil)
	i.Equal(*ptr, val)
}

func TestOfOkWithTrue(t *testing.T) {
	i := is.New(t)
	val := 42
	result := OfOk(val, true)
	i.True(result != nil)
	i.Equal(*result, val)
}

func TestOfOkWithFalse(t *testing.T) {
	i := is.New(t)
	result := OfOk(42, false)
	i.Equal(result, nil)
}
