package encoding

import (
	"testing"

	"github.com/matryer/is"
)

func TestB64EncodeAndDecode(t *testing.T) {
	i := is.New(t)
	original := "Hello, World!"

	encoded := B64Encode(original)
	decoded, err := B64Decode(encoded)
	i.NoErr(err)
	i.Equal(decoded, original)
}

func TestB64URLEncodeAndDecode(t *testing.T) {
	i := is.New(t)
	original := "Hello, World!"

	encoded := B64URLEncode(original)
	decoded, err := B64URLDecode(encoded)
	i.NoErr(err)
	i.Equal(decoded, original)
}

func TestB64RawEncodeAndDecode(t *testing.T) {
	i := is.New(t)
	original := "Hello, World!"

	encoded := B64RawEncode(original)
	decoded, err := B64RawDecode(encoded)
	i.NoErr(err)
	i.Equal(decoded, original)
}

func TestB64MustDecodeWithInvalidInput(t *testing.T) {
	i := is.New(t)

	defer func() {
		r := recover()
		i.True(r != nil) // Should panic for invalid base64
	}()

	MustB64Decode("invalid-base64")
}
