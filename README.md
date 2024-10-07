# Template API Project

This project serves as a template for creating new API projects. It provides a structured foundation that can be easily adapted for different use cases while following best practices in software architecture. The goal is to facilitate the rapid development of robust and scalable APIs.

## Architecture

The project follows the principles of Hexagonal Architecture, which promotes separation of concerns and independence of frameworks. The main components of the architecture are organized as follows:
```
template-go-hexagonal/
├── cmd/                               # Entry point of the application
│    └── myapi/
│       └── main.go                 
├── configs/                           # Configuration files
├── deploy/                            # Deployment scripts and configurations
├── docs/                              # Documentation files 
├── internal/                          # Internal packages
│   ├── adapters/
│   │   ├── entrypoints/               # Entry points for HTTP API
│   │   │   └── v1/                    # DTOs and Mappers
│   │   │       ├── presenters/        # DTOs and Mappers
│   │   │       │   └── ...dto.go      # Example DTO definition
│   │   │       │   └── ...mapper.go   # Example Mapper
│   │   │       └── ...controller.go   # HTTP Controller for API endpoints
│   │   ├── mongodb/                   # MongoDB specific implementations
│   │   │   ├── presenters/            # DTOs and Mappers
│   │   │   │   ├── ...dto.go          # MongoDB DTO
│   │   │   │   └── ...mapper.go       # MongoDB Mapper
│   │   │   └── ...repository.go       # MongoDB repository implementation
│   │   └── ... 
│   ├── dependency-injection/          # Dependency injection containers
│   │   └── ...container.go
│   ├── domain/                        # Domain layer
│   │ ├── entities/                    # Business entities
│   │ │     └── entity.go              # Example entity definition
│   │ ├── gateways/                    # Interfaces for external services, such as APIs, databases, and other services.
│   │ │     └── ...gateway.go          # Example gateway interface
│   │ └── usecases/                    # Application logic
│   │       └── entity_usecase.go      # Example use case
├── scripts                            # Scripts for development or test environment configurations.
├── tests                              # Integration tests only.
├── go.mod                             # Go module definition 
├── go.sum                             # Go module dependencies 
└── README.md                          # Project documentation
```

### Components

- **CMD**: Entry point of the application.
- **Configs**: Configuration file templates or default settings.
- **Adapters Layer**: Contains the implementations that connect the application's business logic to external interfaces, such as APIs, databases, and other services. It is divided into subfolders to organize the different types of adapters.
- **Deploy**: IaaS, PaaS, orchestration system and container deployment configurations and templates (kubernetes/helm, mesos, terraform, bosh).
- **Domain Layer**: Contains the business logic and rules. It defines entities, gateways, and use cases.
- **Entrypoints**: Manages incoming requests and responses. This layer includes controllers and presenters (DTOs and mappers).
- **Tests**: Integration tests only, unit tests are located with the code files.

## Swagger
[Swagger](https://swagger.io/) is an API documentation tool that allows you to describe, consume, and visualize RESTful web services interactively. It provides a graphical interface that makes it easier to understand and use the API, allowing developers to test endpoints directly from the browser.

In this template, the Gin framework was used, and the Swagger implementation was done using [gin-swagger](https://github.com/swaggo/gin-swagger).

To start Swagger and generate the documentation, you can run the following command, where -g indicates the file with the 'swagger general API Info' and -d indicates the folder from which to search for functions with Swagger annotations.
```bash
swag init -g ./swagger/swagger.go -d ./internal/adapters/entrypoints
```

## Getting Started

### Prerequisites

- Go 1.23 or higher
- Docker and Docker Compose

### Installation

1. Clone the repository.

2. Install all dependencies.
```bash
go mod tidy
```
3. Run data base using doker-compose.
```bash
docker compose up -d
```
4. And run the application:
```bash
go run cmd/myapi/main.go
```

### Usage
You can access the API at http://localhost:8080.

## Testes

### Unit

This project uses the `stretchr/testify` testing library and the `Mockery` mock generation tool to ensure code quality and reliability.

#### stretchr/testify

[stretchr/testify](https://github.com/stretchr/testify) is a popular library for Go that provides a set of assertions and utilities that make writing tests easier.

#### Mockery
Mockery is a tool that automatically generates mocks from Go interfaces. This allows you to create consistent and up-to-date mocks without manually implementing each method. The mock generation is done via command, for example:

```bash
mockery --name=EntityGateway --dir=internal/domain/gateways --output=internal/domain/gateways/mocks --outpkg=mocks
```

#### Running the tests
To run the tests from all folders starting from the root, you can use the command:


```bash
go test ./internal/...
```

For a specific folder:

```bash
go test ./internal/domain/usecases
```

### Integration
Integration tests are performed through HTTP calls, the flow should follow the real code end-to-end of the application, only external dependencies should be mocked.
The naming of the test folders follows the route and not the implementation folders.
For database dependencies, a test database will be created, which runs in a Docker container during test execution. The configuration to create the test database is in the file `tests/v1/entities/integration_test.go`.

To run the integration tests, you can use the following command:

```bash
go test -v ./tests/v1/entities
```

## Development Tools

### Official Go Plugin for VSCode
Install the official Go plugin from the Team at Google (go.dev).

### Lint

We are using [golangci-lint](https://github.com/golangci/golangci-lint). To install it, you can use the following command:

```bash
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

To run linting on the project, you can use the following command:

```bash
golangci-lint run
```

#### Integrating with VSCode
If you already have the official Go plugin installed in VSCode, you can configure VSCode to run golangci-lint automatically when saving files. 
Add configurations to the VSCode settings file (settings.json), either user settings or workspace settings (project).
Here it was done in the workspace `template/.vscode/settings.json` so that the preferences are followed by all developers of the project.

### Git Hooks
Git Hooks are used to add checks and commands to be executed during pre-commit and pre-push hooks.
To ensure that all developers use the same Git hooks and maintain code quality, follow these steps after cloning the repository:

1. **Clone the Repository**:
   ```bash
   git clone <repository URL>
   cd <repository name>
   ```
2. **Run the Setup Script**:
The project includes a script that automatically sets up Git hooks. To run it, use the following command:
```bash
chmod +x ./scripts/hooks/setup-hooks.sh # Makes the script executable (only the first time)
./scripts/hooks/setup-hooks.sh          # Runs the script
```

### Debug
```bash
go install -v github.com/go-delve/delve/cmd/dlv@latest
```

## Contributing
Feel free to submit issues, fork the repository, and make pull requests. Contributions are welcome!

## License