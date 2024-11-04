package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type templateData struct {
	Data map[string]any
}

func (app *application) buildTemplateFromDisc(t string) (*template.Template, error) {
	templateSlice := []string{
		"./templates/base.layout.gohtml",
		"./templates/partials/header.partials.gohtml",
		"./templates/partials/footer.partials.gohtml",
		fmt.Sprintf("./templates/%s", t),
	}
	tmpl, err := template.ParseFiles(templateSlice...)
	if err != nil {
		return nil, err
	}

	app.templateMap[t] = tmpl

	return tmpl, nil
}

func (app *application) render(ctx *gin.Context, t string, td *templateData) {
	var tmpl *template.Template

	// use cache
	if app.config.useCache {
		if templateFromMap, ok := app.templateMap[t]; ok {
			tmpl = templateFromMap
		}
	}

	if tmpl == nil {
		newTemplate, err := app.buildTemplateFromDisc(t)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("Building template from disk")
		tmpl = newTemplate
	}

	if td == nil {
		td = &templateData{}
	}

	if err := tmpl.ExecuteTemplate(ctx.Writer, t, td); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}
}
