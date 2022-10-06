module github.com/yehan2002/fastbytes/v2

go 1.19

// This version has a broken mod file
retract v2.0.0

// This version has incorrect import paths and will behave like v1
retract v2.0.1

require (
	github.com/yehan2002/errors v1.4.0
	github.com/yehan2002/is/v2 v2.2.3
	golang.org/x/sys v0.0.0-20220928140112-f11e5e49a4ec
)

require github.com/go-test/deep v1.0.8 // indirect
