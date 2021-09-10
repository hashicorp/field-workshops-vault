package config

import "html/template"

//TPL is the HTML template container
var TPL *template.Template

func init() {
	TPL = template.Must(template.ParseGlob("templates/*.gohtml"))
}
