package tech

import (
	"html/template"
	"io"
)

type ViewData struct {
	Items []*Object
}

func RenderListHTML(wr io.Writer, list []*Object) error {
	t, err := template.ParseFiles("./res/templates/allTechList")
	if err != nil {
		return err
	}
	return t.Execute(wr, ViewData{Items: list})
}
