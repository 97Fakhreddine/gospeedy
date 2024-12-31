package env

import (
    "log"
    "os"
)

// LoadEnv loads the environment configuration
func LoadEnv() {
    if err := os.Setenv("APP_ENV", "development"); err != nil {
        log.Fatalf("Error setting environment variable: %v", err)
    }
}
