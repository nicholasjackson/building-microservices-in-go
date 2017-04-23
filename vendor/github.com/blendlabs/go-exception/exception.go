package exception

import (
	"encoding/json"
	"fmt"
	"io"
)

const (
	// DefaultExceptionMessage is a default exception message.
	DefaultExceptionMessage = "An Exception Occurred"
)

// New returns a new exception by `Sprint`ing the messageComponents.
func New(messageComponents ...interface{}) error {
	message := fmt.Sprint(messageComponents...)
	if len(message) == 0 {
		message = DefaultExceptionMessage
	}
	return &Exception{message: message, stack: callers()}
}

// Newf returns a new exception by `Sprintf`ing the format and the args.
func Newf(format string, args ...interface{}) error {
	message := fmt.Sprintf(format, args...)
	if len(message) == 0 {
		message = DefaultExceptionMessage
	}
	return &Exception{message: message, stack: callers()}
}

// Wrap wraps an exception, will return error-typed `nil` if the exception is nil.
func Wrap(err error) error {
	if err == nil {
		return nil
	}

	if typedEx, isException := err.(*Exception); isException {
		return typedEx
	}
	return WrapError(err)
}

// Exception is an error with a stack trace.
type Exception struct {
	message string
	inner   error
	*stack
}

// Format allows for conditional expansion in printf statements
// based on the token and flags used.
// %+v : message + stack
// %v, %s : message
// %t : stack
func (e *Exception) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			io.WriteString(s, e.message)
			e.stack.Format(s, verb)
			return
		} else if s.Flag('-') {
			e.stack.Format(s, verb)
			return
		}
		fallthrough
	case 's':
		io.WriteString(s, e.message)
	case 'q':
		fmt.Fprintf(s, "%q", e.message)
	}
}

// MarshalJSON is a custom json marshaler.
func (e *Exception) MarshalJSON() ([]byte, error) {
	values := map[string]interface{}{}
	values["Message"] = e.message
	if e.stack != nil {
		values["Stack"] = e.StackTrace()
	}
	if e.inner != nil {
		innerJSON, err := json.Marshal(e.inner)
		if err != nil {
			return nil, err
		}
		values["Inner"] = string(innerJSON)
	}

	return json.Marshal(values)
}

// Inner returns the nested exception.
func (e *Exception) Inner() error {
	return e.inner
}

// Error implements the `error` interface
func (e *Exception) Error() string { return e.message }

// Message returns just the message, it is effectively
// an alias to .Error()
func (e *Exception) Message() string { return e.message }

// StackString returns the stack trace as a string.
func (e *Exception) StackString() string {
	return fmt.Sprintf("%v", e.stack)
}

// Cause wraps an exception and allows user to add a custom message (cause) to the error.
// This is useful when we want to retain the original stack trace but also add customized
// message at the same time. Will return error-typed `nil` if the exception is nil.
func Cause(cause string, err error) error {
	if err == nil {
		return nil
	}

	return &Exception{
		message: cause,
		inner:   Wrap(err),
	}
}

// Nest nests an arbitrary number of exceptions.
func Nest(err ...error) error {
	var ex *Exception
	var last *Exception
	var didSet bool

	for _, e := range err {
		if e != nil {
			var wrappedEx *Exception
			if typedEx, isTyped := e.(*Exception); !isTyped {
				wrappedEx = &Exception{
					message: e.Error(),
					stack:   callers(),
				}
			} else {
				wrappedEx = typedEx
			}

			if wrappedEx != ex {
				if ex == nil {
					ex = wrappedEx
					last = wrappedEx
				} else {
					last.inner = wrappedEx
					last = wrappedEx
				}
				didSet = true
			}
		}
	}
	if didSet {
		return ex
	}
	return nil
}

// WrapError is a shortcut method for wrapping an error by calling .Error() on it.
func WrapError(err error) error {
	if err == nil {
		return nil
	}
	return &Exception{
		message: err.Error(),
		stack:   callers(),
	}
}

// Is is a helper function that returns if an error is an exception.
func Is(err error) bool {
	if _, typedOk := err.(*Exception); typedOk {
		return true
	}
	return false
}

// As is a helper method that returns an error as an exception.
func As(err error) *Exception {
	if typed, typedOk := err.(*Exception); typedOk {
		return typed
	}
	return nil
}
