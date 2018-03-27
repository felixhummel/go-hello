package main

import (
	"os"
)

func main() {
	os.Stdout.WriteString("Hello STDOUT!\n")
	os.Stderr.WriteString("Hello STDERR!\n")
}
