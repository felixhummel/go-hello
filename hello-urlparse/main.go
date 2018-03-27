package main

import (
	"fmt"
	"log"
	"net/url"
	"reflect"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	URL   = kingpin.Arg("URL", "URL to parse").Required().String()
	FIELD = kingpin.Arg("FIELD", "Field of go's url.URL to return").Required().String()
)

func main() {
	kingpin.Version("0.0.1")
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()
	u, err := url.Parse(*URL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reflect.ValueOf(*u).FieldByName(*FIELD))
}
