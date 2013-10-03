package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"regexp"
)

var writeBlock = regexp.MustCompile(`<%=\s*(.*?)\s*%>`)
var execBlock = regexp.MustCompile(`<%\s*(.*?)\s*%>`)

func Convert(name string, in io.Reader, out io.Writer) error {
	bytes, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}

	body := string(bytes)
	body = writeBlock.ReplaceAllString(body, "`);fmt.Print($1);fmt.Print(`")
	body = execBlock.ReplaceAllString(body, "`);$1;fmt.Print(`")

	fmt.Fprint(out, "package template\n")
	fmt.Fprint(out, "\n")
	fmt.Fprintf(out, "func %s() {\n", name)
	fmt.Fprintf(out, "  fmt.Print(`%s`);\n", body)
	fmt.Fprint(out, "}\n")

	return nil
}
