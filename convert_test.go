package main

import (
	"bytes"
	"strings"
	"testing"
)

const pure = `<html>
  <head>
    <title>Hi</title>
  </head>
</html>`

func TestPureHtmlOutput(t *testing.T) {
	buf := bytes.NewBuffer(make([]byte, 0, 1000))
	Convert("templateName", strings.NewReader(pure), buf)

	expected := "package templates\n"
	expected += "\n"
	expected += "import (\n"
	expected += "  \"bytes\"\n"
	expected += "  \"fmt\"\n"
	expected += ")\n"
	expected += "\n"
	expected += "func templateName(locals templateNameLocals) []byte {\n"
	expected += "  buf := bytes.NewBuffer(make([]byte, 0, 275))\n"
	expected += "  fmt.Fprint(buf, `<html>\n"
	expected += "  <head>\n"
	expected += "    <title>Hi</title>\n"
	expected += "  </head>\n"
	expected += "</html>`);\n"
	expected += "  return buf.Bytes()\n"
	expected += "}\n"

	if buf.String() != expected {
		t.Errorf("Unexpected output:\n%s", buf.String())
	}
}

const code = `<html>
<% if this { %>
  <p>Hi</p>
<% } %>
</html>`

func TestCodeOutput(t *testing.T) {
	buf := bytes.NewBuffer(make([]byte, 0, 1000))
	Convert("templateName", strings.NewReader(code), buf)

	expected := "package templates\n"
	expected += "\n"
	expected += "import (\n"
	expected += "  \"bytes\"\n"
	expected += "  \"fmt\"\n"
	expected += ")\n"
	expected += "\n"
	expected += "func templateName(locals templateNameLocals) []byte {\n"
	expected += "  buf := bytes.NewBuffer(make([]byte, 0, 250))\n"
	expected += "  fmt.Fprint(buf, `<html>\n"
	expected += "`);if this {;fmt.Fprint(buf, `\n"
	expected += "  <p>Hi</p>\n"
	expected += "`);};fmt.Fprint(buf, `\n"
	expected += "</html>`);\n"
	expected += "  return buf.Bytes()\n"
	expected += "}\n"

	if buf.String() != expected {
		t.Errorf("Unexpected output:\n%s", buf.String())
	}
}

const printCode = `<html>
  <p><%= var %></p>
</html>`

func TestPrintCodeOutput(t *testing.T) {
	buf := bytes.NewBuffer(make([]byte, 0, 1000))
	Convert("templateName", strings.NewReader(printCode), buf)

	expected := "package templates\n"
	expected += "\n"
	expected += "import (\n"
	expected += "  \"bytes\"\n"
	expected += "  \"fmt\"\n"
	expected += ")\n"
	expected += "\n"
	expected += "func templateName(locals templateNameLocals) []byte {\n"
	expected += "  buf := bytes.NewBuffer(make([]byte, 0, 170))\n"
	expected += "  fmt.Fprint(buf, `<html>\n"
	expected += "  <p>`);fmt.Fprint(buf, var);fmt.Fprint(buf, `</p>\n"
	expected += "</html>`);\n"
	expected += "  return buf.Bytes()\n"
	expected += "}\n"

	if buf.String() != expected {
		t.Errorf("Unexpected output:\n%s", buf.String())
	}
}
