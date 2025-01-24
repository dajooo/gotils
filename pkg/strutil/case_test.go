package strutil

import (
	"reflect"
	"strings"
	"sync"
	"testing"
)

func TestStringCase_String(t *testing.T) {
	tests := []struct {
		name     string
		sc       StringCase
		expected string
	}{
		{
			name:     "Empty CamelCase",
			sc:       StringCase{Kind: CamelCase, Parts: []string{}},
			expected: "",
		},
		{
			name:     "Basic CamelCase",
			sc:       StringCase{Kind: CamelCase, Parts: []string{"hello", "world"}},
			expected: "helloWorld",
		},
		{
			name:     "Basic PascalCase",
			sc:       StringCase{Kind: PascalCase, Parts: []string{"hello", "world"}},
			expected: "HelloWorld",
		},
		{
			name:     "Basic SnakeCase",
			sc:       StringCase{Kind: SnakeCase, Parts: []string{"hello", "world"}},
			expected: "hello_world",
		},
		{
			name:     "Basic KebabCase",
			sc:       StringCase{Kind: KebabCase, Parts: []string{"hello", "world"}},
			expected: "hello-world",
		},
		{
			name:     "Basic ScreamingSnakeCase",
			sc:       StringCase{Kind: ScreamingSnakeCase, Parts: []string{"hello", "world"}},
			expected: "HELLO_WORLD",
		},
		{
			name:     "Unknown Case",
			sc:       StringCase{Kind: "unknown", Parts: []string{"hello", "world"}},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.sc.String(); got != tt.expected {
				t.Errorf("StringCase.String() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestDetectCase(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected StringCaseKind
	}{
		{"Empty string", "", CamelCase},
		{"Snake case", "hello_world", SnakeCase},
		{"Kebab case", "hello-world", KebabCase},
		{"Pascal case", "HelloWorld", PascalCase},
		{"Camel case", "helloWorld", CamelCase},
		{"Multiple underscores", "hello_world_test", SnakeCase},
		{"Multiple hyphens", "hello-world-test", KebabCase},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DetectCase(tt.input); got != tt.expected {
				t.Errorf("DetectCase() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestParseCase(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected StringCase
	}{
		{
			name:  "Snake case",
			input: "hello_world_test",
			expected: StringCase{
				Kind:  SnakeCase,
				Parts: []string{"hello", "world", "test"},
			},
		},
		{
			name:  "Kebab case",
			input: "hello-world-test",
			expected: StringCase{
				Kind:  KebabCase,
				Parts: []string{"hello", "world", "test"},
			},
		},
		{
			name:  "Pascal case",
			input: "HelloWorldTest",
			expected: StringCase{
				Kind:  PascalCase,
				Parts: []string{"hello", "world", "test"},
			},
		},
		{
			name:  "Camel case",
			input: "helloWorldTest",
			expected: StringCase{
				Kind:  CamelCase,
				Parts: []string{"hello", "world", "test"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParseCase(tt.input)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("ParseCase() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestCapitalize(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Empty string", "", ""},
		{"Single letter", "a", "A"},
		{"Word", "hello", "Hello"},
		{"Already capitalized", "Hello", "Hello"},
		{"Mixed case", "hELLO", "HELLO"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := capitalize(tt.input); got != tt.expected {
				t.Errorf("capitalize() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestStringCase_Conversions(t *testing.T) {
	sc := StringCase{
		Kind:  SnakeCase,
		Parts: []string{"hello", "world", "test"},
	}

	tests := []struct {
		name     string
		convert  func() string
		expected string
	}{
		{"ToCamelCase", sc.ToCamelCase, "helloWorldTest"},
		{"ToPascalCase", sc.ToPascalCase, "HelloWorldTest"},
		{"ToSnakeCase", sc.ToSnakeCase, "hello_world_test"},
		{"ToKebabCase", sc.ToKebabCase, "hello-world-test"},
		{"ToScreamingSnakeCase", sc.ToScreamingSnakeCase, "HELLO_WORLD_TEST"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.convert(); got != tt.expected {
				t.Errorf("%s() = %v, want %v", tt.name, got, tt.expected)
			}
		})
	}
}

func TestSplitCamelCase(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{"Empty string", "", nil},
		{"Single word", "hello", []string{"hello"}},
		{"Camel case", "helloWorld", []string{"hello", "World"}},
		{"Pascal case", "HelloWorld", []string{"Hello", "World"}},
		{"Multiple words", "helloWorldTest", []string{"hello", "World", "Test"}},
		{"Consecutive capitals", "HTMLParser", []string{"HTML", "Parser"}},
		{"Mixed case", "myXMLParser", []string{"my", "XML", "Parser"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := splitCamelCase(tt.input)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("splitCamelCase() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestScreamingKebabCase(t *testing.T) {
	tests := []struct {
		name     string
		parts    []string
		expected string
	}{
		{"Empty parts", []string{}, ""},
		{"Single word", []string{"hello"}, "HELLO"},
		{"Multiple words", []string{"hello", "world"}, "HELLO-WORLD"},
		{"Already uppercase", []string{"HELLO", "WORLD"}, "HELLO-WORLD"},
		{"Mixed case", []string{"Hello", "World"}, "HELLO-WORLD"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := screamingKebabCase(tt.parts); got != tt.expected {
				t.Errorf("screamingKebabCase() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected StringCase
	}{
		{
			name:  "Empty string",
			input: "",
			expected: StringCase{
				Kind:  CamelCase,
				Parts: nil,
			},
		},
		{
			name:  "Single underscore",
			input: "_",
			expected: StringCase{
				Kind:  SnakeCase,
				Parts: nil,
			},
		},
		{
			name:  "Single hyphen",
			input: "-",
			expected: StringCase{
				Kind:  KebabCase,
				Parts: nil,
			},
		},
		{
			name:  "Multiple delimiters",
			input: "hello__world--test",
			expected: StringCase{
				Kind:  SnakeCase,
				Parts: []string{"hello", "world", "test"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParseCase(tt.input)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("ParseCase()\ngot: %#v\nwant: %#v", got, tt.expected)
			}
		})
	}
}

func TestNegativeCases(t *testing.T) {
	tests := []struct {
		name          string
		testFunc      func() string
		expectedEmpty bool
	}{
		{
			name: "Nil parts in CamelCase",
			testFunc: func() string {
				return StringCase{Kind: CamelCase, Parts: nil}.String()
			},
			expectedEmpty: true,
		},
		{
			name: "Nil parts in PascalCase",
			testFunc: func() string {
				return StringCase{Kind: PascalCase, Parts: nil}.String()
			},
			expectedEmpty: true,
		},
		{
			name: "Invalid StringCaseKind",
			testFunc: func() string {
				return StringCase{Kind: "invalid_kind", Parts: []string{"test"}}.String()
			},
			expectedEmpty: true,
		},
		{
			name: "Empty parts with invalid kind",
			testFunc: func() string {
				return StringCase{Kind: "invalid", Parts: []string{}}.String()
			},
			expectedEmpty: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.testFunc()
			if tt.expectedEmpty && result != "" {
				t.Errorf("Expected empty string, got %q", result)
			}
		})
	}
}

func TestInvalidInputs(t *testing.T) {
	tests := []struct {
		name string
		test func(t *testing.T)
	}{
		{
			name: "Multiple consecutive delimiters in snake case",
			test: func(t *testing.T) {
				input := "hello___world"
				result := ParseCase(input)
				expected := []string{"hello", "world"}
				if !reflect.DeepEqual(result.Parts, expected) {
					t.Errorf("Expected parts %v, got %v", expected, result.Parts)
				}
			},
		},
		{
			name: "Mixed delimiters",
			test: func(t *testing.T) {
				input := "hello_world-test"
				result := ParseCase(input)
				expected := []string{"hello", "world", "test"}
				if !reflect.DeepEqual(result.Parts, expected) {
					t.Errorf("Expected parts %v, got %v", expected, result.Parts)
				}
			},
		},
		{
			name: "Invalid UTF-8 handling",
			test: func(t *testing.T) {
				input := string([]byte{0xFF, 0xFE, 0xFD})
				result := ParseCase(input)
				if len(result.Parts) != 1 {
					t.Errorf("Expected single part for invalid UTF-8, got %d parts", len(result.Parts))
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.test(t)
		})
	}
}
func TestMalformedInputs(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		testFunc func(string) string
	}{
		{
			name:  "Malformed camel case",
			input: "helloWORLDTest",
			testFunc: func(s string) string {
				return ParseCase(s).ToCamelCase()
			},
		},
		{
			name:  "Malformed snake case",
			input: "hello__WORLD__test_",
			testFunc: func(s string) string {
				return ParseCase(s).ToSnakeCase()
			},
		},
		{
			name:  "Malformed kebab case",
			input: "hello--WORLD--test-",
			testFunc: func(s string) string {
				return ParseCase(s).ToKebabCase()
			},
		},
		{
			name:  "Mixed case with numbers",
			input: "hello123World456",
			testFunc: func(s string) string {
				return ParseCase(s).ToPascalCase()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("Test panicked with input %q: %v", tt.input, r)
				}
			}()

			result := tt.testFunc(tt.input)
			if result == "" {
				t.Errorf("Got empty result for input %q", tt.input)
			}
		})
	}
}

func TestBoundaryConditions(t *testing.T) {
	tests := []struct {
		name string
		test func(t *testing.T)
	}{
		{
			name: "Very long string",
			test: func(t *testing.T) {
				input := strings.Repeat("a", 1000000)
				defer func() {
					if r := recover(); r != nil {
						t.Errorf("Panicked on very long string: %v", r)
					}
				}()
				result := ParseCase(input)
				if result.String() == "" {
					t.Error("Got empty result for very long string")
				}
			},
		},
		{
			name: "Zero-width spaces",
			test: func(t *testing.T) {
				input := "hello\u200Bworld"
				result := ParseCase(input)
				if result.String() == "" {
					t.Error("Got empty result for string with zero-width spaces")
				}
			},
		},
		{
			name: "Control characters",
			test: func(t *testing.T) {
				input := "hello\u0000world\u0007test"
				result := ParseCase(input)
				if result.String() == "" {
					t.Error("Got empty result for string with control characters")
				}
			},
		},
		{
			name: "Unicode characters",
			test: func(t *testing.T) {
				input := "héllöWórld"
				result := ParseCase(input)
				if result.String() == "" {
					t.Error("Got empty result for string with Unicode characters")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.test(t)
		})
	}
}

func TestConcurrentAccess(t *testing.T) {
	sc := StringCase{
		Kind:  CamelCase,
		Parts: []string{"hello", "world", "test"},
	}

	const goroutines = 100
	const iterations = 100

	var wg sync.WaitGroup
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				// Access all conversion methods concurrently
				_ = sc.ToCamelCase()
				_ = sc.ToPascalCase()
				_ = sc.ToSnakeCase()
				_ = sc.ToKebabCase()
				_ = sc.ToScreamingSnakeCase()
			}
		}()
	}

	wg.Wait()
}
