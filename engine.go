package templateManagerGin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"

	TM "github.com/paul-norman/go-template-manager"
)

// Convenience type allowing any variables types to be passed in
type Params map[string]any

// Struct wrapper allowing override of the Render function
type Engine struct {
	*TM.TemplateManager
}

// Creates a new Renderer instance
func Init(directory string, extension string) *Engine {
	return &Engine{ TM.Init(directory, extension) }
}

// Renders a single template
func (e *Engine) Instance(name string, data any) render.Render {
	return &Renderer{ name, parseData(data), e.TemplateManager }
}

// Struct wrapper replacing `render.HTML`
type Renderer struct {
	name string
	params TM.Params
	tm *TM.TemplateManager
}

// Called by Gin's contextual HTML function
func (r *Renderer) Render(w http.ResponseWriter) error {
	r.WriteContentType(w)

	return r.tm.Render(r.name, r.params, w)
}

// Required to fulfil Gin's Interface requirements
func (r *Renderer) WriteContentType(w http.ResponseWriter) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = []string{"text/html; charset=utf-8"}
	}
}

// Converts generic variables to TM variables
func parseData(binding any) TM.Params {
	if binding == nil {
		return TM.Params{}
	}

	if old, ok := binding.(TM.Params); ok {
		return old
	}

	if old, ok := binding.(map[string]any); ok {
		return TM.Params(old)
	}

	if old, ok := binding.(Params); ok {
		return TM.Params(old)
	}

	if old, ok := binding.(gin.H); ok {
		return TM.Params(old)
	}

	return TM.Params{}
}