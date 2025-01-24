package slice

import (
	"testing"
)

func TestMap(t *testing.T) {
	input := []int{1, 2, 3}
	result := Map(input, func(v int) string {
		return string(rune('A' + v - 1))
	})
	expected := []string{"A", "B", "C"}

	for i, v := range result {
		if v != expected[i] {
			t.Errorf("Map at index %d got %v, want %v", i, v, expected[i])
		}
	}
}

func TestMapIndexed(t *testing.T) {
	input := []int{1, 2, 3}
	result := MapIndexed(input, func(i int, v int) string {
		return string(rune('A' + i))
	})
	expected := []string{"A", "B", "C"}

	for i, v := range result {
		if v != expected[i] {
			t.Errorf("MapIndexed at index %d got %v, want %v", i, v, expected[i])
		}
	}
}

func TestMapToPtr(t *testing.T) {
	input := []int{1, 2, 3}
	result := MapToPtr(input, func(v int) string {
		return string(rune('A' + v - 1))
	})

	for i, ptr := range result {
		if ptr == nil {
			t.Errorf("MapToPtr at index %d got nil", i)
			continue
		}
		expected := string(rune('A' + i))
		if *ptr != expected {
			t.Errorf("MapToPtr at index %d got %v, want %v", i, *ptr, expected)
		}
	}
}

func TestMapNonPtrToPtr(t *testing.T) {
	input := []int{1, 2, 3}
	result := MapNonPtrToPtr(input, func(v *int) *string {
		s := string(rune('A' + *v - 1))
		return &s
	})

	for i, ptr := range result {
		if ptr == nil {
			t.Errorf("MapNonPtrToPtr at index %d got nil", i)
			continue
		}
		expected := string(rune('A' + i))
		if *ptr != expected {
			t.Errorf("MapNonPtrToPtr at index %d got %v, want %v", i, *ptr, expected)
		}
	}
}
