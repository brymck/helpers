package webapp

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/brymck/helpers/env"
)

type WebApp interface {
	routes() http.Handler
}

func Serve(app WebApp) {
	port := env.GetPort("8080")

	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Infof("starting server on %s", port)
	log.Fatal(srv.ListenAndServe())
}
