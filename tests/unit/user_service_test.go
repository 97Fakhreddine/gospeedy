// user_service_test.go
package service_test

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "gospeedy/modules/user"
)

func TestCreateUser(t *testing.T) {
    service := user.NewUserService()

    // Simulate the input data
    newUser := user.User{
        Name:  "John Doe",
        Email: "johndoe@example.com",
    }

    // Call the method to be tested
    createdUser, err := service.CreateUser(newUser)

    // Assertions
    assert.Nil(t, err)
    assert.NotNil(t, createdUser)
    assert.Equal(t, newUser.Name, createdUser.Name)
    assert.Equal(t, newUser.Email, createdUser.Email)
}
