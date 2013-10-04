package templates

type Submission struct {
	Url, Title string
}

type ExampleTemplateLocals struct {
	Submissions []Submission
}
