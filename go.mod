module github.com/yehan2002/fastbytes/v2

go 1.19

// This version has a broken mod file
retract v2.0.0

// This version has incorrect import paths and will behave like v1
retract v2.0.1

require (
	github.com/google/go-cmp v0.5.9
	github.com/yehan2002/errors v1.5.4
	github.com/yehan2002/is/v2 v2.4.0
	golang.org/x/sys v0.11.0
)
