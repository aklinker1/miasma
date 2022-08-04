package io

import (
	"bufio"
	"io"
	"time"

	"github.com/aklinker1/miasma/internal"
	"github.com/aklinker1/miasma/internal/server/zero"
)

type LogStream struct {
	rd      io.ReadCloser
	scanner *bufio.Scanner
}

func NewReadCloserLogStream(rd io.ReadCloser) *LogStream {
	return &LogStream{
		rd:      rd,
		scanner: bufio.NewScanner(rd),
	}
}

// Close implements LogStream
func (s *LogStream) Close() {
	s.rd.Close()
}

// NextLog implements LogStream
func (s *LogStream) NextLog() (log internal.Log, done bool, err error) {
	hasNext := s.scanner.Scan()
	if !hasNext {
		return zero.Log, true, s.scanner.Err()
	}
	return internal.Log{
		Message:   s.scanner.Text(),
		Timestamp: time.Now(),
	}, false, nil
}
