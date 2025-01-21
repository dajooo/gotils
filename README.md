# gotils

Simple Go utility functions for common operations.

> [!NOTE]
> This library is under active development. More utility functions will be added in the future.

## Why
I created this library because existing solutions didn't provide all the utilities I needed in my projects. While there are many great utility libraries out there, they either lack certain functions or include too many dependencies. This library aims to be a complete solution for my use cases, though it may be opinionated in its implementation choices.

## Functions

### Pointer
- `Of(v T) *T` - Creates a pointer from a value
- `Resolve(v *T) T` - Safely resolves a pointer to its value

### Slice
- `Map[T, R](slice []T, fn func(T) R) []R` - Maps a slice to another type
- `ToPtr[T](slice []T) []*T` - Converts slice of values to slice of pointers
- `FromPtr[T](slice []*T) []T` - Converts slice of pointers to slice of values

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

## Install
```
go get github.com/dajooo/gotils
```