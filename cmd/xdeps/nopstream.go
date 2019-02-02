package main

import "io"

type nopStream int

func (nopStream) Write(p []byte) (int, error) {
	return len(p), nil
}

func (nopStream) Read(p []byte) (int, error) {
	return 0, io.EOF
}

func (nopStream) Close() error { return nil }
