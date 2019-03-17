/* (C) Copyright 2019 Joe Ellsworth
 MIT LICENSE  
  Download use of library for use by GO get
    go get "github.com/joeatbayes/GoPackaging/sample_library"
 */
 
package main

import (
  "fmt"  
)

import "github.com/joeatbayes/GOPackaging/sample_library"

func main() {
  
  fmt.Printf("In test_git_import.go  Megabytes Used=%.3f", 
sample_library.MBUsed())
}
