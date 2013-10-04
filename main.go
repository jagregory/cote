package main

import (
	"flag"
	"fmt"
	"os"
)

var templateName = flag.String("name", "", "Name of the template, this will be the function name")

func main() {
	flag.Parse()

	if templateName == nil || *templateName == "" {
		fmt.Fprintln(os.Stderr, "Template name not specified. Use -name flag.")
		os.Exit(1)
	}

	if err := Convert(*templateName, os.Stdin, os.Stdout); err != nil {
		panic(err)
	}
}
