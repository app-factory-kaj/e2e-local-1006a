// Package greeter computes greetings; it holds no state and persists nothing.
package greeter

import "fmt"

const (
	defaultName = "World"
	maxNameLen  = 100
)

// ErrNameTooLong is returned when the supplied name exceeds maxNameLen.
type ErrNameTooLong struct {
	Length int
}

func (e *ErrNameTooLong) Error() string {
	return fmt.Sprintf("name must be at most %d characters, got %d", maxNameLen, e.Length)
}

// Greet returns a greeting message for name, defaulting to "World" when name
// is empty, and rejects a name longer than maxNameLen.
func Greet(name string) (string, error) {
	if len(name) > maxNameLen {
		return "", &ErrNameTooLong{Length: len(name)}
	}
	if name == "" {
		name = defaultName
	}
	return fmt.Sprintf("Hello, %s!", name), nil
}
