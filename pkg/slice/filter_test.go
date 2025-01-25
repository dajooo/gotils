package slice

import (
	"testing"

	"github.com/matryer/is"
)

func TestFilterWithInts(t *testing.T) {
	i := is.New(t)
	input := []int{1, 2, 3, 4, 5}
	result := Filter(input, func(v int) bool {
		return v%2 == 0
	})
	i.Equal(result, []int{2, 4})
}

func TestFilterWithStrings(t *testing.T) {
	i := is.New(t)
	input := []string{"a", "b", "c", "aa", "bb"}
	result := Filter(input, func(v string) bool {
		return len(v) > 1
	})
	i.Equal(result, []string{"aa", "bb"})
}

func TestFilterWithEmptySlice(t *testing.T) {
	i := is.New(t)
	var input []int
	result := Filter(input, func(v int) bool {
		return v > 0
	})
	i.Equal(len(result), 0)
}

func TestFilterInstanceOfString(t *testing.T) {
	i := is.New(t)
	input := []any{"hello", 42, "world", true}
	result := FilterInstanceOf[string](input)
	i.Equal(result, []string{"hello", "world"})
}

func TestFilterInstanceOfInt(t *testing.T) {
	i := is.New(t)
	input := []any{1, "hello", 42, true, 3}
	result := FilterInstanceOf[int](input)
	i.Equal(result, []int{1, 42, 3})
}

func TestFilterInstanceOfWithEmptySlice(t *testing.T) {
	i := is.New(t)
	var input []any
	result := FilterInstanceOf[string](input)
	i.Equal(len(result), 0)
}

func TestFilterNotNilWithPointers(t *testing.T) {
	i := is.New(t)
	str1, str2 := "hello", "world"
	input := []*string{&str1, nil, &str2, nil}
	result := FilterNotNil(input)
	i.Equal(len(result), 2)
	i.Equal(result[0], "hello")
	i.Equal(result[1], "world")
}

func TestFilterNotNilWithEmptySlice(t *testing.T) {
	i := is.New(t)
	var input []*string
	result := FilterNotNil(input)
	i.Equal(len(result), 0)
}

func TestFilterNotNilWithAllNil(t *testing.T) {
	i := is.New(t)
	input := []*string{nil, nil, nil}
	result := FilterNotNil(input)
	i.Equal(len(result), 0)
}

func TestFilterNotNilWithAllNotNil(t *testing.T) {
	i := is.New(t)
	str1, str2, str3 := "a", "b", "c"
	input := []*string{&str1, &str2, &str3}
	result := FilterNotNil(input)
	i.Equal(len(result), 3)
	i.Equal(result[0], "a")
	i.Equal(result[1], "b")
	i.Equal(result[2], "c")
}
