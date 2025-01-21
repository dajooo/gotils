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

## Install
```
go get github.com/dajooo/gotils
```