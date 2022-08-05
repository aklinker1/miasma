package io

import (
	"bufio"
	"io"
	"strings"
	"time"

	"github.com/aklinker1/miasma/internal"
	"github.com/aklinker1/miasma/internal/server/zero"
)

type LogStream struct {
	rd      io.ReadCloser
	scanner *bufio.Scanner
}

func NewLogStream(rd io.ReadCloser) *LogStream {
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
	// The first 8 characters are some crazy unicode text (\x0000\x0001 or something) that break
	// GQLGen's subscription or Golang's channel implementations. So we remove them!
	message := s.scanner.Text()[8:]
	sections := strings.SplitN(message, " ", 2)
	timestamp, err := time.Parse("2006-01-02T15:04:05.999999999Z", sections[0])
	if err != nil {
		return zero.Log, false, err
	}
	return internal.Log{
		Timestamp: timestamp,
		Message:   sections[1],
	}, false, nil
}
