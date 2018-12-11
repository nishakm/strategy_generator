// Business Strategy Generator written in Go
// Inspired by http://blog.gardeviance.org/2014/07/a-quick-route-to-building-strategy.html
// Doubles as a learning guide for Go (so there are more comments than actual lines of code)
// Copyright (c) Nisha Kumar (@nishakm) All Rights Reserved
// SPDX-License-Identifier: BSD-2-Clause

// Entry point of any Go program is the main package
// The file name can be called anything
// I am calling it 'main.go' for clarity
package main

// Import other golang packages here. Names are in ""
import (
	"fmt"
/*
If you either put the pkg in ~/go/src, or set GOPATH to include the parent dir,
the relative form shoud work.  The relative form is generally preferred.
*/
	"strategy_generator/pkg"
)

// this is the main function
func main() {
	statement := pkg.Generate()

	fmt.Println(statement)
}
