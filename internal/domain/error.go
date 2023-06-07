package domain

import (
	"bytes"
	"fmt"
)

// TODO: fix declaring error's messages in different layers

const (
	// CodeInvalidArgument indicates client specified an invalid argument
	CodeInvalidArgument Code = "invalid_argument"

	// CodeNotFound means requested entity was not found
	CodeNotFound Code = "not_found"

	// CodeAlreadyExists means an attempt to create an entity failed because one already exists
	CodeAlreadyExists Code = "already_exists"

	// CodePermissionDenied indicates the caller does not have permission to execute the specified operation
	CodePermissionDenied Code = "permission_denied"

	// CodeInternal means an internal error occured
	CodeInternal Code = "internal"

	// CodeUnauthenticated indicates the request does not have valid authentication credentials for the operation
	CodeUnauthenticated Code = "unauthenticated"
)

type Code string

type Error struct {
	Code Code   // Code provides general information about the error
	Msg  string // Msg provides additional context in human-readable form
	Op   string // Op (operation) provides additional context about error's location
	Err  error  // Err is a nested error
}

func NewErrWithMsg(code Code, msg string) error {
	return &Error{
		Code: code,
		Msg:  msg,
	}
}

func NewInternalError(Op string, err error) error {
	return &Error{
		Code: CodeInternal,
		Op:   Op,
		Err:  err,
	}
}

// Error returns the string representation of the Error
func (e *Error) Error() string {
	var buf bytes.Buffer

	if e.Op != "" {
		fmt.Fprintf(&buf, "%s: ", e.Op)
	}

	if e.Err != nil {
		buf.WriteString(e.Err.Error())
	} else {
		if e.Code != "" {
			fmt.Fprintf(&buf, "<%s> ", e.Code)
		}
		buf.WriteString(e.Msg)
	}

	return buf.String()
}

// ErrorCode returns the code of the root error, if available. Otherwise returns Internal.
func ErrorCode(err error) Code {
	if err == nil {
		return ""
	} else if e, ok := err.(*Error); ok && e.Code != "" {
		return e.Code
	} else if ok && e.Err != nil {
		return ErrorCode(e.Err)
	}
	return CodeInternal
}

// ErrorMessage returns the message of the error, if available. Otherwise returns a error message.
func ErrorMessage(err error) string {
	if err == nil {
		return ""
	} else if e, ok := err.(*Error); ok && e.Msg != "" {
		return e.Msg
	} else if ok && e.Err != nil {
		return ErrorMessage(e.Err)
	}
	return "An internal error has occurred"
}
