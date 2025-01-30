package password

import (
	"crypto/rand"
	"errors"
	"math/big"
	"strings"
)

const (
	upperChars   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowerChars   = "abcdefghijklmnopqrstuvwxyz"
	numberChars  = "0123456789"
	specialChars = "!@#$%^&*()_+-=[]{}|;:,.<>?"
)

var (
	ErrInvalidCharset = errors.New("invalid charset: no characters available for password generation")
)

// GenerateConfig holds password generation settings
type GenerateConfig struct {
	Length        int
	UseUpper      bool
	UseLower      bool
	UseNumbers    bool
	UseSpecial    bool
	CustomCharset string
	ExcludeChars  string
}

// GenerateOption is a function that modifies GenerateConfig
type GenerateOption func(*GenerateConfig)

// Generate creates a random password based on the specified options
func Generate(options ...GenerateOption) (string, error) {
	config := GenerateConfig{
		Length:     16,
		UseUpper:   true,
		UseLower:   true,
		UseNumbers: true,
		UseSpecial: true,
	}

	for _, opt := range options {
		opt(&config)
	}

	var charset strings.Builder

	if config.CustomCharset != "" {
		charset.WriteString(config.CustomCharset)
	} else {
		if config.UseUpper {
			charset.WriteString(upperChars)
		}
		if config.UseLower {
			charset.WriteString(lowerChars)
		}
		if config.UseNumbers {
			charset.WriteString(numberChars)
		}
		if config.UseSpecial {
			charset.WriteString(specialChars)
		}
	}

	if config.ExcludeChars != "" {
		charsetStr := charset.String()
		for _, c := range config.ExcludeChars {
			charsetStr = strings.ReplaceAll(charsetStr, string(c), "")
		}
		charset.Reset()
		charset.WriteString(charsetStr)
	}

	if charset.Len() == 0 {
		return "", ErrInvalidCharset
	}

	var password strings.Builder
	charsetStr := charset.String()
	charsetLength := big.NewInt(int64(len(charsetStr)))

	for i := 0; i < config.Length; i++ {
		n, err := rand.Int(rand.Reader, charsetLength)
		if err != nil {
			return "", err
		}
		password.WriteByte(charsetStr[n.Int64()])
	}

	return password.String(), nil
}

// MustGenerate is a helper that wraps Generate and panics if an error occurs
func MustGenerate(options ...GenerateOption) string {
	password, err := Generate(options...)
	if err != nil {
		panic(err)
	}
	return password
}

// GenerateWithLengthOption sets the desired length for the generated password
func GenerateWithLengthOption(length int) GenerateOption {
	return func(c *GenerateConfig) {
		if length > 0 {
			c.Length = length
		}
	}
}

// GenerateWithoutUpperOption excludes uppercase letters from password generation
func GenerateWithoutUpperOption() GenerateOption {
	return func(c *GenerateConfig) {
		c.UseUpper = false
	}
}

// GenerateWithoutLowerOption excludes lowercase letters from password generation
func GenerateWithoutLowerOption() GenerateOption {
	return func(c *GenerateConfig) {
		c.UseLower = false
	}
}

// GenerateWithoutNumbersOption excludes numbers from password generation
func GenerateWithoutNumbersOption() GenerateOption {
	return func(c *GenerateConfig) {
		c.UseNumbers = false
	}
}

// GenerateWithoutSpecialOption excludes special characters from password generation
func GenerateWithoutSpecialOption() GenerateOption {
	return func(c *GenerateConfig) {
		c.UseSpecial = false
	}
}

// GenerateWithCustomCharsetOption sets a custom character set for password generation
func GenerateWithCustomCharsetOption(charset string) GenerateOption {
	return func(c *GenerateConfig) {
		c.CustomCharset = charset
	}
}

// GenerateWithExcludedCharsOption specifies characters to exclude from password generation
func GenerateWithExcludedCharsOption(chars string) GenerateOption {
	return func(c *GenerateConfig) {
		c.ExcludeChars = chars
	}
}
