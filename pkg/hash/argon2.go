package hash

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/argon2"
)

type Argon2idParams struct {
	Memory      uint32
	Iterations  uint32
	Parallelism uint8
	KeyLen      uint32
}

var Argon2idDefaultParams = Argon2idParams{
	Memory:      2,
	Iterations:  32 * 1024,
	Parallelism: 4,
	KeyLen:      32,
}

func generateArgon2idSalt() ([]byte, error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

func formatArgon2id(salt, hash []byte, p Argon2idParams) string {
	return fmt.Sprintf("$argon2id$v=19$m=%d,t=%d,p=%d$%s$%s",
		p.Memory,
		p.Iterations,
		p.Parallelism,
		base64.RawStdEncoding.EncodeToString(salt),
		base64.RawStdEncoding.EncodeToString(hash))
}

// []byte output
func Argon2idBytes(password []byte) ([]byte, error) {
	salt, err := generateArgon2idSalt()
	if err != nil {
		return nil, err
	}
	return Argon2idBytesWithParams(password, salt, Argon2idDefaultParams), nil
}

func Argon2idBytesWithSalt(password, salt []byte) []byte {
	return Argon2idBytesWithParams(password, salt, Argon2idDefaultParams)
}

func Argon2idBytesWithParams(password, salt []byte, p Argon2idParams) []byte {
	return argon2.IDKey(password, salt, p.Iterations, p.Memory, p.Parallelism, p.KeyLen)
}

// string output
func Argon2idBytesToString(password []byte) (string, error) {
	salt, err := generateArgon2idSalt()
	if err != nil {
		return "", err
	}
	hash := Argon2idBytesWithParams(password, salt, Argon2idDefaultParams)
	return formatArgon2id(salt, hash, Argon2idDefaultParams), nil
}

func Argon2idBytesToStringWithSalt(password, salt []byte) string {
	hash := Argon2idBytesWithSalt(password, salt)
	return formatArgon2id(salt, hash, Argon2idDefaultParams)
}

func Argon2idBytesToStringWithParams(password, salt []byte, p Argon2idParams) string {
	hash := Argon2idBytesWithParams(password, salt, p)
	return formatArgon2id(salt, hash, p)
}

// string input -> []byte output
func Argon2idString(password string) ([]byte, error) {
	return Argon2idBytes([]byte(password))
}

func Argon2idStringWithSalt(password string, salt []byte) []byte {
	return Argon2idBytesWithSalt([]byte(password), salt)
}

func Argon2idStringWithParams(password string, salt []byte, p Argon2idParams) []byte {
	return Argon2idBytesWithParams([]byte(password), salt, p)
}

// string input -> string output
func Argon2idStringToString(password string) (string, error) {
	return Argon2idBytesToString([]byte(password))
}

func Argon2idStringToStringWithSalt(password string, salt []byte) string {
	return Argon2idBytesToStringWithSalt([]byte(password), salt)
}

func Argon2idStringToStringWithParams(password string, salt []byte, p Argon2idParams) string {
	return Argon2idBytesToStringWithParams([]byte(password), salt, p)
}
