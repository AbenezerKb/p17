package httpclient

import "io"

// this is to ensure that the client requests body is retirable
type nopReadCloser struct{ io.Reader }

func (nopReadCloser) Close() error { return nil }
