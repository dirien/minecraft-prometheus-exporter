package template

import (
	_ "embed"
	"text/template"
)

//go:embed index.html
var index string

func NewIndexTemplate() *template.Template {
	return template.Must(template.New("index").Parse(index))
}
