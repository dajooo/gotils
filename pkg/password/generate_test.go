package password

import (
	"strings"
	"sync"
	"testing"

	"github.com/matryer/is"
)

func TestGenerateDefault(t *testing.T) {
	i := is.New(t)
	password, err := Generate()
	i.NoErr(err)
	i.Equal(len(password), 16)
}

func TestGenerateWithCustomLength(t *testing.T) {
	i := is.New(t)
	password, err := Generate(GenerateWithLengthOption(20))
	i.NoErr(err)
	i.Equal(len(password), 20)
}

func TestGenerateWithoutUpper(t *testing.T) {
	i := is.New(t)
	password, err := Generate(
		GenerateWithoutUpperOption(),
		GenerateWithCustomCharsetOption(lowerChars+numberChars+specialChars),
	)
	i.NoErr(err)
	i.True(!containsAny(password, upperChars))
}

func TestGenerateWithoutLower(t *testing.T) {
	i := is.New(t)
	password, err := Generate(
		GenerateWithoutLowerOption(),
		GenerateWithCustomCharsetOption(upperChars+numberChars+specialChars),
	)
	i.NoErr(err)
	i.True(!containsAny(password, lowerChars))
}

func TestGenerateWithoutNumbers(t *testing.T) {
	i := is.New(t)
	password, err := Generate(
		GenerateWithoutNumbersOption(),
		GenerateWithCustomCharsetOption(upperChars+lowerChars+specialChars),
	)
	i.NoErr(err)
	i.True(!containsAny(password, numberChars))
}

func TestGenerateWithoutSpecial(t *testing.T) {
	i := is.New(t)
	password, err := Generate(
		GenerateWithoutSpecialOption(),
		GenerateWithCustomCharsetOption(upperChars+lowerChars+numberChars),
	)
	i.NoErr(err)
	i.True(!containsAny(password, specialChars))
}

func TestGenerateWithCustomCharset(t *testing.T) {
	i := is.New(t)
	charset := "ABC123"
	password, err := Generate(GenerateWithCustomCharsetOption(charset))
	i.NoErr(err)
	i.True(containsOnly(password, charset))
}

func TestGenerateWithExcludedChars(t *testing.T) {
	i := is.New(t)
	excluded := "0O1lI"
	password, err := Generate(GenerateWithExcludedCharsOption(excluded))
	i.NoErr(err)
	i.True(!containsAny(password, excluded))
}

func TestGenerateEmptyResult(t *testing.T) {
	i := is.New(t)
	_, err := Generate(
		GenerateWithoutUpperOption(),
		GenerateWithoutLowerOption(),
		GenerateWithoutNumbersOption(),
		GenerateWithoutSpecialOption(),
	)
	i.True(err != nil)
	i.Equal(err, ErrInvalidCharset)
}

func TestMustGenerate(t *testing.T) {
	i := is.New(t)
	password := MustGenerate()
	i.Equal(len(password), 16)
}

func TestMustGeneratePanic(t *testing.T) {
	i := is.New(t)
	defer func() {
		r := recover()
		i.True(r != nil)
		if err, ok := r.(error); ok {
			i.Equal(err, ErrInvalidCharset)
		}
	}()
	MustGenerate(
		GenerateWithoutUpperOption(),
		GenerateWithoutLowerOption(),
		GenerateWithoutNumbersOption(),
		GenerateWithoutSpecialOption(),
	)
}

func TestConcurrentGeneration(t *testing.T) {
	i := is.New(t)
	const goroutines = 100
	const iterations = 100

	var wg sync.WaitGroup
	wg.Add(goroutines)

	for g := 0; g < goroutines; g++ {
		go func() {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				password, err := Generate()
				i.NoErr(err)
				i.Equal(len(password), 16)
			}
		}()
	}

	wg.Wait()
}

func containsAny(s, chars string) bool {
	for _, c := range chars {
		if strings.ContainsRune(s, c) {
			return true
		}
	}
	return false
}

func containsOnly(s, chars string) bool {
	for _, c := range s {
		if !strings.ContainsRune(chars, c) {
			return false
		}
	}
	return true
}
