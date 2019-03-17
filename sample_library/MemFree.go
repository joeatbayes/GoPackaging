/*
 (C) Copyright 2019 Joe Ellsworth
 MIT LICENSE  */

// Package stringutil contains utility functions for working with strings.
// Include this package using: 
//    import "github.com/GOPackaging/sample_library"
// Download the package with import 
//  go get github.com/GOPackaging

package sample_library

import (
  "runtime"
)

const BytesPerMBF = 1048576.0

// Reverse returns its argument string reversed rune-wise left to right.
func MBUsed() float64 {
  var m runtime.MemStats
  runtime.ReadMemStats(&m)
  return float64(m.TotalAlloc) / BytesPerMBF 
}

