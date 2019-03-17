# Go Packaging
Demonstrate building a golang library that can be directory imported from github using only the standard git command line tools.    This is essential to reduce temptation of including shared libraries using cut and paste.

Testing ability to package GOLANG package for direct use by GO compiler.  

### Download sample repository for use by GO 

Make sure you have set GOPATH before attempting to  download and use packages.

```
  go get -u "github.com/joeatbayes/GoPackaging/sample_library"
```

### Use the Sample Library as imported package

```
import "github.com/joeatbayes/GOPackaging/sample_library"
```

### Make & run Sample Program

```
go get "github.com/joeatbayes/GoPackaging/sample_library"

go build test_git_import.go

.\test_git_import 
```

Use go get -u to fetch most recent changes of a library.

### Important Files

* [make_test_git_import.bat](make_test_git_import.bat) - Demonstrates using go get and go build commands.
* [test_git_import.go](test_git_import.go) - Sample program that requires package from github to build. 
* [MemFree.go](sample_library/MemFree.go) - Sample source in a package imported from github



### License:

Copyright 2018 Joseph Ellsworth  [MIT License](https://opensource.org/licenses/MIT) 

