# gotils

Simple Go utility functions for common operations.

> [!NOTE]
> This library is under active development. More utility functions will be added in the future.

## Why
I created this library because existing solutions didn't provide all the utilities I needed in my projects. While there are many great utility libraries out there, they either lack certain functions or include too many dependencies. This library aims to be a complete solution for my use cases, though it may be opinionated in its implementation choices.

## Functions

### Pointer
- `Of(v T) *T` - Creates a pointer from a value
- `OfOk(v T, ok bool) *T` - Creates a pointer from a value if ok is true, nil otherwise
- `Resolve(v *T) T` - Safely resolves a pointer to its value
- `ResolveOrDefault(v *T) T` - Resolves a pointer or returns zero value if nil
- `ResolveOr(v *T, defaultValue T) T` - Resolves a pointer or returns provided default value if nil

### Error
- `Must[T](value T, err error) T` - Returns value or panics if error occurs
- `MustOk[T](value T, ok bool) T` - Returns value or panics if ok is false

### Slice
- `Map[T, R](slice []T, fn func(T) R) []R` - Maps a slice to another type
- `MapIndexed[I, O](slice []I, fn func(index int, value I) O) []O` - Maps a slice to another type with index access
- `MapToPtr[I, O](slice []I, fn func(I) O) []*O` - Maps slice to another type and converts to pointers
- `MapFromPtr[I, O](slice []*I, fn func(I) O) []O` - Maps slice of pointers to another type
- `ToPtr[T](slice []T) []*T` - Converts slice of values to slice of pointers
- `FromPtr[T](slice []*T) []T` - Converts slice of pointers to slice of values
- `MapNonPtrToPtr[I, O](slice []I, fn func(*I) *O) []*O` - Maps non-pointer slice using a function that takes and returns pointers
- `Filter[T](s []T, f func(T) bool) []T` - Filters slice elements based on predicate function
- `FilterInstanceOf[T](s []any) []T` - Filters slice elements by type assertion
- `FilterNotNil[T](s []*T) []T` - Filters out nil elements from slice of pointers

### File
#### Read
- `Read(path string) ([]byte, error)` - Reads file as bytes
- `ReadString(path string) (string, error)` - Reads file as string
- `ReadLines(path string) ([]string, error)` - Reads file as string array
- `ReadJson[T](path string, unmarshaler ...json.Unmarshaler) (T, error)` - Reads JSON file into type T

#### Write
- `Write(path string, data []byte) error` - Writes bytes to file
- `WriteString(path, data string) error` - Writes string to file
- `WriteLines(path string, lines []string) error` - Writes string array to file
- `WriteJson[T](path string, data T, indent ...string) error` - Writes type T as JSON to file

### Encoding
#### JSON
- `MarshalJSON[T](v T, marshaler ...json.Marshaler) ([]byte, error)` - Marshals type T into JSON bytes with optional custom marshaler
- `MustMarshalJSON[T](v T, marshaler ...json.Marshaler) []byte` - Same as MarshalJSON but panics on error
- `UnmarshalJSON[T](data []byte, unmarshaler ...json.Unmarshaler) (T, error)` - Unmarshals JSON bytes into type T with optional custom unmarshaler
- `MustUnmarshalJSON[T](data []byte, unmarshaler ...json.Unmarshaler) T` - Same as UnmarshalJSON but panics on error

#### Base64
String operations:
- `B64Encode(data string, padding ...rune) string` - Encodes string to base64 string
- `B64Decode(data string) (string, error)` - Decodes base64 string to string
- `MustB64Decode(data string) string` - Same as B64Decode but panics on error

Bytes operations:
- `B64EncodeBytes(data []byte, padding ...rune) string` - Encodes bytes to base64 string
- `B64DecodeBytes(data []byte) (string, error)` - Decodes base64 bytes to string
- `B64EncodeBytesToBytes(data []byte, padding ...rune) []byte` - Encodes bytes to base64 bytes
- `B64DecodeBytesToBytes(data []byte) ([]byte, error)` - Decodes base64 bytes to bytes

URL-safe variants:
- `B64URLEncode(data string, padding ...rune) string` - URL-safe base64 encoding
- `B64URLDecode(data string) (string, error)` - URL-safe base64 decoding
- `B64URLEncodeBytes(data []byte, padding ...rune) string` - URL-safe base64 encoding of bytes
- `B64URLDecodeBytes(data []byte) (string, error)` - URL-safe base64 decoding to string

### Case
- `DetectCase(s string) StringCaseKind` - Detects the case style of a string
- `ParseCase(s string) StringCase` - Parses a string into its case parts
- `(StringCase) ToCamelCase() string` - Converts to camelCase
- `(StringCase) ToPascalCase() string` - Converts to PascalCase
- `(StringCase) ToSnakeCase() string` - Converts to snake_case
- `(StringCase) ToKebabCase() string` - Converts to kebab-case
- `(StringCase) ToScreamingSnakeCase() string` - Converts to SCREAMING_SNAKE_CASE

Supported case types:
- CamelCase (e.g. "camelCase")
- PascalCase (e.g. "PascalCase")
- SnakeCase (e.g. "snake_case")
- KebabCase (e.g. "kebab-case")
- ScreamingSnakeCase (e.g. "SCREAMING_SNAKE_CASE")

## Install
```
go get dario.lol/gotils
```