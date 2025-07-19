package slice

import (
	"fmt"
	"testing"

	"github.com/matryer/is"
)

func TestRepeat(t *testing.T) {
	is := is.New(t)

	t.Run("repeat string", func(t *testing.T) {
		is := is.New(t)
		result := Repeat(5, "hello")
		expected := []string{"hello", "hello", "hello", "hello", "hello"}
		is.Equal(result, expected)
	})

	t.Run("repeat integer", func(t *testing.T) {
		is := is.New(t)
		result := Repeat(3, 42)
		expected := []int{42, 42, 42}
		is.Equal(result, expected)
	})

	t.Run("repeat boolean", func(t *testing.T) {
		is := is.New(t)
		result := Repeat(2, true)
		expected := []bool{true, true}
		is.Equal(result, expected)
	})

	t.Run("repeat with count 1", func(t *testing.T) {
		is := is.New(t)
		result := Repeat(1, "single")
		expected := []string{"single"}
		is.Equal(result, expected)
	})

	t.Run("repeat with count 0", func(t *testing.T) {
		is := is.New(t)
		result := Repeat(0, "empty")
		expected := []string{}
		is.Equal(result, expected)
	})

	t.Run("repeat struct", func(t *testing.T) {
		is := is.New(t)
		type Person struct {
			Name string
			Age  int
		}
		person := Person{Name: "John", Age: 30}
		result := Repeat(3, person)
		expected := []Person{person, person, person}
		is.Equal(result, expected)
	})

	t.Run("repeat slice", func(t *testing.T) {
		is := is.New(t)
		slice := []int{1, 2, 3}
		result := Repeat(2, slice)
		expected := [][]int{slice, slice}
		is.Equal(result, expected)
	})
}

func TestRepeatFunc(t *testing.T) {
	is := is.New(t)

	t.Run("repeat func with index", func(t *testing.T) {
		is := is.New(t)
		result := RepeatFunc(5, func(index int) int {
			return index * 2
		})
		expected := []int{0, 2, 4, 6, 8}
		is.Equal(result, expected)
	})

	t.Run("repeat func with string formatting", func(t *testing.T) {
		is := is.New(t)
		result := RepeatFunc(3, func(index int) string {
			return fmt.Sprintf("item-%d", index)
		})
		expected := []string{"item-0", "item-1", "item-2"}
		is.Equal(result, expected)
	})

	t.Run("repeat func with boolean logic", func(t *testing.T) {
		is := is.New(t)
		result := RepeatFunc(4, func(index int) bool {
			return index%2 == 0
		})
		expected := []bool{true, false, true, false}
		is.Equal(result, expected)
	})

	t.Run("repeat func with count 1", func(t *testing.T) {
		is := is.New(t)
		result := RepeatFunc(1, func(index int) string {
			return "only"
		})
		expected := []string{"only"}
		is.Equal(result, expected)
	})

	t.Run("repeat func with count 0", func(t *testing.T) {
		is := is.New(t)
		result := RepeatFunc(0, func(index int) string {
			return "never called"
		})
		expected := []string{}
		is.Equal(result, expected)
	})

	t.Run("repeat func with struct creation", func(t *testing.T) {
		is := is.New(t)
		type Item struct {
			ID    int
			Value string
		}
		result := RepeatFunc(3, func(index int) Item {
			return Item{ID: index, Value: fmt.Sprintf("value-%c", 'A'+index)}
		})
		expected := []Item{
			{ID: 0, Value: "value-A"},
			{ID: 1, Value: "value-B"},
			{ID: 2, Value: "value-C"},
		}
		is.Equal(result, expected)
	})

	t.Run("repeat func with closure", func(t *testing.T) {
		is := is.New(t)
		multiplier := 10
		result := RepeatFunc(2, func(index int) int {
			return index * multiplier
		})
		expected := []int{0, 10}
		is.Equal(result, expected)
	})
}

func TestRepeatStream(t *testing.T) {
	is := is.New(t)

	t.Run("repeat stream string", func(t *testing.T) {
		is := is.New(t)
		ch := RepeatStream(3, "hello")
		var result []string
		for val := range ch {
			result = append(result, val)
		}
		expected := []string{"hello", "hello", "hello"}
		is.Equal(result, expected)
	})

	t.Run("repeat stream integer", func(t *testing.T) {
		is := is.New(t)
		ch := RepeatStream(2, 42)
		var result []int
		for val := range ch {
			result = append(result, val)
		}
		expected := []int{42, 42}
		is.Equal(result, expected)
	})

	t.Run("repeat stream with count 0", func(t *testing.T) {
		is := is.New(t)
		ch := RepeatStream(0, "empty")
		var result []string
		for val := range ch {
			result = append(result, val)
		}
		expected := []string{}
		is.Equal(result, expected)
	})

	t.Run("repeat stream with count 1", func(t *testing.T) {
		is := is.New(t)
		ch := RepeatStream(1, true)
		var result []bool
		for val := range ch {
			result = append(result, val)
		}
		expected := []bool{true}
		is.Equal(result, expected)
	})
}

func TestRepeatStreamFunc(t *testing.T) {
	is := is.New(t)

	t.Run("repeat stream func with index", func(t *testing.T) {
		is := is.New(t)
		ch := RepeatStreamFunc(4, func(index int) int {
			return index * 3
		})
		var result []int
		for val := range ch {
			result = append(result, val)
		}
		expected := []int{0, 3, 6, 9}
		is.Equal(result, expected)
	})

	t.Run("repeat stream func with string formatting", func(t *testing.T) {
		is := is.New(t)
		ch := RepeatStreamFunc(3, func(index int) string {
			return fmt.Sprintf("item-%d", index)
		})
		var result []string
		for val := range ch {
			result = append(result, val)
		}
		expected := []string{"item-0", "item-1", "item-2"}
		is.Equal(result, expected)
	})

	t.Run("repeat stream func with count 0", func(t *testing.T) {
		is := is.New(t)
		ch := RepeatStreamFunc(0, func(index int) string {
			return "never called"
		})
		var result []string
		for val := range ch {
			result = append(result, val)
		}
		expected := []string{}
		is.Equal(result, expected)
	})

	t.Run("repeat stream func with boolean logic", func(t *testing.T) {
		is := is.New(t)
		ch := RepeatStreamFunc(3, func(index int) bool {
			return index%2 == 0
		})
		var result []bool
		for val := range ch {
			result = append(result, val)
		}
		expected := []bool{true, false, true}
		is.Equal(result, expected)
	})
}
