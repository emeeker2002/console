package consoleerrors

import (
	"fmt"
)

type NotFoundError struct {
	ConsoleError
}

func (e NotFoundError) Error() string {
	return "requested resource not found"
}

type NotUniqueError struct {
	ConsoleError
}

func (e NotUniqueError) Error() string {
	return "unique constraint violation"
}

type DatabaseError struct {
	Console ConsoleError
}

func (e DatabaseError) Error() string {
	return "database error"
}

func (e DatabaseError) Wrap(call, function string, err error) error {
	_ = e.Console.Wrap(call, function, err)
	return e
}

type AMTError struct {
	ConsoleError
}

func (e AMTError) Error() string {
	return "amt error"
}

type ConsoleError struct {
	file          string
	Function      string
	Call          string
	Message       string
	InnerTrace    string
	OriginalError error
}

func (e ConsoleError) Error() string {
	return fmt.Sprintf("%s - %s - %s: %w", e.file, e.Function, e.Call, e.OriginalError)
}
func (e ConsoleError) FriendlyMessage() string {
	return e.Message
}
func (e *ConsoleError) Wrap(call, function string, err error) error {
	e.Call = call
	e.Function = function
	e.OriginalError = err

	if err != nil {
		e.InnerTrace = err.Error()
	}
	return e
}
func CreateConsoleError(file string) ConsoleError {
	var message = ""

	return ConsoleError{
		file:    file,
		Message: message,
	}
}
