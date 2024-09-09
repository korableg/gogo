package io

import (
	"errors"
	"io"
)

type multiReadCloser struct {
	mr io.Reader
	c  []io.Closer
}

func MultiReadCloser(r ...io.ReadCloser) io.ReadCloser {
	readers := make([]io.Reader, len(r))
	closers := make([]io.Closer, len(r))

	for i, v := range r {
		readers[i] = v
		closers[i] = v
	}

	return &multiReadCloser{
		mr: io.MultiReader(readers...),
		c:  closers,
	}
}

func (mrc *multiReadCloser) Read(p []byte) (n int, err error) {
	return mrc.mr.Read(p)
}

func (mrc *multiReadCloser) Close() (err error) {
	for _, c := range mrc.c {
		err = errors.Join(err, c.Close())
	}
	return err
}

var _ io.ReadCloser = (*multiReadCloser)(nil)
