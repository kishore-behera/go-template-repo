# Go Template Repository

A template repository for Go projects with best practices and standard project layout.

## Using this Template

1. Click the "Use this template" button on GitHub
2. Create a new repository with your desired name
3. Clone your new repository
4. Update the module name in `go.mod`:
   ```bash
   go mod edit -module github.com/yourusername/your-repo-name
   ```
5. Copy `.env.example` to `.env` and update values
6. Start coding!

## Project Structure

```
.
├── cmd/                    # Main applications
│   └── app/               # Application entry point
├── internal/              # Private application and library code
│   ├── config/           # Configuration management
│   ├── handler/          # HTTP handlers
│   ├── middleware/       # HTTP middleware
│   ├── model/           # Data models
│   └── service/         # Business logic
├── pkg/                  # Library code that's ok to use by external applications
├── api/                  # API definitions, OpenAPI/Swagger specs
├── configs/              # Configuration files
├── docs/                 # Documentation
├── scripts/              # Build, install, analysis, etc. scripts
├── test/                 # Additional external test applications and test data
├── .env                  # Environment variables
├── .gitignore           # Git ignore file
├── go.mod               # Go module definition
├── go.sum               # Go module checksums
├── Makefile             # Build automation
└── README.md            # Project documentation
```

## Getting Started

1. Clone this repository
2. Update the module name in `go.mod`
3. Start coding!

## Development

```bash
# Install dependencies
go mod tidy

# Run tests
go test ./...

# Build the application
go build -o bin/app cmd/app/main.go

# Run the application
./bin/app
```

## Available Make Commands

```bash
make help        # Show available commands
make build       # Build the application
make test        # Run tests
make run         # Run the application
make clean       # Clean build artifacts
```

## Best Practices

1. Use `internal/` for private application code
2. Use `pkg/` for code that can be used by external applications
3. Keep `cmd/` as thin as possible
4. Use dependency injection
5. Follow Go standard project layout
6. Write tests for your code
7. Use environment variables for configuration

## License

MIT License 