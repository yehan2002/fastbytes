[![Test](https://github.com/yehan2002/fastbytes/actions/workflows/go.yml/badge.svg)](https://github.com/yehan2002/fastbytes/actions/workflows/go.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/yehan2002/fastbytes/v2)](https://goreportcard.com/report/github.com/yehan2002/fastbytes/v2)[![Go Reference](https://pkg.go.dev/badge/github.com/yehan2002/fastbytes.svg)](https://pkg.go.dev/github.com/yehan2002/fastbytes)

# FastBytes

FastBytes is a go package for translating between slices with fixed-size integers/floats and byte slices.

## Supported Types
Floats and all signed/unsigned integers except uint and int can be translated by this package.
`uint` and `int` are not supported since their size is platform dependent.

## Usage

[GoDoc](https://pkg.go.dev/github.com/yehan2002/fastbytes)


#### Usage of assembly code
This package uses assembly for copying data on certain platforms.
To disable the usage of assembly set the `purego` build tag when building.

#### Usage of `unsafe.Pointer`

This package uses the `unsafe` package to covert between slice/array types and to extract pointers from `interface{}` values. To disable the usage of unsafe set the `no_unsafe` build tag when building. Note that this also disables the usage of assembly.
