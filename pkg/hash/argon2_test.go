package hash

import (
	"bytes"
	"encoding/base64"
	"strings"
	"testing"

	"github.com/matryer/is"
)

func TestArgon2idBytes(t *testing.T) {
	i := is.New(t)
	hash, err := Argon2idBytes([]byte("password"))
	i.NoErr(err)
	i.Equal(len(hash), 32)
}

func TestArgon2idBytesEmpty(t *testing.T) {
	i := is.New(t)
	hash, err := Argon2idBytes([]byte(""))
	i.NoErr(err)
	i.Equal(len(hash), 32)
}

func TestArgon2idString(t *testing.T) {
	i := is.New(t)
	hash, err := Argon2idString("password")
	i.NoErr(err)
	i.Equal(len(hash), 32)
}

func TestArgon2idStringToString(t *testing.T) {
	i := is.New(t)
	hash, err := Argon2idStringToString("password")
	i.NoErr(err)
	parts := strings.Split(hash, "$")
	i.Equal(len(parts), 6)
	i.Equal(parts[1], "argon2id")
	i.Equal(parts[2], "v=19")
	i.True(strings.HasPrefix(parts[3], "m=2,t=32768,p=4"))

	_, err = base64.RawStdEncoding.DecodeString(parts[4])
	i.NoErr(err)
	_, err = base64.RawStdEncoding.DecodeString(parts[5])
	i.NoErr(err)
}

func TestArgon2idConsistencyWithSalt(t *testing.T) {
	i := is.New(t)
	salt := []byte("0123456789abcdef")
	hash1 := Argon2idStringWithSalt("password", salt)
	hash2 := Argon2idStringWithSalt("password", salt)
	i.True(bytes.Equal(hash1, hash2))
}

func TestArgon2idDifferentPasswords(t *testing.T) {
	i := is.New(t)
	salt := []byte("0123456789abcdef")
	hash1 := Argon2idStringWithSalt("password1", salt)
	hash2 := Argon2idStringWithSalt("password2", salt)
	i.True(!bytes.Equal(hash1, hash2))
}

func TestArgon2idDifferentSalts(t *testing.T) {
	i := is.New(t)
	hash1, err := Argon2idStringToString("password")
	i.NoErr(err)
	hash2, err := Argon2idStringToString("password")
	i.NoErr(err)
	i.True(hash1 != hash2)
}

func TestArgon2idCustomParams(t *testing.T) {
	i := is.New(t)
	params := Argon2idParams{
		Memory:      4,
		Iterations:  1,
		Parallelism: 2,
		KeyLen:      64,
	}
	salt := []byte("0123456789abcdef")
	hash := Argon2idStringToStringWithParams("password", salt, params)
	i.True(strings.Contains(hash, "m=4,t=1,p=2"))
}

func TestArgon2idMinimalParams(t *testing.T) {
	i := is.New(t)
	params := Argon2idParams{
		Memory:      1,
		Iterations:  1,
		Parallelism: 1,
		KeyLen:      16,
	}
	salt := []byte("0123456789abcdef")
	hash := Argon2idStringToStringWithParams("password", salt, params)
	i.True(strings.Contains(hash, "m=1,t=1,p=1"))
}

func TestArgon2idShortSalt(t *testing.T) {
	i := is.New(t)
	salt := []byte("short")
	hash := Argon2idStringWithSalt("password", salt)
	i.Equal(len(hash), 32)
}

func TestArgon2idLongSalt(t *testing.T) {
	i := is.New(t)
	salt := make([]byte, 100)
	hash := Argon2idStringWithSalt("password", salt)
	i.Equal(len(hash), 32)
}

func TestArgon2idNilSalt(t *testing.T) {
	i := is.New(t)
	hash := Argon2idStringWithSalt("password", nil)
	i.Equal(len(hash), 32)
}

func TestArgon2idUnicodePassword(t *testing.T) {
	i := is.New(t)
	hash, err := Argon2idStringToString("パスワード")
	i.NoErr(err)
	i.True(strings.HasPrefix(hash, "$argon2id$v=19$"))
}

func TestArgon2idGeneratedSaltLength(t *testing.T) {
	i := is.New(t)
	hash, err := Argon2idStringToString("password")
	i.NoErr(err)
	parts := strings.Split(hash, "$")
	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	i.NoErr(err)
	i.Equal(len(salt), 16)
}

func TestArgon2idAllBytesVariants(t *testing.T) {
	i := is.New(t)
	password := []byte("password")
	salt := []byte("0123456789abcdef")
	params := Argon2idDefaultParams

	hash1, err := Argon2idBytes(password)
	i.NoErr(err)
	i.Equal(len(hash1), 32)

	hash2 := Argon2idBytesWithSalt(password, salt)
	i.Equal(len(hash2), 32)

	hash3 := Argon2idBytesWithParams(password, salt, params)
	i.Equal(len(hash3), 32)
}

func TestArgon2idAllStringVariants(t *testing.T) {
	i := is.New(t)
	password := "password"
	salt := []byte("0123456789abcdef")
	params := Argon2idDefaultParams

	hash1, err := Argon2idStringToString(password)
	i.NoErr(err)
	i.True(strings.HasPrefix(hash1, "$argon2id$"))

	hash2 := Argon2idStringToStringWithSalt(password, salt)
	i.True(strings.HasPrefix(hash2, "$argon2id$"))

	hash3 := Argon2idStringToStringWithParams(password, salt, params)
	i.True(strings.HasPrefix(hash3, "$argon2id$"))
}
