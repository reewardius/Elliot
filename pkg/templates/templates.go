package templates

// === IMPORTS ===

import "html/template"

// === PUBLIC METHODS ===

// NewTemplates generates a new structure associated with the HTML templates served on the website
func NewTemplates() *template.Template {
	templates := template.Must(template.ParseGlob("frontend/views/*.html"))
	return template.Must(templates.ParseGlob("frontend/views/modules/*.html"))
}
