package core

import "fmt"

// DIContainer holds registered services and components
type DIContainer struct {
	services map[string]interface{}
}

// NewDIContainer initializes a new Dependency Injection container
func NewDIContainer() *DIContainer {
	return &DIContainer{
		services: make(map[string]interface{}),
	}
}

// Register registers a service in the container
func (c *DIContainer) Register(name string, service interface{}) {
	c.services[name] = service
}

// Get retrieves a service by its name
func (c *DIContainer) Get(name string) (interface{}, error) {
	service, exists := c.services[name]
	if !exists {
		return nil, fmt.Errorf("service not found: %s", name)
	}
	return service, nil
}
