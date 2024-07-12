package errors

import "errors"

var (
	ErrSomethingWentWrong = errors.New("something went wrong")
	ErrResourceNotFound   = errors.New("resource not found")
	ErrTokenExpired       = errors.New("token has expired")
	ErrTokenMissing       = errors.New("bearer token missing")
	ErrTokenInvalid       = errors.New("invalid token")
	ErrUnauthorized       = errors.New("not authorized to perform this action")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrInvalidID          = errors.New("invalid id")
	ErrInvalidRequest     = errors.New("invalid request body")
)

// New returns an error that formats as the given text.
// Each call to New returns a distinct error value even if the text is identical.
func New(text string) error {
	return errors.New(text)
}

// As finds the first error in err's tree that matches target, and if one is found, sets
// target to that error value and returns true. Otherwise, it returns false.
//
// The tree consists of err itself, followed by the errors obtained by repeatedly
// calling its Unwrap() error or Unwrap() []error method. When err wraps multiple
// errors, As examines err followed by a depth-first traversal of its children.
//
// An error matches target if the error's concrete value is assignable to the value
// pointed to by target, or if the error has a method As(interface{}) bool such that
// As(target) returns true. In the latter case, the As method is responsible for
// setting target.
//
// An error type might provide an As method so it can be treated as if it were a
// different error type.
//
// As panics if target is not a non-nil pointer to either a type that implements
// error, or to any interface type.
func As(err error, target any) bool {
	return errors.As(err, target)
}

// Is reports whether any error in err's tree matches target.
//
// The tree consists of err itself, followed by the errors obtained by repeatedly
// calling its Unwrap() error or Unwrap() []error method. When err wraps multiple
// errors, Is examines err followed by a depth-first traversal of its children.
//
// An error is considered to match a target if it is equal to that target or if
// it implements a method Is(error) bool such that Is(target) returns true.
//
// An error type might provide an Is method so it can be treated as equivalent
// to an existing error. For example, if MyError defines
//
//	func (m MyError) Is(target error) bool { return target == fs.ErrExist }
//
// then Is(MyError{}, fs.ErrExist) returns true. See [syscall.Errno.Is] for
// an example in the standard library. An Is method should only shallowly
// compare err and the target and not call [Unwrap] on either.
func Is(err error, target error) bool {
	return errors.Is(err, target)
}
