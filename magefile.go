// +build mage

package main

import (
	"fmt"
	"os/exec"
)

// Default target to run when none is specified
// If not set, running mage will list available targets
// var Default = Build

// Fmt is the format step
func Fmt() error {
	fmt.Println("Formatting...")
	cmd := exec.Command("go", "fmt", "./...")
	return cmd.Run()
}

// Lint is the lint step
func Lint() error {
	Fmt()
	fmt.Println("Linting...")
	cmd := exec.Command("golint", "./...")
	return cmd.Run()
}

// Vet is the analysis step
func Vet() error {
	Lint()
	fmt.Println("Analizing...")
	cmd := exec.Command("go", "fmt", "./...")
	return cmd.Run()
}

// Test is the test step
func Test() error {
	fmt.Println("Testing...")
	cmd := exec.Command("go", "test")
	return cmd.Run()
}

// Build is the build step that execute all other
func Build() error {
	Vet()
	Test()
	fmt.Println("Building...")
	cmd := exec.Command("go", "build", "-o", "bin/hello")
	return cmd.Run()
}
