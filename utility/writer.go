package utility

import (
	"io"
	"sync"
)

// SyncWriter wraps a [io.Writer] and synchronizes all access to it.
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
