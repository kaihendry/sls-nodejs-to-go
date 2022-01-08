package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"time"

	"github.com/apex/gateway/v2"
	"github.com/apex/log"
	jsonhandler "github.com/apex/log/handlers/json"
	"github.com/apex/log/handlers/text"
)

//go:embed templates
var tmpl embed.FS

func main() {
	t, err := template.ParseFS(tmpl, "templates/*.html")
	if err != nil {
		log.WithError(err).Fatal("Failed to parse templates")
	}

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "text/html")
		err = t.ExecuteTemplate(rw, "index.html", struct {
			Now time.Time
		}{
			time.Now(),
		})
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			log.WithError(err).Fatal("Failed to execute templates")
		}
	})

	log.SetHandler(jsonhandler.Default)

	if _, ok := os.LookupEnv("AWS_LAMBDA_FUNCTION_NAME"); ok {
		err = gateway.ListenAndServe("", nil)
	} else {
		log.SetHandler(text.Default)
		err = http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil)
	}
	log.WithError(err).Fatal("error listening")
}
