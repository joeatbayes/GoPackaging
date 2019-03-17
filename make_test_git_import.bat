set GOPATH=%cd%

go get -u "github.com/joeatbayes/GoPackaging/sample_library"

go build test_git_import.go

.\test_git_import
