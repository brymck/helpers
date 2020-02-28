package webapp

import (
	"fmt"
	"net/http"
	"runtime"
	"runtime/debug"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/brymck/helpers/env"
)

type WebApp interface {
	Routes() http.Handler
}

func Serve(app WebApp) {
	port := env.GetPort("8080")

	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      app.Routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Infof("starting server on %s", port)
	log.Fatal(srv.ListenAndServe())
}

func ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	if _, file, line, ok := runtime.Caller(2); ok {
		log.WithFields(log.Fields{
			"file": file,
			"line": line,
		}).Error(trace)
	} else {
		log.Error(trace)
	}

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func ClientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func NotFound(w http.ResponseWriter) {
	ClientError(w, http.StatusNotFound)
}
