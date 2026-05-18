package validation

import (
	"regexp"
	"strings"
	"time"
)

// Validator collects validation errors.
type Validator struct {
	errors []string
}

// New creates a new Validator.
func New() *Validator {
	return &Validator{}
}

// Required checks that a string is not empty.
func (v *Validator) Required(field, value string) *Validator {
	if strings.TrimSpace(value) == "" {
		v.errors = append(v.errors, field+" is required")
	}
	return v
}

// MinLength checks minimum string length.
func (v *Validator) MinLength(field, value string, min int) *Validator {
	if len(value) < min {
		v.errors = append(v.errors, field+" must be at least "+string(rune(min))+" characters")
	}
	return v
}

// MaxLength checks maximum string length.
func (v *Validator) MaxLength(field, value string, max int) *Validator {
	if len(value) > max {
		v.errors = append(v.errors, field+" must be at most "+string(rune(max))+" characters")
	}
	return v
}

// InFuture checks that a time is in the future.
func (v *Validator) InFuture(field string, t time.Time) *Validator {
	if !t.After(time.Now()) {
		v.errors = append(v.errors, field+" must be in the future")
	}
	return v
}

// Matches checks that a string matches a regex pattern.
func (v *Validator) Matches(field, value, pattern string) *Validator {
	matched, _ := regexp.MatchString(pattern, value)
	if !matched {
		v.errors = append(v.errors, field+" has invalid format")
	}
	return v
}

// IsValid returns true if no errors.
func (v *Validator) IsValid() bool {
	return len(v.errors) == 0
}

// Errors returns all validation errors.
func (v *Validator) Errors() []string {
	return v.errors
}

// Error returns a combined error string, or nil if valid.
func (v *Validator) Error() error {
	if v.IsValid() {
		return nil
	}
	return &ValidationError{messages: v.errors}
}

// ValidationError is a collection of validation messages.
type ValidationError struct {
	messages []string
}

func (e *ValidationError) Error() string {
	return strings.Join(e.messages, "; ")
}

func (e *ValidationError) Messages() []string {
	return e.messages
}
