package mux

import (
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"sg/core/alert"
	"sg/tmpls/page"

	"github.com/Rockup-Consulting/std/core/mid"
	"github.com/Rockup-Consulting/std/core/web"
)

type service struct {
	l       *log.Logger
	alerter alert.Alerter
}

func Init(
	l *log.Logger,
	alerter alert.Alerter,
	staticFS fs.FS,
) http.Handler {
	app := web.NewApp(
		mid.Log(l),
		mid.CatchErr(l),
		mid.CatchPanic(),
		mid.TryGzip,
	)

	s := service{
		l:       l,
		alerter: alerter,
	}

	staticHandler := web.StaticHandler(staticFS, l, web.DefaultStaticCacheSeconds, true)
	app.HandleStd(http.MethodGet, "/assets/*", staticHandler.ServeHTTP)

	app.Handle(http.MethodGet, "/", func(w http.ResponseWriter, r *http.Request) error {
		if r.URL.Path != "/" {
			fmt.Fprint(w, "<h1>Could not find this route</h1>")
			return nil
		}

		return page.Home().Render(r.Context(), w)
	})

	app.Handle(http.MethodGet, "/about", s.render(page.About()))
	app.Handle(http.MethodGet, "/portfolio", s.render(page.Portfolio()))
	app.Handle(http.MethodGet, "/renders", s.render(page.Renders()))
	app.Handle(http.MethodGet, "/renders/{slug}", func(w http.ResponseWriter, r *http.Request) error {
		switch r.PathValue("slug") {
		case page.RenderPage.AutumnalHarmony.Slug:
			return page.Render(page.RenderPage.AutumnalHarmony).Render(r.Context(), w)
		case page.RenderPage.TimelessTranquility.Slug:
			return page.Render(page.RenderPage.TimelessTranquility).Render(r.Context(), w)
		}

		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)

		return nil
	})
	app.Handle(http.MethodGet, "/boards", s.render(page.Boards()))
	app.Handle(http.MethodGet, "/gallery", s.render(page.Gallery()))
	app.Handle(http.MethodGet, "/contact", s.render(page.Contact(page.ContactForm{})))
	app.Handle(http.MethodPost, "/contact", s.contactAction)

	return app
}
