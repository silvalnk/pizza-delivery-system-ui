package http

import (
	"encoding/json"
	"html/template"

	"github.com/gin-gonic/gin"
)

func LoadTemplates(r *gin.Engine, pattern string) error {
	funcMap := template.FuncMap{
		"add": func(a, b int) int { return a + b },
		"json": func(v interface{}) template.JS {
			b, _ := json.Marshal(v)
			return template.JS(b)
		},
	}
	tmpl, err := template.New("").Funcs(funcMap).ParseGlob(pattern)
	if err != nil {
		return err
	}
	r.SetHTMLTemplate(tmpl)
	return nil
}
