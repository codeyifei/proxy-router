package middleware

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
)

func AccessLog(next http.Handler) http.Handler {
	startTime := time.Now()
	return handlers.CustomLoggingHandler(log.Writer(), next, func(w io.Writer, params handlers.LogFormatterParams) {
		_, _ = fmt.Fprintf(w, "method: %s path: %s status: %d time: %s, %s\n",
			params.Request.Method,
			params.URL.RequestURI(),
			params.StatusCode,
			time.Now().Sub(startTime),
			params.TimeStamp)
	})
}
