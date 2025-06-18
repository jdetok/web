package main

// modified tutorial from
// https://blog.questionable.services/article/guide-logging-middleware-go/

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"time"
)

// responseWriter is a minimal wrapper for http.ResponseWriter that allows the
// written HTTP status code to be captured for logging.
type responseWriter struct {
  http.ResponseWriter
  status      int
  wroteHeader bool
}

func wrapResponseWriter(w http.ResponseWriter) *responseWriter {
  return &responseWriter{ResponseWriter: w}
}

func (rw *responseWriter) Status() int {
  return rw.status
}

func (rw *responseWriter) WriteHeader(code int) {
  if rw.wroteHeader {
    return
  }

  rw.status = code
  rw.ResponseWriter.WriteHeader(code)
  rw.wroteHeader = true

  //return
}

// LoggingMiddleware logs the incoming HTTP request & its duration.
func (app *application) LoggingMiddleware() func(http.Handler) http.Handler {
  return func(next http.Handler) http.Handler {
    fn := func(w http.ResponseWriter, r *http.Request) {
      defer func() {
        if err := recover(); err != nil {
            w.WriteHeader(http.StatusInternalServerError)
		  	fmt.Printf("err: %v | trace: %v", err, debug.Stack())
        }
      }()

      start := time.Now()
      wrapped := wrapResponseWriter(w)
      next.ServeHTTP(wrapped, r)
	  fmt.Printf("status: %v | method: %v | path: %v | duration: %v |", 
	  wrapped.status, r.Method, r.URL.EscapedPath(), time.Since(start))
    }

    return http.HandlerFunc(fn)
  }
}
