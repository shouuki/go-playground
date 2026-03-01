package logging

import (
	"io"
	"sync/atomic"
)

type atomicWriter struct {
	writer atomic.Value
}

func (w *atomicWriter) Set(writer io.Writer) {
	w.writer.Store(writer)
}

func (w *atomicWriter) Write(p []byte) (n int, err error) {
	return w.writer.Load().(io.Writer).Write(p)
}

type writerHolder struct {
	writer io.Writer
}

func (w *writerHolder) Write(p []byte) (n int, err error) {
	return w.writer.Write(p)
}

func newWriterHolder(writer io.Writer) *writerHolder {
	return &writerHolder{
		writer: writer,
	}
}
