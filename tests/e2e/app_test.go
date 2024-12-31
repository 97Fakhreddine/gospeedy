// app_test.go
package e2e_test

import (
    "testing"
    "github.com/playwright-community/playwright-go"
    "github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
    pw, err := playwright.Run()
    if err != nil {
        t.Fatal(err)
    }
    defer pw.Stop()

    browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{Headless: playwright.Bool(true)})
    if err != nil {
        t.Fatal(err)
    }
    defer browser.Close()

    page, err := browser.NewPage()
    if err != nil {
        t.Fatal(err)
    }

    // Navigate to the login page
    err = page.Goto("http://localhost:8080/login")
    if err != nil {
        t.Fatal(err)
    }

    // Simulate entering login details and submit
    err = page.Fill("input[name='email']", "johndoe@example.com")
    err = page.Fill("input[name='password']", "password123")
    err = page.Click("button[type='submit']")
    
    // Assert successful login
    assert.Nil(t, err)
    assert.Contains(t, page.Title(), "Dashboard")
}
