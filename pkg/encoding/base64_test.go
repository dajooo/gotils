package encoding

import (
	"testing"
)

func TestBase64Encode(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Empty string", "", ""},
		{"Simple string", "hello", "aGVsbG8="},
		{"With special chars", "hello world!", "aGVsbG8gd29ybGQh"},
		{"With padding", "hello", "aGVsbG8="},
		{"Without padding", "hell", "aGVsbA=="},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := B64Encode(tt.input)
			if result != tt.expected {
				t.Errorf("B64Encode() = %v, want %v", result, tt.expected)
			}

			// Test decode
			decoded, err := B64Decode(result)
			if err != nil {
				t.Errorf("B64Decode() error = %v", err)
			}
			if decoded != tt.input {
				t.Errorf("B64Decode() = %v, want %v", decoded, tt.input)
			}
		})
	}
}

func TestBase64URLEncode(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"URL unsafe chars", "hello+world/", "aGVsbG8rd29ybGQv"},
		{"URL safe output", "hello world!", "aGVsbG8gd29ybGQh"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := B64URLEncode(tt.input)
			if result != tt.expected {
				t.Errorf("B64URLEncode() = %v, want %v", result, tt.expected)
			}

			// Test decode
			decoded, err := B64URLDecode(result)
			if err != nil {
				t.Errorf("B64URLDecode() error = %v", err)
			}
			if decoded != tt.input {
				t.Errorf("B64URLDecode() = %v, want %v", decoded, tt.input)
			}
		})
	}
}

func TestBase64BytesConversion(t *testing.T) {
	input := []byte("hello world")

	// Test bytes to string
	encoded := B64EncodeBytes(input)
	decoded, err := B64DecodeToBytes(encoded)
	if err != nil {
		t.Errorf("B64DecodeToBytes() error = %v", err)
	}
	if string(decoded) != string(input) {
		t.Errorf("B64DecodeToBytes() = %v, want %v", string(decoded), string(input))
	}

	// Test bytes to bytes
	encodedBytes := B64EncodeBytesToBytes(input)
	decodedBytes, err := B64DecodeBytesToBytes(encodedBytes)
	if err != nil {
		t.Errorf("B64DecodeBytesToBytes() error = %v", err)
	}
	if string(decodedBytes) != string(input) {
		t.Errorf("B64DecodeBytesToBytes() = %v, want %v", string(decodedBytes), string(input))
	}
}

func TestCustomPadding(t *testing.T) {
	input := "hello"
	expected := "aGVsbG8*"

	result := B64Encode(input, '*')
	if result != expected {
		t.Errorf("B64Encode() with custom padding = %v, want %v", result, expected)
	}
}

func TestMustDecode(t *testing.T) {
	t.Run("valid input", func(t *testing.T) {
		input := B64Encode("hello")
		result := MustB64Decode(input)
		if result != "hello" {
			t.Errorf("MustB64Decode() = %v, want hello", result)
		}
	})

	t.Run("invalid input", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("MustB64Decode() did not panic with invalid input")
			}
		}()
		MustB64Decode("invalid base64!")
	})
}

func TestInvalidInputs(t *testing.T) {
	_, err := B64Decode("invalid base64!")
	if err == nil {
		t.Error("B64Decode() did not return error for invalid input")
	}

	_, err = B64URLDecode("invalid base64!")
	if err == nil {
		t.Error("B64URLDecode() did not return error for invalid input")
	}

	_, err = B64DecodeToBytes("invalid base64!")
	if err == nil {
		t.Error("B64DecodeToBytes() did not return error for invalid input")
	}
}
