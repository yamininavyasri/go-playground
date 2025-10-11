//go:build linux
// +build linux

/*
Build Tags (or build constraints) are special comments that control when a Go file should be included in a build.
They allow you to include/exclude files for specific OS, architecture, or custom conditions.
Useful for:
Platform-specific code (windows, linux, darwin)
Debug vs production builds
Feature flags
Syntax
go:build <expression>
+build <expression>  // older syntax, optional in Go 1.17+
*/

package main

import "fmt"

func main() {
	fmt.Println("This program runs only on Linux")
}
