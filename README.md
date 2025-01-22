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
- `MapToPtr[I, O](slice []I, fn func(I) O) []*O` - Maps slice to another type and converts to pointers
- `MapFromPtr[I, O](slice []*I, fn func(I) O) []O` - Maps slice of pointers to another type
- `ToPtr[T](slice []T) []*T` - Converts slice of values to slice of pointers
- `FromPtr[T](slice []*T) []T` - Converts slice of pointers to slice of values
- `MapNonPtrToPtr[I, O](slice []I, fn func(*I) *O) []*O` - Maps non-pointer slice using a function that takes and returns pointers

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
- `MarshalJSON[T](v T, marshaler ...json.Marshaler) ([]byte, error)` - Marshals type T into JSON bytes with optional custom marshaler
- `MustMarshalJSON[T](v T, marshaler ...json.Marshaler) []byte` - Same as MarshalJSON but panics on error
- `UnmarshalJSON[T](data []byte, unmarshaler ...json.Unmarshaler) (T, error)` - Unmarshals JSON bytes into type T with optional custom unmarshaler
- `MustUnmarshalJSON[T](data []byte, unmarshaler ...json.Unmarshaler) T` - Same as UnmarshalJSON but panics on error

## Install
```
go get github.com/dajooo/gotils
```