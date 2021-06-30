package main

import (
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	x := 42

	tpl.ExecuteTemplate(os.Stdout, "index.gohtml", x)
}
