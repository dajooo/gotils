package pointer

import "testing"

func TestResolve(t *testing.T) {
	val := 42
	ptr := &val
	result := Resolve(ptr)
	if result != val {
		t.Errorf("Resolve got %v, want %v", result, val)
	}
}

func TestResolveOrDefault(t *testing.T) {
	val := 42
	var nilPtr *int

	result1 := ResolveOrDefault(&val)
	if result1 != val {
		t.Errorf("ResolveOrDefault got %v, want %v", result1, val)
	}

	result2 := ResolveOrDefault(nilPtr)
	if result2 != 0 {
		t.Errorf("ResolveOrDefault got %v, want 0", result2)
	}
}

func TestResolveOr(t *testing.T) {
	val := 42
	defaultVal := 100
	var nilPtr *int

	result1 := ResolveOr(&val, defaultVal)
	if result1 != val {
		t.Errorf("ResolveOr got %v, want %v", result1, val)
	}

	result2 := ResolveOr(nilPtr, defaultVal)
	if result2 != defaultVal {
		t.Errorf("ResolveOr got %v, want %v", result2, defaultVal)
	}
}
