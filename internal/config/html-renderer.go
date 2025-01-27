package config

import (
	"context"
	"errors"
	"net/http"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin/render"
)

//Custom renderer to use templ generated functions
//Implements the gin HTML renderer interface
type TemplRender struct {
	Code int
	Data templ.Component
}

// Renders the html template from the templ component
func (t TemplRender) Render(w http.ResponseWriter) error {
	t.WriteContentType(w)
	w.WriteHeader(t.Code)
	if t.Data != nil {
		return t.Data.Render(context.Background(), w)
	}
	return errors.New("renderer should always have a component when rendering")
}

func (t TemplRender) WriteContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
}

func (t *TemplRender) Instance(name string, data interface{}) render.Render {
	if templData, ok := data.(templ.Component); ok {
		return &TemplRender{
			Code: http.StatusOK,
			Data: templData,
		}
	}
	return nil
}

