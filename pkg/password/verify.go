package password

import (
	"fmt"
	"strings"
	"unicode"
)

// VerifyConfig holds password validation settings
type VerifyConfig struct {
	MinLength        int
	MaxLength        int
	RequireUppercase bool
	RequireLowercase bool
	RequireNumbers   bool
	RequireSpecial   bool
}

// VerifyOption is a function that modifies VerifyConfig
type VerifyOption func(*VerifyConfig)

// Verify Validate password against common criteria
func Verify(password string, options ...VerifyOption) error {
	// Default configuration
	config := VerifyConfig{
		MinLength:        8,
		RequireUppercase: true,
		RequireLowercase: true,
		RequireNumbers:   true,
		RequireSpecial:   true,
		MaxLength:        128,
	}

	for _, opt := range options {
		opt(&config)
	}

	if len(password) < config.MinLength {
		return fmt.Errorf("password must be at least %d characters long", config.MinLength)
	}
	if len(password) > config.MaxLength {
		return fmt.Errorf("password must not exceed %d characters", config.MaxLength)
	}

	var (
		hasUpper   bool
		hasLower   bool
		hasNumber  bool
		hasSpecial bool
	)

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	var violations []string
	if config.RequireUppercase && !hasUpper {
		violations = append(violations, "uppercase letter")
	}
	if config.RequireLowercase && !hasLower {
		violations = append(violations, "lowercase letter")
	}
	if config.RequireNumbers && !hasNumber {
		violations = append(violations, "number")
	}
	if config.RequireSpecial && !hasSpecial {
		violations = append(violations, "special character")
	}

	if len(violations) > 0 {
		return fmt.Errorf("password must contain at least one %s", strings.Join(violations, ", "))
	}

	return nil
}

// VerifyWithMinLengthOption sets minimum password length
func VerifyWithMinLengthOption(length int) VerifyOption {
	return func(c *VerifyConfig) {
		c.MinLength = length
	}
}

// VerifyWithMaxLengthOption sets maximum password length
func VerifyWithMaxLengthOption(length int) VerifyOption {
	return func(c *VerifyConfig) {
		c.MaxLength = length
	}
}

// VerifyWithoutUppercaseOption removes uppercase requirement
func VerifyWithoutUppercaseOption() VerifyOption {
	return func(c *VerifyConfig) {
		c.RequireUppercase = false
	}
}

// VerifyWithoutLowercaseOption removes lowercase requirement
func VerifyWithoutLowercaseOption() VerifyOption {
	return func(c *VerifyConfig) {
		c.RequireLowercase = false
	}
}

// VerifyWithoutNumbersOption removes numbers requirement
func VerifyWithoutNumbersOption() VerifyOption {
	return func(c *VerifyConfig) {
		c.RequireNumbers = false
	}
}

// VerifyWithoutSpecialOption removes special characters requirement
func VerifyWithoutSpecialOption() VerifyOption {
	return func(c *VerifyConfig) {
		c.RequireSpecial = false
	}
}
