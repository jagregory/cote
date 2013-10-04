package templates

import (
	"io/ioutil"
	"os"
	"os/exec"
	"testing"
)

func TestExampleWithNoSubmissions(t *testing.T) {
	locals := ExampleTemplateLocals{}
	output := string(ExampleTemplate(locals))
	expected := `<!DOCTYPE html>
<html>
  <head>
    <title>News</title>
  </head>
  <body>
    <h1>Today</h1>
  
    <p>Nothing has been reported yet.</p>
  
  </body>
</html>`

	if output != expected {
		d, _ := diff([]byte(output), []byte(expected))
		t.Errorf("Unexpected output.\nDiff:\n%s", d)
	}
}

func TestExampleWithSubmissions(t *testing.T) {
	locals := ExampleTemplateLocals{
		Submissions: []Submission{
			Submission{"http://example.com", "Hello World"},
		},
	}
	output := string(ExampleTemplate(locals))
	expected := `<!DOCTYPE html>
<html>
  <head>
    <title>News</title>
  </head>
  <body>
    <h1>Today</h1>
  
    <ol>
      
        <li>
          <a href="http://example.com">Hello World</a>
        </li>
      
    </ol>
  
  </body>
</html>`

	if output != expected {
		d, _ := diff([]byte(output), []byte(expected))
		t.Errorf("Unexpected output.\nDiff:\n%s", d)
	}
}

// diff two strings, ripped from gofmt
func diff(actual, expected []byte) (data []byte, err error) {
	f1, err := ioutil.TempFile("", "actual")
	if err != nil {
		return
	}
	defer os.Remove(f1.Name())
	defer f1.Close()

	f2, err := ioutil.TempFile("", "expected")
	if err != nil {
		return
	}
	defer os.Remove(f2.Name())
	defer f2.Close()

	f1.Write(actual)
	f2.Write(expected)

	data, err = exec.Command("diff", "-u", f1.Name(), f2.Name()).CombinedOutput()
	if len(data) > 0 {
		// diff exits with a non-zero status when the files don't match.
		// Ignore that failure as long as we get output.
		err = nil
	}
	return

}
