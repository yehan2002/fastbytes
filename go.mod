module github.com/yehan2002/fastbytes/v2

go 1.19

// This version has a broken mod file
retract v2.0.0

// This version has incorrect import paths and will behave like v1
retract v2.0.1

require (
	github.com/mmcloughlin/avo v0.4.0
	github.com/yehan2002/errors v1.4.0
	github.com/yehan2002/is/v2 v2.2.3
	golang.org/x/sys v0.0.0-20220731174439-a90be440212d
)

require (
	github.com/go-test/deep v1.0.8 // indirect
	golang.org/x/mod v0.6.0-dev.0.20220419223038-86c51ed26bb4 // indirect
	golang.org/x/tools v0.1.12 // indirect
)
