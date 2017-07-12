package view

import (
	"bytes"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
	"github.com/tdewolff/minify/html"
	"github.com/tdewolff/minify/js"
)

var (
	tpIndex         = parseTemplate("root.tmpl", "index.tmpl")
	tpNews          = parseTemplate("root.tmpl", "news.tmpl")
	tpAdminLogin    = parseTemplate("root.tmpl", "admin/login.tmpl")
	tpAdminRegister = parseTemplate("root.tmpl", "admin/register.tmpl")
	tpAdminList     = parseTemplate("root.tmpl", "admin/list.tmpl")
	tpAdminCreate   = parseTemplate("root.tmpl", "admin/create.tmpl")
	tpAdminEdit     = parseTemplate("root.tmpl", "admin/edit.tmpl")
)

var m = minify.New()

const templateDir = "template"

func init() {
	m.AddFunc("text/html", html.Minify)
	m.AddFunc("text/css", css.Minify)
	m.AddFunc("text/javascript", js.Minify)
}

func joinTemplateDir(files ...string) []string {
	r := make([]string, len(files))
	for i, f := range files {
		r[i] = filepath.Join(templateDir, f)
	}
	return r
}

func parseTemplate(files ...string) *template.Template {
	t := template.New("")
	t.Funcs(template.FuncMap{})
	_, err := t.ParseFiles(joinTemplateDir(files...)...)
	if err != nil {
		panic(err)
	}
	t = t.Lookup("root")
	return t
}

func render(t *template.Template, w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	buf := bytes.Buffer{}
	err := t.Execute(&buf, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	m.Minify("text/html", w, &buf)
}
