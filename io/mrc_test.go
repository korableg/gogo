package io

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MultiReadCloser(t *testing.T) {
	r1 := io.NopCloser(bytes.NewBuffer([]byte("Hello ")))
	r2 := io.NopCloser(bytes.NewBuffer([]byte("world!")))

	m := MultiReadCloser(r1, r2)

	data, err := io.ReadAll(m)
	assert.NoError(t, err)
	assert.Equal(t, data, []byte("Hello world!"))

	assert.NoError(t, m.Close())
}
