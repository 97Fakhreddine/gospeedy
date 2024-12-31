# GoSpeedy Framework

Welcome to **GoSpeedy**, the next-generation web framework designed for developers who value simplicity, speed, and efficiency. With **zero configuration**, an intuitive structure, and best-in-class performance, GoSpeedy makes building web services easier and faster than ever.

## **Main Goals**

- **Zero Config to Start**: With GoSpeedy, you can immediately start building your web service app without worrying about complex configurations.
- **Insanely Fast**: Leverage Go's speed to handle high request volumes and manage databases efficiently.
- **Simple and Intuitive**: Designed for developers, the framework provides a clear structure with built-in features to speed up your development cycle.
- **RBAC Support**: Built-in Role-Based Access Control (RBAC) with decorators for seamless integration.
- **CLI Tools**: Generate controllers, services, modules, and more using our powerful CLI commands.

---

## **Directory Structure**

Here’s a high-level overview of the GoSpeedy framework structure:

```
gospeedy/
├── core/                        # Core functionalities of the framework
│   ├── app.go                   # App bootstrap and server logic
│   ├── controller.go            # Controller definitions
│   ├── service.go               # Service definitions
│   ├── decorators.go            # Route and Guard decorators
│   ├── middleware.go            # Middleware system
│   ├── generator.go             # CLI generators for modules, controllers, etc.
│   ├── di.go                    # Dependency injection system
│   ├── auth.go                  # Authentication and authorization system
│   ├── jwt.go                   # JWT token management
│   ├── oauth2.go                # OAuth2 integration
│   ├── validator.go             # Request validation
│   ├── events.go                # Event-driven architecture
│   ├── cache.go                 # In-memory caching
│   ├── scheduler.go             # Task scheduling
├── env/                         # Environment configuration folder
│   ├── config.go                # Configuration loader
│   ├── development.env          # Development environment variables
│   ├── production.env           # Production environment variables
├── modules/                     # Example app modules
│   ├── main.go                  # Main entry point for the app
│   ├── user/                    # Example user module
│   │   ├── user_controller.go   # User controller
│   │   ├── user_service.go      # User service
│   │   ├── user_routes.go       # User route constants
│   ├── file/                    # File management module
│   │   ├── file_controller.go   # File controller
│   │   ├── file_service.go      # File upload service
│   ├── order/                   # Order processing module
│   │   ├── order_controller.go  # Order controller
│   │   ├── order_service.go     # Order service
├── tests/                       # Testing folder
│   ├── unit/                    # Unit tests for individual components
│   │   ├── auth_test.go         # Test for authentication
│   │   ├── user_service_test.go # Test for User service
│   ├── integration/             # Integration tests
│   │   ├── database_test.go     # Test for database connections
│   │   ├── api_test.go          # Test for HTTP APIs
│   ├── e2e/                     # End-to-End tests
│   │   ├── app_test.go          # Full application test
├── templates/                   # CLI templates
│   ├── controller.tpl           # Controller template
│   ├── service.tpl              # Service template
│   ├── module.tpl               # Module template
├── cli/                         # CLI entry point and commands
│   ├── main.go                  # CLI main file
│   ├── generate.go              # CLI generators
│   ├── templates/
│   │   ├── controller_template.go
│   │   ├── service_template.go
│   │   ├── module_template.go
│   │   ├── resolver_template.go
│   │   ├── mutation_template.go
│   │   ├── controller_test_template.go
│   │   ├── service_test_template.go
│   │   ├── module_test_template.go
│   │   ├── resolver_test_template.go
│   │   ├── mutation_test_template.go
│   └── ...
├── utils/                       # Utility functions
│   ├── logger.go                # Logger utility
│   ├── helper.go                # Helper functions
│   ├── cache.go                 # Caching helpers
│   ├── validator.go             # Validation helpers
└── README.md                    # Documentation

```

---

## **Features**

### 1. **Core Functionalities**

- Controllers, Services, and Decorators for organizing your business logic.
- Middleware support for request handling and preprocessing.
- Dependency injection for managing services and components efficiently.

### 2. **Built-in RBAC Support**

- Use decorators to implement Role-Based Access Control (RBAC).
- Predefined Auth Guards make it easy to secure your routes.

### 3. **Environment Management**

- Structured environment folder for managing configurations seamlessly.
- Supports `development.env`, `production.env`, and more.

### 4. **CLI Tool**

- Generate controllers, services, and modules quickly with the CLI.
- Example:
  ```bash
  gospeedy generate module user
  ```

### 5. **Database Management**

- Optimized for working with various databases like PostgreSQL, MySQL, or MongoDB.

### 6. **Performance Focused**

- Designed to handle thousands of requests per second.
- Minimal resource usage thanks to Go’s efficiency.

---

## **Getting Started**

### Installation

1. **Clone the repository**:

   ```bash
   git clone https://github.com/yourusername/gospeedy.git
   ```

2. **Navigate to the directory**:

   ```bash
   cd gospeedy
   ```

3. **Run the CLI to generate your first app**:

   ```bash
   gospeedy new myapp
   ```

4. **Start your app**:

   ```bash
   go run main.go
   ```

---

## **Why GoSpeedy?**

- Faster request handling than existing frameworks.
- Clear and consistent structure for your apps.
- Predefined tools and templates to kickstart your project.
- Go’s simplicity ensures robust and maintainable code.

---

Stay tuned for further updates and documentation. Let’s build something great, together with GoSpeedy!

