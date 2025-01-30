package renderers

import (
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
	"path/filepath"
)

type HTMLTemplate struct {
	Dir string
	Ext string
}

func (t *HTMLTemplate) Render(w io.Writer, n string, d interface{}, c echo.Context) error {
	path := filepath.Join(t.Dir + "/" + n + t.Ext)
	a, err := template.ParseFiles(path)
	if err != nil {
		return err
	}
	return a.Execute(w, d)
}
