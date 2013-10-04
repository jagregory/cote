package templates

import (
  "bytes"
  "fmt"
)

func ExampleTemplate(locals ExampleTemplateLocals) []byte {
  buf := bytes.NewBuffer(make([]byte, 0, 1925))
  fmt.Fprint(buf, `<!DOCTYPE html>
<html>
  <head>
    <title>News</title>
  </head>
  <body>
    <h1>Today</h1>
  `);if len(locals.Submissions) == 0 {;fmt.Fprint(buf, `
    <p>Nothing has been reported yet.</p>
  `);} else {;fmt.Fprint(buf, `
    <ol>
      `);for _, s := range locals.Submissions {;fmt.Fprint(buf, `
        <li>
          <a href="`);fmt.Fprint(buf, s.Url);fmt.Fprint(buf, `">`);fmt.Fprint(buf, s.Title);fmt.Fprint(buf, `</a>
        </li>
      `);};fmt.Fprint(buf, `
    </ol>
  `);};fmt.Fprint(buf, `
  </body>
</html>`);
  return buf.Bytes()
}
