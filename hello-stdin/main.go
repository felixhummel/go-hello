package main

import (
	"os"
	"io/ioutil"
	"fmt"
)

const template = `STDIN data:
----
%v
----
`

func main() {
	data, _ := ioutil.ReadAll(os.Stdin)
	os.Stdout.WriteString("Hello STDOUT!\n")
	os.Stderr.WriteString("Hello STDERR!\n")
	fmt.Printf(template, string(data))
}
