package config

import (
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var TplAuthenticated *template.Template
var TplUnauthenticated *template.Template

var TplUnauthenticatedAbout *template.Template
var Tpl *template.Template

func init() {
	//TplAuthenticated = template.Must(template.ParseGlob("./frontend/client/view/templates/authenticated/*.html"))
	//TplUnauthenticated = template.Must(template.ParseGlob("./frontend/client/view/templates/unauthenticated/*.html"))
	//TplUnauthenticatedAbout = template.Must(template.ParseGlob("./frontend/client/view/templates/unauthenticated/about/*.html"))

	Tpl = template.Must(ParseTemplates(), nil)

}

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
