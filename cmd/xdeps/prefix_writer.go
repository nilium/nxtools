package main

import (
	"bytes"
	"io"
)

type prefixWriter struct {
	prefix []byte
	w      io.Writer
	first  bool
}

func newPrefixWriter(w io.Writer, prefix string) *prefixWriter {
	return &prefixWriter{
		prefix: []byte("\n" + prefix),
		w:      w,
	}
}

func (w *prefixWriter) Write(p []byte) (n int, err error) {
	if !w.first {
		w.first = true
		if _, err := w.w.Write(w.prefix); err != nil {
			return 0, err
		}
	}
	n = len(p)
	p = bytes.Replace(p, []byte("\n"), w.prefix, -1)
	mn, err := w.w.Write(p)
	if mn < n {
		return mn, err
	}
	return n, err
}
