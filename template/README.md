# Template API Project

This project serves as a template for creating new API projects. It provides a structured foundation that can be easily adapted for different use cases while following best practices in software architecture. The goal is to facilitate the rapid development of robust and scalable APIs.

## Architecture

The project follows the principles of Hexagonal Architecture, which promotes separation of concerns and independence of frameworks. The main components of the architecture are organized as follows:
```
template/
├── cmd/                                # Entry point of the application
│    └── myapi/
│       └── main.go                 
├── configs/                            # Configuration files
├── deploy/                             # Deployment scripts and configurations
├── docs/                               # Documentation files 
├── internal/                           # Internal packages
│   ├── adapters/
│   │   ├── entrypoints/                # Entry points for HTTP API
│   │   │   └── v1/                     # DTOs and Mappers
│   │   │       ├── presenters/         # DTOs and Mappers
│   │   │       │   └── ...dto.go       # Example DTO definition
│   │   │       │   └── ...mapper.go    # Example Mapper
│   │   │       └── ...controller.go    # HTTP Controller for API endpoints
│   │   ├── mongodb/                    # MongoDB specific implementations
│   │   │   ├── presenters/             # DTOs and Mappers
│   │   │   │   ├── ...dto.go           # MongoDB DTO
│   │   │   │   └── ...mapper.go        # MongoDB Mapper
│   │   │   └── ...repository.go        # MongoDB repository implementation
│   │   └── ... 
│   ├── dependency-injection/           # Dependency injection containers
│   │   └── ...container.go
│   ├── domain/                         # Domain layer
│   │ ├── entities/                     # Business entities
│   │ │     └── entity.go               # Example entity definition
│   │ ├── gateways/                     # Interfaces for external services, such as APIs, databases, and other services.
│   │ │     └── ...gateway.go           # Example gateway interface
│   │ └── usecases/                     # Application logic
│   │       └── entity_usecase.go       # Example use case
├── go.mod                              # Go module definition 
├── go.sum                              # Go module dependencies 
└── README.md                           # Project documentation
```

### Components

- **CMD**: Entry point of the application.
- **Configs**: Configuration file templates or default settings.
- **Adapters Layer**: Contains the implementations that connect the application's business logic to external interfaces, such as APIs, databases, and other services. It is divided into subfolders to organize the different types of adapters.
- **Deploy**: IaaS, PaaS, orchestration system and container deployment configurations and templates (kubernetes/helm, mesos, terraform, bosh).
- **Domain Layer**: Contains the business logic and rules. It defines entities, gateways, and use cases.
- **Entrypoints**: Manages incoming requests and responses. This layer includes controllers and presenters (DTOs and mappers).


#### Swagger
[Swagger](https://swagger.io/) is an API documentation tool that allows you to describe, consume, and visualize RESTful web services interactively. It provides a graphical interface that makes it easier to understand and use the API, allowing developers to test endpoints directly from the browser.

Neste template foi usado o framework Gin e a implementação do swagger foi utilizando a [gin-swagger](https://github.com/swaggo/gin-swagger).

Para iniciar o swagger e gerar a documentação pode rodar o seguinte comando, onde -g indica o arquivo com onde está 'swagger general API Info' e o -d indica a partir de qual pasta procurar por funções que tenham anotações swagger.
```bash
swag init -g ./swagger/swagger.go -d ./internal/adapters/entrypoints
```

## Getting Started

### Prerequisites

- Go 1.16 or higher
- Docker and Docker Compose

### Installation

1. Clone the repository.
2. Run data base using doker-compose.
3. And run the application:

```bash
docker compose up -d
go run cmd/myapi/main.go
```

### Usage
You can access the API at http://localhost:8080.

### Contributing
Feel free to submit issues, fork the repository, and make pull requests. Contributions are welcome!

### License


