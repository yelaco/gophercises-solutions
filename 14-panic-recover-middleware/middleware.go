package main

import (
	"log"
	"net/http"
	"runtime/debug"
)

type middleware func(http.Handler) http.Handler

func createStack(xs ...middleware) middleware {
	return func(next http.Handler) http.Handler {
		for i := len(xs) - 1; i >= 0; i-- {
			x := xs[i]
			next = x(next)
		}
		return next
	}
}

func panicRecover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bw := NewBufferedResponseWriter(w)
		defer func() {
			if r := recover(); r != nil {
				log.Println(r)
				debug.PrintStack()
				bw.Reset()
				bw.WriteHeader(500)
				var buf []byte
				if mode == DEV_MODE {
					buf = debug.Stack()
				} else {
					buf = []byte("Something went wrong")
				}
				w.Write(buf)
			}
		}()

		// Call the next handler
		next.ServeHTTP(bw, r)

		// Flush the response if no panic occurred
		bw.Flush()
	})
}
