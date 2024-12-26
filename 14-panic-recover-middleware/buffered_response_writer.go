package main

import (
	"bytes"
	"log"
	"net/http"
)

type BufferedResponseWriter struct {
	writer  http.ResponseWriter
	buffer  bytes.Buffer
	status  int
	written bool
}

func NewBufferedResponseWriter(w http.ResponseWriter) *BufferedResponseWriter {
	return &BufferedResponseWriter{
		writer: w,
		buffer: bytes.Buffer{},
	}
}

func (bw *BufferedResponseWriter) WriteHeader(statusCode int) {
	if bw.written {
		return
	}
	bw.status = statusCode
}

func (bw *BufferedResponseWriter) Header() http.Header {
	return bw.writer.Header()
}

func (bw *BufferedResponseWriter) Write(data []byte) (int, error) {
	return bw.buffer.Write(data)
}

func (bw *BufferedResponseWriter) Flush() {
	if bw.written {
		return
	}
	if bw.status != 0 {
		bw.writer.WriteHeader(bw.status)
	}
	_, err := bw.buffer.WriteTo(bw.writer)
	if err != nil {
		log.Printf("Failed to write response: %v", err)
	}
	bw.written = true
}

func (bw *BufferedResponseWriter) Reset() {
	bw.buffer.Reset()
}
