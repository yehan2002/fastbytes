[![Test](https://github.com/yehan2002/fastbytes/actions/workflows/go.yml/badge.svg)](https://github.com/yehan2002/fastbytes/actions/workflows/go.yml)

# FastBytes

FastBytes is a go package for translating between slices with fixed-size integers/floats and byte slices.


### Usage of `unsafe.Pointer`

This package uses the `unsafe` package to covert between slice/array types and to extract pointers from `interface{}` values. To disable the usage of unsafe set the `no_unsafe` build tag when building.