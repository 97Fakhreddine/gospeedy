package core

import "fmt"

// Validator defines validation logic for requests
type Validator struct {}

// NewValidator creates a new validator instance
func NewValidator() *Validator {
	return &Validator{}
}

// Validate checks if the given data is valid
func (v *Validator) Validate(data interface{}) bool {
	// Add validation logic here
	return true
}

// ValidateEmail validates email format
func (v *Validator) ValidateEmail(email string) bool {
	// Add email validation logic here
	return true
}
