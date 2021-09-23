package utils

import (
	"html/template"
	"net/http"
)

var templates *template.Template

// CarregarTemplates insere template html na variável templetes
func CarregarTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
}

//ExecutarTemplates renderizar uma pagina html na tela
func ExecutarTemplates(w http.ResponseWriter, template string, dados interface{}) {
	templates.ExecuteTemplate(w, template, dados)
}
