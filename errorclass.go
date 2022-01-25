package errorclass

import "fmt"

type Err = error

type WrappedError struct{ Err }

func (e WrappedError) Unwrap() error {
	return e.Err
}

func (e WrappedError) Error() string {
	if e.Err == nil {
		return "(no inner error)"
	}
	return e.Err.Error()
}

type ErrUnexpected WrappedError
type ErrUsage WrappedError
type ErrAbuse WrappedError
type ErrInput WrappedError
type ErrEndUser WrappedError
type ErrClient WrappedError
type ErrExternal WrappedError
type ErrExists WrappedError
type ErrDoesntExist WrappedError
type ErrExhaustedRetries WrappedError
type ErrTODO WrappedError

func (e ErrUnexpected) Error() string  { return fmt.Sprintf("ErrUnexpected: %v", WrappedError(e)) }
func (e ErrUsage) Error() string       { return fmt.Sprintf("ErrUsage: %v", WrappedError(e)) }
func (e ErrAbuse) Error() string       { return fmt.Sprintf("ErrAbuse: %v", WrappedError(e)) }
func (e ErrInput) Error() string       { return fmt.Sprintf("ErrInput: %v", WrappedError(e)) }
func (e ErrEndUser) Error() string     { return fmt.Sprintf("ErrEndUser: %v", WrappedError(e)) }
func (e ErrClient) Error() string      { return fmt.Sprintf("ErrClient: %v", WrappedError(e)) }
func (e ErrExternal) Error() string    { return fmt.Sprintf("ErrExternal: %v", WrappedError(e)) }
func (e ErrExists) Error() string      { return fmt.Sprintf("ErrExists: %v", WrappedError(e)) }
func (e ErrDoesntExist) Error() string { return fmt.Sprintf("ErrDoesntExist: %v", WrappedError(e)) }
func (e ErrExhaustedRetries) Error() string {
	return fmt.Sprintf("ErrExhaustedRetries: %v", WrappedError(e))
}
func (e ErrTODO) Error() string { return fmt.Sprintf("ErrTODO: %v", WrappedError(e)) }

func (e ErrUnexpected) Unwrap() error       { return WrappedError(e).Unwrap() }
func (e ErrUsage) Unwrap() error            { return WrappedError(e).Unwrap() }
func (e ErrAbuse) Unwrap() error            { return WrappedError(e).Unwrap() }
func (e ErrInput) Unwrap() error            { return WrappedError(e).Unwrap() }
func (e ErrEndUser) Unwrap() error          { return WrappedError(e).Unwrap() }
func (e ErrClient) Unwrap() error           { return WrappedError(e).Unwrap() }
func (e ErrExternal) Unwrap() error         { return WrappedError(e).Unwrap() }
func (e ErrExists) Unwrap() error           { return WrappedError(e).Unwrap() }
func (e ErrDoesntExist) Unwrap() error      { return WrappedError(e).Unwrap() }
func (e ErrExhaustedRetries) Unwrap() error { return WrappedError(e).Unwrap() }
func (e ErrTODO) Unwrap() error             { return WrappedError(e).Unwrap() }

func (e ErrUnexpected) Is(t error) bool       { _, ok := t.(ErrUnexpected); return ok }
func (e ErrUsage) Is(t error) bool            { _, ok := t.(ErrUsage); return ok }
func (e ErrAbuse) Is(t error) bool            { _, ok := t.(ErrAbuse); return ok }
func (e ErrInput) Is(t error) bool            { _, ok := t.(ErrInput); return ok }
func (e ErrEndUser) Is(t error) bool          { _, ok := t.(ErrEndUser); return ok }
func (e ErrClient) Is(t error) bool           { _, ok := t.(ErrClient); return ok }
func (e ErrExternal) Is(t error) bool         { _, ok := t.(ErrExternal); return ok }
func (e ErrExists) Is(t error) bool           { _, ok := t.(ErrExists); return ok }
func (e ErrDoesntExist) Is(t error) bool      { _, ok := t.(ErrDoesntExist); return ok }
func (e ErrExhaustedRetries) Is(t error) bool { _, ok := t.(ErrExhaustedRetries); return ok }
func (e ErrTODO) Is(t error) bool             { _, ok := t.(ErrTODO); return ok }
