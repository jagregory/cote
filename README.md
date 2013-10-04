# cote - COmpiled TEmplates for Go

Cote is a template compiler for Go. Use it to precompile HTML (or other) templates and embed them in your Go packages.

## Install

```bash
go get github.com/jagregory/cote
```

Coat has an executable `coat` which you'll need to use, make sure it's available on your path (or you know where it is).

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

## Usage

The best approach to using Cote is to treat it as a pre-compile step.

Put your Coat templates in your package, then run Coat over each one specifying a template name. Do this whenever you change a template.

## Examples

*templates/example.coat*

    <p>Hi <%= locals.Name %></p>

Which can be compiled like so (this'll get easier soon):

```bash
cat templates/example.coat | ./coat -name=example > templates/example.coat.go
```

This will produce a template called `example`.

```go
example(locals exampleLocals) []byte {
  ...
}
```

You will need to define a `*Locals` struct for each template you use. E.g.

```go
struct exampleLocals {
  Name string
}
```

Finally, you can use your template:

```go
func(w http.ResponseWriter, r *http.Request) {
  html := example(exampleLocals{ Name: "James" })
  w.Write(html)
}
```

    <p>Hi James</p>