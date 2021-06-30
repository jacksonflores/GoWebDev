package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type hotel struct {
	Name    string
	Address string
	City    string
	Zip     int
	Region  string
}

type hotels []hotel

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	h := hotels{
		hotel{
			Name:    "coom resort",
			Address: "123 wiener road",
			City:    "Los Angeles",
			Zip:     69420,
			Region:  "Southern",
		},
		hotel{
			Name:    "tard house",
			Address: "456 dumptruck street",
			City:    "Gotham",
			Zip:     77777,
			Region:  "Northern",
		},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", h)
	if err != nil {
		log.Fatalln(err)
	}
}
