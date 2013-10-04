# cote - COmpiled TEmplates for Go

Cote is a bare-bones templating language for Go. It does not have an interpreter, or a custom language. It simply mixes Go with HTML (or your other text-based language of choice), and compiles them to pure Go. *Cote is a compile-time template language.*

Templates are written in HTML with Go embedded with `<% %>` and `<%= %>` tags. The `cote` utility converts these templates into Go code files. The generated code is then compiled into your package, and usable without the original HTML files.

## Install

```bash
go get github.com/jagregory/cote
```

Cote has an executable `cote` which you'll need to use, make sure it's available on your `PATH` (or you know where it is).

## Rationalè / Q&A

**Why not html/template?** I did not enjoy using html/template. For a language designed by the people responsible for Go, a lovely elegant language, it's obtuse and bizarre.

**But, logicless templates!** I'm a big proponent of logicless templates; however, I prefer to exclude the logic myself rather than be hamstrung by the language. For example, html/templates only got an equality operator in Go 1.2.

**So Cote templates are compiled?** Correct. They're embedded in your package as Go code. At no point during the running of your code will anything be interpreted. This is obviously quite fast, and your templates are statically typed.

**But doesn't that mean you have to recompile your code every time you make a template change?** Yes. I would prefer to not do it, but Go doesn't have runtime code generation so there aren't any other options besides writing a Go interpreter. At least the compiler is fast.

### Pros

  * Fast
  * No runtime overhead
  * No dependencies on files
  * Can use any valid Go code in your templates
  * Statically typed/compile-time checked

### Cons

  * Have to regenerate and recompile for every change
  * Can use any valid Go code in your templates!

## Examples

```html
<!DOCTYPE html>
<html>
  <head>
    <title><%= locals.Title %></title>
  </head>
  <body>
    <h1><%= locals.Title %></h1>
    <% if locals.IsHome { %>
      <p>Welcome home</p>
    <% } %>
  </body>
</html>
```

```go
import "templates"

func(w http.ResponseWriter, r *http.Request) {
  html := templates.Home(HomeLocals{
    Title: "Our Site",
    IsHome: true,
  })
  w.Write(html)
}
```

## Usage

The best approach to using Cote is to treat it as a pre-compile step.

Put your Cote templates in your package, then run Cote over each one specifying a template name. Do this whenever you change a template.

    cat yourtemplate.cote | ./cote -name=yourtemplate > yourtemplate.cote.go

Or

    ./cote -input=yourtemplate.cote -output=yourtemplate.cote.go

A bit more detailed. If you had a template named *templates/example.cote*.

    <p>Hi <%= locals.Name %></p>

You could compile the template by either piping its content into `cote` and redirecting the output to a file.

e.g. `cat templates/example.cote | ./cote -name=Example > templates/example.cote.go`

Or alternatively, you can use `cote` with `-input` and `-output` flags.

e.g. `./cote -input=templates/example.cote -output=templates/example.cote.go`

Either approach will produce a template named `Example`, you can override this with the `-name` flag. Whatever name you use needs to be a valid Go method name, as it will be the method which you call to render the template.

```go
package templates

func Example(locals exampleLocals) []byte {
  ...
}
```

All templates take a `locals` structure, which you can use within the template to access any variables you need. You will need to declare the structure yourself, with a name of `*templateName*Locals`, e.g. `struct ExampleLocals { Name string }`

Finally, you can use your template by calling the template method with a locals instance.

```go
import "templates"

func(w http.ResponseWriter, r *http.Request) {
  html := templates.Example(ExampleLocals{ Name: "James" })
  w.Write(html)
}
```

This would produce: `<p>Hi James</p>`.

## Known issues

  * No way to handle extra imports (e.g. `include "time"`)
