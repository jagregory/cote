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
	body = writeBlock.ReplaceAllString(body, "`);fmt.Fprint(buf, $1);fmt.Fprint(buf, `")
	body = execBlock.ReplaceAllString(body, "`);$1;fmt.Fprint(buf, `")

	fmt.Fprint(out, "package templates\n")
	fmt.Fprint(out, "\n")
	fmt.Fprint(out, "import (\n")
	fmt.Fprint(out, "  \"bytes\"\n")
	fmt.Fprint(out, "  \"fmt\"\n")
	fmt.Fprint(out, ")\n")
	fmt.Fprint(out, "\n")
	fmt.Fprintf(out, "func %s(locals %sLocals) []byte {\n", name, name)
	fmt.Fprintf(out, "  buf := bytes.NewBuffer(make([]byte, 0, %d))\n", len(bytes)*5)
	fmt.Fprintf(out, "  fmt.Fprint(buf, `%s`);\n", body)
	fmt.Fprint(out, "  return buf.Bytes()\n")
	fmt.Fprint(out, "}\n")

	return nil
}
