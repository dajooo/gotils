package slice

import "testing"

func TestToPtr(t *testing.T) {
	input := []int{1, 2, 3}
	result := ToPtr(input)

	for i, ptr := range result {
		if ptr == nil {
			t.Errorf("ToPtr at index %d got nil", i)
			continue
		}
		if *ptr != input[i] {
			t.Errorf("ToPtr at index %d got %v, want %v", i, *ptr, input[i])
		}
	}
}

func TestFromPtr(t *testing.T) {
	one, two, three := 1, 2, 3
	input := []*int{&one, &two, &three}
	result := FromPtr(input)

	expected := []int{1, 2, 3}
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("FromPtr at index %d got %v, want %v", i, v, expected[i])
		}
	}
}
