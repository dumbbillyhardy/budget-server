package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dumbbillyhardy/budget-server/objects"

	context "github.com/dumbbillyhardy/budget-server/context"
)

// Handler which takes a request and writes a response.
type Handler func(http.ResponseWriter, *http.Request) error

// ErrorHandler handles the errors, logging n'at
func ErrorHandler(f Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			if _, t := err.(*objects.NotFoundError); t {
				ServeContent(w, r, err.Error(), http.StatusNotFound)
			} else {
				ServeContent(w, r, err.Error(), http.StatusInternalServerError)
			}
		}
	}
}

// RootHandlerFactory creates a handler that shows the server is running
func RootHandlerFactory(ctx context.Context) Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		return ServeContent(w, r, "<h1>Server is running.</h1>", http.StatusOK)
	}
}

// ServeContent writes the content to the response, setting headers n'at
func ServeContent(w http.ResponseWriter, r *http.Request, body string, statusCode int) error {
	logHeader := fmt.Sprintf(time.Now().Format("Jan 2 15:04:05 EST"))
	path := fmt.Sprintf(", request on path '%s' ", r.URL.Path)
	status := fmt.Sprintf("returned status code: %d", statusCode)

	if statusCode > 299 {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		errString := fmt.Sprintf(" with error: %s", body)
		status = status + errString
	}

	log.Println(logHeader + path + status)
	w.WriteHeader(statusCode)
	w.Write([]byte(body))

	return nil
}
