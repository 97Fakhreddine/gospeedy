package core

import "fmt"

// Service provides business logic for handling requests
type Service interface {
	Execute() string
}

// BaseService is a basic implementation of a service
type BaseService struct {}

// Execute performs the business logic for the service
func (s *BaseService) Execute() string {
	return fmt.Sprintf("Executing %T service", s)
}
