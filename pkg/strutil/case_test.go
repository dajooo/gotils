package strutil

import (
	"sync"
	"testing"

	"github.com/matryer/is"
)

func TestEmptyCamelCase(t *testing.T) {
	i := is.New(t)
	sc := StringCase{Kind: CamelCase, Parts: []string{}}
	i.Equal(sc.String(), "")
}

func TestBasicCamelCase(t *testing.T) {
	i := is.New(t)
	sc := StringCase{Kind: CamelCase, Parts: []string{"hello", "world"}}
	i.Equal(sc.String(), "helloWorld")
}

func TestBasicPascalCase(t *testing.T) {
	i := is.New(t)
	sc := StringCase{Kind: PascalCase, Parts: []string{"hello", "world"}}
	i.Equal(sc.String(), "HelloWorld")
}

func TestBasicSnakeCase(t *testing.T) {
	i := is.New(t)
	sc := StringCase{Kind: SnakeCase, Parts: []string{"hello", "world"}}
	i.Equal(sc.String(), "hello_world")
}

func TestBasicKebabCase(t *testing.T) {
	i := is.New(t)
	sc := StringCase{Kind: KebabCase, Parts: []string{"hello", "world"}}
	i.Equal(sc.String(), "hello-world")
}

func TestBasicScreamingSnakeCase(t *testing.T) {
	i := is.New(t)
	sc := StringCase{Kind: ScreamingSnakeCase, Parts: []string{"hello", "world"}}
	i.Equal(sc.String(), "HELLO_WORLD")
}

func TestUnknownCase(t *testing.T) {
	i := is.New(t)
	sc := StringCase{Kind: "unknown", Parts: []string{"hello", "world"}}
	i.Equal(sc.String(), "")
}

func TestDetectEmptyString(t *testing.T) {
	i := is.New(t)
	result := DetectCase("")
	i.Equal(result, CamelCase)
}

func TestDetectSnakeCase(t *testing.T) {
	i := is.New(t)
	result := DetectCase("hello_world")
	i.Equal(result, SnakeCase)
}

func TestDetectKebabCase(t *testing.T) {
	i := is.New(t)
	result := DetectCase("hello-world")
	i.Equal(result, KebabCase)
}

func TestDetectPascalCase(t *testing.T) {
	i := is.New(t)
	result := DetectCase("HelloWorld")
	i.Equal(result, PascalCase)
}

func TestDetectCamelCase(t *testing.T) {
	i := is.New(t)
	result := DetectCase("helloWorld")
	i.Equal(result, CamelCase)
}

func TestParseSnakeCase(t *testing.T) {
	i := is.New(t)
	result := ParseCase("hello_world_test")
	expected := StringCase{
		Kind:  SnakeCase,
		Parts: []string{"hello", "world", "test"},
	}
	i.Equal(result, expected)
}

func TestParseKebabCase(t *testing.T) {
	i := is.New(t)
	result := ParseCase("hello-world-test")
	expected := StringCase{
		Kind:  KebabCase,
		Parts: []string{"hello", "world", "test"},
	}
	i.Equal(result, expected)
}

func TestParsePascalCase(t *testing.T) {
	i := is.New(t)
	result := ParseCase("HelloWorldTest")
	expected := StringCase{
		Kind:  PascalCase,
		Parts: []string{"hello", "world", "test"},
	}
	i.Equal(result, expected)
}

func TestParseCamelCase(t *testing.T) {
	i := is.New(t)
	result := ParseCase("helloWorldTest")
	expected := StringCase{
		Kind:  CamelCase,
		Parts: []string{"hello", "world", "test"},
	}
	i.Equal(result, expected)
}

func TestCapitalizeEmptyString(t *testing.T) {
	i := is.New(t)
	result := capitalize("")
	i.Equal(result, "")
}

func TestCapitalizeSingleLetter(t *testing.T) {
	i := is.New(t)
	result := capitalize("a")
	i.Equal(result, "A")
}

func TestCapitalizeWord(t *testing.T) {
	i := is.New(t)
	result := capitalize("hello")
	i.Equal(result, "Hello")
}

func TestStringCaseConversionsToCamelCase(t *testing.T) {
	i := is.New(t)
	sc := StringCase{Kind: SnakeCase, Parts: []string{"hello", "world", "test"}}
	i.Equal(sc.ToCamelCase(), "helloWorldTest")
}

func TestStringCaseConversionsToPascalCase(t *testing.T) {
	i := is.New(t)
	sc := StringCase{Kind: SnakeCase, Parts: []string{"hello", "world", "test"}}
	i.Equal(sc.ToPascalCase(), "HelloWorldTest")
}

func TestStringCaseConversionsToSnakeCase(t *testing.T) {
	i := is.New(t)
	sc := StringCase{Kind: CamelCase, Parts: []string{"hello", "world", "test"}}
	i.Equal(sc.ToSnakeCase(), "hello_world_test")
}

func TestStringCaseConversionsToKebabCase(t *testing.T) {
	i := is.New(t)
	sc := StringCase{Kind: CamelCase, Parts: []string{"hello", "world", "test"}}
	i.Equal(sc.ToKebabCase(), "hello-world-test")
}

func TestConcurrentStringCaseAccess(t *testing.T) {
	i := is.New(t)
	sc := StringCase{
		Kind:  CamelCase,
		Parts: []string{"hello", "world", "test"},
	}

	const goroutines = 100
	const iterations = 100

	var wg sync.WaitGroup
	wg.Add(goroutines)

	for g := 0; g < goroutines; g++ {
		go func() {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				result := sc.ToCamelCase()
				i.Equal(result, "helloWorldTest")
			}
		}()
	}

	wg.Wait()
}
