package config

import (
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Tpl creates a pointer varriable used to execute templates
var Tpl *template.Template

// init runs when the program executes
func init() {
	Tpl = template.Must(ParseTemplates(), nil)
}

// ParseTemplates parses all html documents in the location
// and returns a pointer of type Template
func ParseTemplates() *template.Template {
	templ := template.New("")
	err := filepath.Walk("./frontend/client/view/templates/", func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, ".html") {
			_, err = templ.ParseFiles(path)
			if err != nil {
				log.Println(err)
			}
		}

		return err
	})

	if err != nil {
		panic(err)
	}

	return templ
}
