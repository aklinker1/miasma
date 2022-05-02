package server

import (
	"bytes"
	"fmt"
)

const (
	ECONFLICT = "conflict"  // action cannot be performed
	EINTERNAL = "internal"  // internal error
	EINVALID  = "invalid"   // validation failed
	ENOTFOUND = "not_found" // entity does not exist
)

// Error defines a standard application error.
type Error struct {
	// Machine-readable error code.
	Code string

	// Human-readable message that gets logged.
	InternalMessage string

	// Human-readable message that is shown to the user.
	ExternalMessage string

	Op  string
	Err error
}

// ErrorCode returns the code of the root error, if available. Otherwise returns EINTERNAL.
func ErrorCode(err error) string {
	if err == nil {
		return ""
	} else if e, ok := err.(*Error); ok && e.Code != "" {
		return e.Code
	} else if ok && e.Err != nil {
		return ErrorCode(e.Err)
	}
	return EINTERNAL
}

func ExternalErrorMessage(err error) string {
	if err == nil {
		return ""
	} else if e, ok := err.(*Error); ok && e.ExternalMessage != "" {
		return e.ExternalMessage
	} else if ok && e.Err != nil {
		return ExternalErrorMessage(e.Err)
	}
	return "Unknown error"
}

// Error returns the string representation of the error message.
func (e *Error) Error() string {
	var buf bytes.Buffer

	// Print the current operation in our stack, if any.
	if e.Op != "" {
		fmt.Fprintf(&buf, "%s: ", e.Op)
	}

	// If wrapping an error, print its Error() message.
	// Otherwise print the error code & message.
	if e.Err != nil {
		buf.WriteString(e.Err.Error())
	} else {
		if e.Code != "" {
			fmt.Fprintf(&buf, "<%s> ", e.Code)
		}
		if e.InternalMessage != "" {
			buf.WriteString(e.InternalMessage)
		} else {
			buf.WriteString(e.ExternalMessage)
		}
	}
	return buf.String()
}
