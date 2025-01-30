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

### Hash
#### Argon2id
Password hashing functions using Argon2id with configurable parameters:

Parameters struct:
```go
type Argon2idParams struct {
    Memory      uint32 // Memory usage in KB
    Iterations  uint32 // Number of iterations
    Parallelism uint8  // Degree of parallelism
    KeyLen      uint32 // Length of the hash in bytes
}
```

Byte operations:
- `Argon2idBytes(password []byte) ([]byte, error)` - Hash password bytes with default parameters
- `Argon2idBytesWithSalt(password, salt []byte) []byte` - Hash password bytes with custom salt
- `Argon2idBytesWithParams(password, salt []byte, p Argon2idParams) []byte` - Hash password bytes with custom salt and parameters

String input operations:
- `Argon2idString(password string) ([]byte, error)` - Hash password string, output bytes
- `Argon2idStringWithSalt(password string, salt []byte) []byte` - Hash password string with custom salt
- `Argon2idStringWithParams(password string, salt []byte, p Argon2idParams) []byte` - Hash password string with custom salt and parameters

String output operations:
- `Argon2idBytesToString(password []byte) (string, error)` - Hash password bytes, output formatted string
- `Argon2idBytesToStringWithSalt(password, salt []byte) string` - Hash password bytes with custom salt, output formatted string
- `Argon2idBytesToStringWithParams(password, salt []byte, p Argon2idParams) string` - Hash password bytes with custom salt and parameters, output formatted string

String input/output operations:
- `Argon2idStringToString(password string) (string, error)` - Hash password string, output formatted string
- `Argon2idStringToStringWithSalt(password string, salt []byte) string` - Hash password string with custom salt, output formatted string
- `Argon2idStringToStringWithParams(password string, salt []byte, p Argon2idParams) string` - Hash password string with custom salt and parameters, output formatted string

Verification operations:
- `VerifyArgon2id(hashedPassword string, password []byte) (bool, error)` - Verify password bytes against a hashed password
- `VerifyArgon2idString(hashedPassword, password string) (bool, error)` - Verify password string against a hashed password
- `MustVerifyArgon2id(hashedPassword string, password []byte) bool` - Same as VerifyArgon2id but panics on error
- `MustVerifyArgon2idString(hashedPassword, password string) bool` - Same as VerifyArgon2idString but panics on error

Default parameters:
- Memory: 2 KB
- Iterations: 32,768
- Parallelism: 4
- Key Length: 32 bytes

String output format: `$argon2id$v=19$m=memory,t=iterations,p=parallelism$salt$hash`

### Password
#### Generation
Generate secure passwords with configurable options:

```go
type GenerateConfig struct {
    Length        int     // Password length
    UseUpper      bool    // Include uppercase letters
    UseLower      bool    // Include lowercase letters
    UseNumbers    bool    // Include numbers
    UseSpecial    bool    // Include special characters
    CustomCharset string  // Custom character set
    ExcludeChars  string  // Characters to exclude
}
```

Generation functions:
- `Generate(options ...GenerateOption) (string, error)` - Generates a password with specified options
- `MustGenerate(options ...GenerateOption) string` - Same as Generate but panics on error

Generation options:
- `GenerateWithLengthOption(length int)` - Sets password length (default: 16)
- `GenerateWithoutUpperOption()` - Excludes uppercase letters
- `GenerateWithoutLowerOption()` - Excludes lowercase letters
- `GenerateWithoutNumbersOption()` - Excludes numbers
- `GenerateWithoutSpecialOption()` - Excludes special characters
- `GenerateWithCustomCharsetOption(charset string)` - Uses custom character set
- `GenerateWithExcludedCharsOption(chars string)` - Excludes specific characters

Default character sets:
- Uppercase: A-Z
- Lowercase: a-z
- Numbers: 0-9
- Special: !@#$%^&*()_+-=[]{}|;:,.<>?

#### Verification
Validate passwords against common security criteria:

```go
type VerifyConfig struct {
    MinLength    int     // Minimum length required
    MaxLength    int     // Maximum length allowed
    UseUpper     bool    // Require uppercase letters
    UseLower     bool    // Require lowercase letters
    UseNumbers   bool    // Require numbers
    UseSpecial   bool    // Require special characters
}
```

Verification functions:
- `Verify(password string, options ...VerifyOption) error` - Validates password against criteria
- `MustVerify(password string, options ...VerifyOption)` - Same as Verify but panics on error

Verification options:
- `VerifyWithMinLengthOption(length int)` - Sets minimum length (default: 8)
- `VerifyWithMaxLengthOption(length int)` - Sets maximum length (default: 128)
- `VerifyWithoutUpperOption()` - Removes uppercase requirement
- `VerifyWithoutLowerOption()` - Removes lowercase requirement
- `VerifyWithoutNumbersOption()` - Removes numbers requirement
- `VerifyWithoutSpecialOption()` - Removes special characters requirement

Default requirements:
- Minimum length: 8 characters
- Maximum length: 128 characters
- Must contain at least one:
    - Uppercase letter
    - Lowercase letter
    - Number
    - Special character

## Install
```
go get dario.lol/gotils
```