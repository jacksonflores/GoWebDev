package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("beep.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "beep.gohtml", "this is some data")
	if err != nil {
		log.Fatalln(err)
	}
}
