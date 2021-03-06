package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path"
	"unicode"
)

var templateName = flag.String("name", "", "Name of the template, this will be the function name")
var inputFile = flag.String("input", "", "Template path e.g. templates/example.cote")
var outputFile = flag.String("output", "", "Compiled output path e.g. templates/example.cote.go")

func main() {
	flag.Parse()

	in, name := inputReader()
	defer in.Close()

	out := outputWriter()
	defer out.Close()

	if err := Convert(name, in, out); err != nil {
		panic(err)
	}
}

// Reads the input file from the -input flag or Stdin. Exits if
// no template name given when using stdin, or no input path given.
func inputReader() (f io.ReadCloser, name string) {
	if isUsingStdin() {
		if *templateName == "" {
			fmt.Fprintln(os.Stderr, "Template name not specified. Use -name flag.")
			os.Exit(1)
		}

		f = os.Stdin
		name = *templateName
	} else {
		if *inputFile == "" {
			fmt.Fprintln(os.Stderr, "Input path not specified. Use -input flag or Stdin")
			os.Exit(1)
		}

		var err error
		f, err = os.Open(*inputFile)
		if err != nil {
			panic(err)
		}

		name = *templateName
		if name == "" {
			name = templateNameFromFilePath(*inputFile)
		}
	}

	return
}

// Extract a sensible template name from the file path
func templateNameFromFilePath(p string) string {
	name := path.Base(p)
	if ext := path.Ext(name); ext != "" {
		name = name[:len(name)-len(ext)]
	}
	return string(unicode.ToUpper(rune(name[0]))) + name[1:]
}

func isUsingStdin() bool {
	s, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	return s.Size() > 0
}

// Gets an output writer, either a file or Stdout
func outputWriter() io.WriteCloser {
	if *outputFile == "" {
		return os.Stdout
	}

	w, err := os.Create(*outputFile)
	if err != nil {
		panic(err)
	}
	return w
}
