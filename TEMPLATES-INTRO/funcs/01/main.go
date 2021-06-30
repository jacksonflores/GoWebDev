package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

type homie struct {
	Fname string
	Lname string
}

var tpl *template.Template

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	if len(s) >= 3 {
		s = s[:3]
	}
	return s
}

func main() {
	j := homie{
		Fname: "jackson",
		Lname: "flores",
	}

	d := homie{
		Fname: "davis",
		Lname: "taylor",
	}

	homies := []homie{j, d}

	data := struct {
		People []homie
	}{
		homies,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}
