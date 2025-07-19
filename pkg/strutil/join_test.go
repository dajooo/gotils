package strutil

import (
	"fmt"
	"testing"

	"github.com/matryer/is"
)

func TestJoinStream(t *testing.T) {
	is := is.New(t)

	t.Run("join stream with multiple strings", func(t *testing.T) {
		is := is.New(t)
		ch := make(chan string, 3)
		ch <- "hello"
		ch <- "world"
		ch <- "test"
		close(ch)

		result := JoinStream(", ", ch)
		expected := "hello, world, test"
		is.Equal(result, expected)
	})

	t.Run("join stream with single string", func(t *testing.T) {
		is := is.New(t)
		ch := make(chan string, 1)
		ch <- "single"
		close(ch)

		result := JoinStream(", ", ch)
		expected := "single"
		is.Equal(result, expected)
	})

	t.Run("join stream with empty channel", func(t *testing.T) {
		is := is.New(t)
		ch := make(chan string)
		close(ch)

		result := JoinStream(", ", ch)
		expected := ""
		is.Equal(result, expected)
	})

	t.Run("join stream with different separator", func(t *testing.T) {
		is := is.New(t)
		ch := make(chan string, 4)
		ch <- "a"
		ch <- "b"
		ch <- "c"
		ch <- "d"
		close(ch)

		result := JoinStream(" | ", ch)
		expected := "a | b | c | d"
		is.Equal(result, expected)
	})

	t.Run("join stream with empty separator", func(t *testing.T) {
		is := is.New(t)
		ch := make(chan string, 3)
		ch <- "hello"
		ch <- "world"
		ch <- "test"
		close(ch)

		result := JoinStream("", ch)
		expected := "helloworldtest"
		is.Equal(result, expected)
	})

	t.Run("join stream with empty strings", func(t *testing.T) {
		is := is.New(t)
		ch := make(chan string, 3)
		ch <- ""
		ch <- "middle"
		ch <- ""
		close(ch)

		result := JoinStream("-", ch)
		expected := "-middle-"
		is.Equal(result, expected)
	})

	t.Run("join stream with whitespace strings", func(t *testing.T) {
		is := is.New(t)
		ch := make(chan string, 3)
		ch <- "  "
		ch <- "text"
		ch <- "\t"
		close(ch)

		result := JoinStream(",", ch)
		expected := "  ,text,\t"
		is.Equal(result, expected)
	})
}

func BenchmarkJoinStream(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ch := make(chan string, 100)
		for j := 0; j < 100; j++ {
			ch <- "benchmark"
		}
		close(ch)
		JoinStream(", ", ch)
	}
}

func ExampleJoinStream() {
	ch := make(chan string, 3)
	ch <- "apple"
	ch <- "banana"
	ch <- "cherry"
	close(ch)

	result := JoinStream(", ", ch)
	fmt.Print(result)
	// Output: apple, banana, cherry
}
