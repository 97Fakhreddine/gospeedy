package core

import (
	"fmt"
	"os"
)

// GenerateModule creates a new module with a controller and service
func GenerateModule(moduleName string) {
	// Create module folder
	err := os.Mkdir(moduleName, os.ModePerm)
	if err != nil {
		fmt.Printf("Error creating module directory: %v\n", err)
		return
	}

	// Create controller and service files
	createFile(moduleName, fmt.Sprintf("%s_controller.go", moduleName), "Controller content for "+moduleName)
	createFile(moduleName, fmt.Sprintf("%s_service.go", moduleName), "Service content for "+moduleName)

	fmt.Printf("Module '%s' generated successfully.\n", moduleName)
}

func createFile(moduleName, fileName, content string) {
	filePath := fmt.Sprintf("%s/%s", moduleName, fileName)
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()

	file.WriteString(content)
}
