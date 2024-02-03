package mux

import (
	"net/http"

	"github.com/Rockup-Consulting/std/core/web"
	"github.com/a-h/templ"
)

func (s service) render(cmp templ.Component) web.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		return cmp.Render(r.Context(), w)
	}
}
