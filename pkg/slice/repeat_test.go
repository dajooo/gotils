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
		is.Equal(result, "hello")
	})

	t.Run("repeat integer", func(t *testing.T) {
		is := is.New(t)
		result := Repeat(3, 42)
		is.Equal(result, 42)
	})

	t.Run("repeat boolean", func(t *testing.T) {
		is := is.New(t)
		result := Repeat(2, true)
		is.Equal(result, true)
	})

	t.Run("repeat with count 1", func(t *testing.T) {
		is := is.New(t)
		result := Repeat(1, "single")
		is.Equal(result, "single")
	})

	t.Run("repeat with count 0", func(t *testing.T) {
		is := is.New(t)
		defer func() {
			if r := recover(); r != nil {
				is.True(true)
			} else {
				is.Fail()
			}
		}()
		Repeat(0, "empty")
	})

	t.Run("repeat struct", func(t *testing.T) {
		is := is.New(t)
		type Person struct {
			Name string
			Age  int
		}
		person := Person{Name: "John", Age: 30}
		result := Repeat(3, person)
		is.Equal(result, person)
	})

	t.Run("repeat slice", func(t *testing.T) {
		is := is.New(t)
		slice := []int{1, 2, 3}
		result := Repeat(2, slice)
		is.Equal(result, slice)
	})
}

func TestRepeatFunc(t *testing.T) {
	is := is.New(t)

	t.Run("repeat func with index", func(t *testing.T) {
		is := is.New(t)
		result := RepeatFunc(5, func(index int) int {
			return index * 2
		})
		is.Equal(result, 0)
	})

	t.Run("repeat func with string formatting", func(t *testing.T) {
		is := is.New(t)
		result := RepeatFunc(3, func(index int) string {
			return "item-" + string(rune('0'+index))
		})
		is.Equal(result, "item-0")
	})

	t.Run("repeat func with boolean logic", func(t *testing.T) {
		is := is.New(t)
		result := RepeatFunc(4, func(index int) bool {
			return index%2 == 0
		})
		is.Equal(result, true)
	})

	t.Run("repeat func with count 1", func(t *testing.T) {
		is := is.New(t)
		result := RepeatFunc(1, func(index int) string {
			return "only"
		})
		is.Equal(result, "only")
	})

	t.Run("repeat func with count 0", func(t *testing.T) {
		is := is.New(t)

		defer func() {
			if r := recover(); r != nil {

				is.True(true)
			} else {
				is.Fail()
			}
		}()
		RepeatFunc(0, func(index int) string {
			return "never called"
		})
	})

	t.Run("repeat func with struct creation", func(t *testing.T) {
		is := is.New(t)
		type Item struct {
			ID    int
			Value string
		}
		result := RepeatFunc(3, func(index int) Item {
			return Item{ID: index, Value: "value-" + string(rune('A'+index))}
		})
		expected := Item{ID: 0, Value: "value-A"}
		is.Equal(result, expected)
	})

	t.Run("repeat func with closure", func(t *testing.T) {
		is := is.New(t)
		multiplier := 10
		result := RepeatFunc(2, func(index int) int {
			return index * multiplier
		})
		is.Equal(result, 0)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat(100, "benchmark")
	}
}

func BenchmarkRepeatFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatFunc(100, func(index int) int {
			return index * 2
		})
	}
}

func ExampleRepeat() {
	result := Repeat(3, "hello")
	fmt.Println(result)

}

func ExampleRepeatFunc() {
	result := RepeatFunc(3, func(index int) int {
		return index * 10
	})
	fmt.Println(result)

}
