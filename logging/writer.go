package logging

import (
	"io"
	"sync"
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

type SyncWriter struct {
	mutex  sync.Mutex
	writer io.Writer
}

func (w *SyncWriter) Write(p []byte) (n int, err error) {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	return w.writer.Write(p)
}

func NewSyncWriter(writer io.Writer) *SyncWriter {
	return &SyncWriter{
		mutex:  sync.Mutex{},
		writer: writer,
	}
}
