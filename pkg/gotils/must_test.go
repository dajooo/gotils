// must_test.go
package gotils

import (
	"errors"
	"testing"
)

func TestMust(t *testing.T) {
	tests := []struct {
		name      string
		value     string
		err       error
		wantPanic bool
	}{
		{
			name:      "no error",
			value:     "success",
			err:       nil,
			wantPanic: false,
		},
		{
			name:      "with error",
			value:     "",
			err:       errors.New("test error"),
			wantPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if tt.wantPanic && r == nil {
					t.Error("Expected panic but got none")
				}
				if !tt.wantPanic && r != nil {
					t.Errorf("Unexpected panic: %v", r)
				}
			}()

			result := Must(tt.value, tt.err)
			if tt.err == nil && result != tt.value {
				t.Errorf("Must() = %v, want %v", result, tt.value)
			}
		})
	}
}

func TestMustOk(t *testing.T) {
	tests := []struct {
		name      string
		value     int
		ok        bool
		wantPanic bool
	}{
		{
			name:      "ok true",
			value:     42,
			ok:        true,
			wantPanic: false,
		},
		{
			name:      "ok false",
			value:     0,
			ok:        false,
			wantPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if tt.wantPanic && r == nil {
					t.Error("Expected panic but got none")
				}
				if !tt.wantPanic && r != nil {
					t.Errorf("Unexpected panic: %v", r)
				}
			}()

			result := MustOk(tt.value, tt.ok)
			if tt.ok && result != tt.value {
				t.Errorf("MustOk() = %v, want %v", result, tt.value)
			}
		})
	}
}

// TestMustWithDifferentTypes tests Must with various types
func TestMustWithDifferentTypes(t *testing.T) {
	t.Run("with int", func(t *testing.T) {
		result := Must(42, nil)
		if result != 42 {
			t.Errorf("Must(int) = %v, want 42", result)
		}
	})

	t.Run("with struct", func(t *testing.T) {
		type testStruct struct {
			value string
		}
		expected := testStruct{value: "test"}
		result := Must(expected, nil)
		if result != expected {
			t.Errorf("Must(struct) = %v, want %v", result, expected)
		}
	})

	t.Run("with slice", func(t *testing.T) {
		expected := []int{1, 2, 3}
		result := Must(expected, nil)
		if len(result) != len(expected) {
			t.Errorf("Must(slice) = %v, want %v", result, expected)
		}
	})
}

// TestMustOkWithDifferentTypes tests MustOk with various types
func TestMustOkWithDifferentTypes(t *testing.T) {
	t.Run("with string", func(t *testing.T) {
		result := MustOk("test", true)
		if result != "test" {
			t.Errorf("MustOk(string) = %v, want test", result)
		}
	})

	t.Run("with map", func(t *testing.T) {
		expected := map[string]int{"test": 42}
		result := MustOk(expected, true)
		if result["test"] != 42 {
			t.Errorf("MustOk(map) = %v, want %v", result, expected)
		}
	})
}
