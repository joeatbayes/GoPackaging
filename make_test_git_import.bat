set GOPATH=%cd%
::  Run get-dependencies.bat before running this
::  script.
go get "github.com/joeatbayes/GoPackaging/sample_library"

go build test_git_import.go
