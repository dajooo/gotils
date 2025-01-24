package pointer

import "testing"

func TestOf(t *testing.T) {
	val := 42
	ptr := Of(val)
	if *ptr != val {
		t.Errorf("Of got %v, want %v", *ptr, val)
	}
}

func TestOfOk(t *testing.T) {
	tests := []struct {
		name    string
		value   int
		ok      bool
		wantNil bool
	}{
		{"ok true", 42, true, false},
		{"ok false", 42, false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := OfOk(tt.value, tt.ok)
			if tt.wantNil && result != nil {
				t.Error("Expected nil result")
			}
			if !tt.wantNil && result == nil {
				t.Error("Expected non-nil result")
			}
		})
	}
}
