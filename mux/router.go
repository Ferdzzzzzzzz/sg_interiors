package mux

import (
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

	app.Handle(http.MethodGet, "/", s.render(page.Home()))
	app.Handle(http.MethodGet, "/about", s.render(page.About()))
	app.Handle(http.MethodGet, "/portfolio", s.render(page.Portfolio()))
	app.Handle(http.MethodGet, "/gallery", s.render(page.Gallery()))
	app.Handle(http.MethodGet, "/contact", s.render(page.Contact(page.ContactForm{})))
	app.Handle(http.MethodPost, "/contact", s.contactAction)

	return app
}
