package slice

import (
	"testing"

	"github.com/matryer/is"
)

func TestMapIntToString(t *testing.T) {
	i := is.New(t)
	input := []int{1, 2, 3}
	result := Map(input, func(v int) string {
		return string(rune('A' + v - 1))
	})
	expected := []string{"A", "B", "C"}
	i.Equal(result, expected)
}

func TestMapIndexedIntToString(t *testing.T) {
	i := is.New(t)
	input := []int{1, 2, 3}
	result := MapIndexed(input, func(i int, v int) string {
		return string(rune('A' + i))
	})
	expected := []string{"A", "B", "C"}
	i.Equal(result, expected)
}

func TestMapToPtrWithInts(t *testing.T) {
	i := is.New(t)
	input := []int{1, 2, 3}
	result := MapToPtr(input, func(v int) string {
		return string(rune('A' + v - 1))
	})

	i.True(len(result) == 3)
	for idx, ptr := range result {
		i.True(ptr != nil)
		expected := string(rune('A' + idx))
		i.Equal(*ptr, expected)
	}
}

func TestMapNonPtrToPtrWithInts(t *testing.T) {
	i := is.New(t)
	input := []int{1, 2, 3}
	result := MapNonPtrToPtr(input, func(v *int) *string {
		s := string(rune('A' + *v - 1))
		return &s
	})

	i.True(len(result) == 3)
	for idx, ptr := range result {
		i.True(ptr != nil)
		expected := string(rune('A' + idx))
		i.Equal(*ptr, expected)
	}
}
