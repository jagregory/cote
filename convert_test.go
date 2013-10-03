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

	expected := "package template\n"
	expected += "\n"
	expected += "func templateName() {\n"
	expected += "  fmt.Print(`<html>\n"
	expected += "  <head>\n"
	expected += "    <title>Hi</title>\n"
	expected += "  </head>\n"
	expected += "</html>`);\n"
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

	expected := "package template\n"
	expected += "\n"
	expected += "func templateName() {\n"
	expected += "  fmt.Print(`<html>\n"
	expected += "`);if this {;fmt.Print(`\n"
	expected += "  <p>Hi</p>\n"
	expected += "`);};fmt.Print(`\n"
	expected += "</html>`);\n"
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

	expected := "package template\n"
	expected += "\n"
	expected += "func templateName() {\n"
	expected += "  fmt.Print(`<html>\n"
	expected += "  <p>`);fmt.Print(var);fmt.Print(`</p>\n"
	expected += "</html>`);\n"
	expected += "}\n"

	if buf.String() != expected {
		t.Errorf("Unexpected output:\n%s", buf.String())
	}
}
