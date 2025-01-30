package password

import (
	"strings"
	"testing"

	"github.com/matryer/is"
)

func TestVerifyWithDefaults(t *testing.T) {
	i := is.New(t)
	err := Verify("Test123!@#")
	i.NoErr(err)
}

func TestVerifyTooShort(t *testing.T) {
	i := is.New(t)
	err := Verify("Test1!")
	i.True(err != nil)
	i.True(strings.Contains(err.Error(), "at least 8 characters"))
}

func TestVerifyTooLong(t *testing.T) {
	i := is.New(t)
	err := Verify(strings.Repeat("a", 129))
	i.True(err != nil)
	i.True(strings.Contains(err.Error(), "must not exceed 128 characters"))
}

func TestVerifyWithoutUpper(t *testing.T) {
	i := is.New(t)
	err := Verify("test123!@#", VerifyWithoutUppercaseOption())
	i.NoErr(err)
}
func TestVerifyWithoutLower(t *testing.T) {
	i := is.New(t)
	err := Verify("TEST123!@#", VerifyWithoutLowercaseOption())
	i.NoErr(err)
}

func TestVerifyWithoutNumbers(t *testing.T) {
	i := is.New(t)
	err := Verify("TestTest!@#", VerifyWithoutNumbersOption())
	i.NoErr(err)
}

func TestVerifyWithoutSpecial(t *testing.T) {
	i := is.New(t)
	err := Verify("Test1234", VerifyWithoutSpecialOption())
	i.NoErr(err)
}

func TestVerifyCustomLength(t *testing.T) {
	i := is.New(t)
	err := Verify("Test1!", VerifyWithMinLengthOption(6))
	i.NoErr(err)
}

func TestVerifyMultipleRequirements(t *testing.T) {
	i := is.New(t)
	err := Verify("test",
		VerifyWithMinLengthOption(4),
		VerifyWithoutUppercaseOption(),
		VerifyWithoutNumbersOption(),
		VerifyWithoutSpecialOption(),
	)
	i.NoErr(err)
}
