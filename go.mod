module github.com/yehan2002/fastbytes/v2

go 1.17

// This version has a broken mod file
retract v2.0.0

// This version has incorrect import paths and will behave like v1
retract v2.0.1

require (
	github.com/mmcloughlin/avo v0.2.0
	github.com/yehan2002/errors v1.1.1
	github.com/yehan2002/is/v2 v2.1.0
	golang.org/x/sys v0.0.0-20210403161142-5e06dd20ab57
)

require (
	github.com/go-test/deep v1.0.7 // indirect
	golang.org/x/mod v0.3.0 // indirect
	golang.org/x/tools v0.1.0 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
)
