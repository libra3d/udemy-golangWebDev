package main

import (
	"html/template"
	"os"
	"log"
)

type Page struct {
	Title string
	Heading string
	Input string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	home := Page{
		Title: "Escaped",
		Heading: "Danger is escaped with html/template",
		Input: `<script>alert("Howdy!");</script>`,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", home)
	if err != nil {
		log.Fatalln(err)
	}
}
