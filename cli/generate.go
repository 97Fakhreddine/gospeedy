package cli

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"io/ioutil"
	"strings"
)

// Function to load template files
func loadTemplate(filePath string) (string, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("Error reading template file: %v", err)
	}
	return string(data), nil
}

// Generate generates the specified type (module, controller, service, resolver, etc.)
func Generate(generateType string, name string, withTests bool) {
	switch generateType {
	case "module":
		GenerateModule(name, withTests)
	case "controller":
		GenerateController(name, withTests)
	case "service":
		GenerateService(name, withTests)
	case "resolver":
		GenerateResolver(name, withTests)
	case "mutation":
		GenerateMutation(name, withTests)
	default:
		fmt.Println("Unknown generation type:", generateType)
	}
}

// GenerateModule generates the module files (controller, service) and optionally test files
func GenerateModule(moduleName string, withTests bool) {
	modulePath := filepath.Join("modules", moduleName)
	if err := os.MkdirAll(modulePath, os.ModePerm); err != nil {
		fmt.Println("Error creating module directory:", err)
		return
	}

	// Load templates
	controllerTemplate, err := loadTemplate("cli/templates/controller_template.go")
	if err != nil {
		fmt.Println("Error loading controller template:", err)
		return
	}
	serviceTemplate, err := loadTemplate("cli/templates/service_template.go")
	if err != nil {
		fmt.Println("Error loading service template:", err)
		return
	}
	moduleTestTemplate, err := loadTemplate("cli/templates/module_test_template.go")
	if err != nil {
		fmt.Println("Error loading module test template:", err)
		return
	}

	// Create the controller file
	controllerFilePath := filepath.Join(modulePath, moduleName+"_controller.go")
	createFileFromTemplate(controllerFilePath, "ControllerName", moduleName, controllerTemplate)

	// Create the service file
	serviceFilePath := filepath.Join(modulePath, moduleName+"_service.go")
	createFileFromTemplate(serviceFilePath, "ServiceName", moduleName, serviceTemplate)

	// If --with-tests flag is passed, generate the test file
	if withTests {
		testPath := filepath.Join("tests", "unit", moduleName+"_test.go")
		createFileFromTemplate(testPath, "ModuleName", moduleName, moduleTestTemplate)
	}

	fmt.Printf("Module '%s' generated successfully\n", moduleName)
}

// GenerateController generates a controller file and optionally a test file
func GenerateController(controllerName string, withTests bool) {
	// Load templates
	controllerTemplate, err := loadTemplate("cli/templates/controller_template.go")
	if err != nil {
		fmt.Println("Error loading controller template:", err)
		return
	}
	controllerTestTemplate, err := loadTemplate("cli/templates/controller_test_template.go")
	if err != nil {
		fmt.Println("Error loading controller test template:", err)
		return
	}

	controllerPath := filepath.Join("modules", controllerName, controllerName+"_controller.go")
	createFileFromTemplate(controllerPath, "ControllerName", controllerName, controllerTemplate)

	// If --with-tests flag is passed, generate the test file
	if withTests {
		testPath := filepath.Join("tests", "unit", controllerName+"_controller_test.go")
		createFileFromTemplate(testPath, "ControllerName", controllerName, controllerTestTemplate)
	}

	fmt.Printf("Controller '%s' generated successfully\n", controllerName)
}

// GenerateService generates a service file and optionally a test file
func GenerateService(serviceName string, withTests bool) {
	// Load templates
	serviceTemplate, err := loadTemplate("cli/templates/service_template.go")
	if err != nil {
		fmt.Println("Error loading service template:", err)
		return
	}
	serviceTestTemplate, err := loadTemplate("cli/templates/service_test_template.go")
	if err != nil {
		fmt.Println("Error loading service test template:", err)
		return
	}

	servicePath := filepath.Join("modules", serviceName, serviceName+"_service.go")
	createFileFromTemplate(servicePath, "ServiceName", serviceName, serviceTemplate)

	// If --with-tests flag is passed, generate the test file
	if withTests {
		testPath := filepath.Join("tests", "unit", serviceName+"_service_test.go")
		createFileFromTemplate(testPath, "ServiceName", serviceName, serviceTestTemplate)
	}

	fmt.Printf("Service '%s' generated successfully\n", serviceName)
}

// GenerateResolver generates a GraphQL resolver file and optionally a test file
func GenerateResolver(resolverName string, withTests bool) {
	// Load templates
	graphqlResolverTemplate, err := loadTemplate("cli/templates/resolver_template.go")
	if err != nil {
		fmt.Println("Error loading resolver template:", err)
		return
	}
	graphqlResolverTestTemplate, err := loadTemplate("cli/templates/resolver_test_template.go")
	if err != nil {
		fmt.Println("Error loading resolver test template:", err)
		return
	}

	resolverPath := filepath.Join("modules", resolverName, resolverName+"_resolver.go")
	createFileFromTemplate(resolverPath, "ModuleName", resolverName, graphqlResolverTemplate)

	// If --with-tests flag is passed, generate the test file
	if withTests {
		testPath := filepath.Join("tests", "unit", resolverName+"_resolver_test.go")
		createFileFromTemplate(testPath, "ModuleName", resolverName, graphqlResolverTestTemplate)
	}

	fmt.Printf("Resolver '%s' generated successfully\n", resolverName)
}

// GenerateMutation generates a GraphQL mutation file and optionally a test file
func GenerateMutation(mutationName string, withTests bool) {
	// Load templates
	graphqlMutationTemplate, err := loadTemplate("cli/templates/mutation_template.go")
	if err != nil {
		fmt.Println("Error loading mutation template:", err)
		return
	}
	graphqlMutationTestTemplate, err := loadTemplate("cli/templates/mutation_test_template.go")
	if err != nil {
		fmt.Println("Error loading mutation test template:", err)
		return
	}

	mutationPath := filepath.Join("modules", mutationName, mutationName+"_mutation.go")
	createFileFromTemplate(mutationPath, "ModuleName", mutationName, graphqlMutationTemplate)

	// If --with-tests flag is passed, generate the test file
	if withTests {
		testPath := filepath.Join("tests", "unit", mutationName+"_mutation_test.go")
		createFileFromTemplate(testPath, "ModuleName", mutationName, graphqlMutationTestTemplate)
	}

	fmt.Printf("Mutation '%s' generated successfully\n", mutationName)
}

// createFileFromTemplate creates a file from the template with the provided values
func createFileFromTemplate(filePath string, placeholder string, value string, tmpl string) {
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	tmplParsed, err := template.New("file").Parse(tmpl)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}

	data := map[string]string{
		placeholder: value,
	}

	err = tmplParsed.Execute(file, data)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return
	}

	fmt.Printf("Created file: %s\n", filePath)
}
