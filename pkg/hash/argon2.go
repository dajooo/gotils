package hash

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"strings"

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

func Argon2idString(password string) ([]byte, error) {
	return Argon2idBytes([]byte(password))
}

func Argon2idStringWithSalt(password string, salt []byte) []byte {
	return Argon2idBytesWithSalt([]byte(password), salt)
}

func Argon2idStringWithParams(password string, salt []byte, p Argon2idParams) []byte {
	return Argon2idBytesWithParams([]byte(password), salt, p)
}

func Argon2idStringToString(password string) (string, error) {
	return Argon2idBytesToString([]byte(password))
}

func Argon2idStringToStringWithSalt(password string, salt []byte) string {
	return Argon2idBytesToStringWithSalt([]byte(password), salt)
}

func Argon2idStringToStringWithParams(password string, salt []byte, p Argon2idParams) string {
	return Argon2idBytesToStringWithParams([]byte(password), salt, p)
}

func VerifyArgon2id(hashedPassword string, password []byte) (bool, error) {
	if len(hashedPassword) == 0 || !strings.HasPrefix(hashedPassword, "$") {
		return false, fmt.Errorf("invalid hash format: must start with $")
	}

	parts := strings.Split(hashedPassword, "$")
	if len(parts) != 6 {
		return false, fmt.Errorf("invalid hash format: expected 6 parts, got %d", len(parts))
	}

	if parts[1] != "argon2id" {
		return false, fmt.Errorf("invalid algorithm: expected argon2id, got %s", parts[1])
	}

	if parts[2] != "v=19" {
		return false, fmt.Errorf("invalid version: expected v=19, got %s", parts[2])
	}

	var memory uint32
	var iterations uint32
	var parallelism uint8

	paramCount, err := fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &memory, &iterations, &parallelism)
	if err != nil || paramCount != 3 {
		return false, fmt.Errorf("invalid parameters format: failed to parse m,t,p values")
	}

	if memory == 0 || iterations == 0 || parallelism == 0 {
		return false, fmt.Errorf("invalid parameters: values must be greater than 0")
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return false, fmt.Errorf("illegal base64")
	}

	if len(salt) == 0 {
		return false, fmt.Errorf("invalid salt: cannot be empty")
	}

	decodedHash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return false, fmt.Errorf("illegal base64")
	}

	if len(decodedHash) == 0 {
		return false, fmt.Errorf("invalid hash: cannot be empty")
	}

	params := Argon2idParams{
		Memory:      memory,
		Iterations:  iterations,
		Parallelism: parallelism,
		KeyLen:      uint32(len(decodedHash)),
	}

	computedHash := Argon2idBytesWithParams(password, salt, params)

	return bytes.Equal(decodedHash, computedHash), nil
}

func VerifyArgon2idString(hashedPassword, password string) (bool, error) {
	return VerifyArgon2id(hashedPassword, []byte(password))
}

func MustVerifyArgon2id(hashedPassword string, password []byte) bool {
	match, err := VerifyArgon2id(hashedPassword, password)
	if err != nil {
		panic(err)
	}
	return match
}

func MustVerifyArgon2idString(hashedPassword, password string) bool {
	return MustVerifyArgon2id(hashedPassword, []byte(password))
}
