package views

import (
	"fmt"
	"html/template"
	"path/filepath"
)

// Global vars to construct filepaths
var (
	LayoutDir   string = "views/layouts/"
	TemplateExt string = ".gohtml"
)

// View struct for new layout files to be auto implemented in templates.
type View struct {
	Template *template.Template
	Layout   string
}

// layoutFiles to glob the layout directory with the .gohtml files
func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err != nil {
		panic(err)
	}
	return files
}

// NewView func to append files returned layoutFiles func and then
// parse those files.
func NewView(layout string, files ...string) *View {
	files = append(files, layoutFiles()...)
	t, err := template.ParseFiles(files...)
	if err != nil {
		fmt.Printf("Error parsing template files: %v", err)
	}

	return &View{
		Template: t,
		Layout:   layout,
	}
}
