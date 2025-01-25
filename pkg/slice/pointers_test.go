package slice

import (
	"testing"

	"github.com/matryer/is"
)

func TestToPtrWithInts(t *testing.T) {
	i := is.New(t)
	input := []int{1, 2, 3}
	result := ToPtr(input)

	i.True(len(result) == len(input))
	for idx, ptr := range result {
		i.True(ptr != nil)
		i.Equal(*ptr, input[idx])
	}
}

func TestFromPtrWithInts(t *testing.T) {
	i := is.New(t)
	one, two, three := 1, 2, 3
	input := []*int{&one, &two, &three}
	result := FromPtr(input)

	expected := []int{1, 2, 3}
	i.Equal(result, expected)
}
