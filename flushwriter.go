package goproxy

import (
	"io"
	"net/http"
)

// FlushWriter is a wrapper for writers with auto flush capabilities
type FlushWriter struct {
	w io.Writer
}

// NewFlushWriter creates a flush writer, wrapping a given writer
func NewFlushWriter(w io.Writer) *FlushWriter {
	return &FlushWriter{w: w}
}

// Write overrides the default write behaviour and automatically flushes on writing
func (w *FlushWriter) Write(b []byte) (int, error) {
	bytes, err := w.w.Write(b)
	if f, ok := w.w.(http.Flusher); ok {
		f.Flush()
	}
	return bytes, err
}
