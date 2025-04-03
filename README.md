# Go Template Repository

A production-ready Go project template with best practices, standard project layout, and pre-configured CI/CD. This template helps you quickly bootstrap new Go projects with a consistent structure and modern development practices.

## Features

- 🏗️ **Standard Project Layout**: Follows Go project layout best practices
- 🔧 **Pre-configured CI/CD**: GitHub Actions workflow for testing and linting
- 📝 **Code Quality**: Pre-configured golangci-lint with essential linters
- 🧪 **Testing Setup**: Ready for unit and integration tests
- 🔐 **Environment Management**: `.env` support with example configuration
- 🛠️ **Build Automation**: Makefile with common development commands
- 📚 **Documentation**: Ready for API docs, OpenAPI/Swagger specs
- 🔍 **Code Analysis**: Static analysis and linting tools

## Quick Start

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

## Development Workflow

### Prerequisites
- Go 1.23.4 or later
- Make
- Docker (optional, for containerized development)

### Getting Started

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

### Available Make Commands

```bash
make help        # Show available commands
make build       # Build the application
make test        # Run tests
make run         # Run the application
make clean       # Clean build artifacts
make lint        # Run linters
make fmt         # Format code
```

## Code Quality

This template includes several code quality tools:

- **golangci-lint**: Static code analysis with multiple linters
- **go test**: Built-in testing framework
- **go fmt**: Code formatting
- **go vet**: Code analysis

### Linting Configuration

The template uses golangci-lint with the following linters enabled:
- govet: Reports suspicious constructs
- errcheck: Checks for unchecked errors
- staticcheck: Advanced static analysis
- ineffassign: Detects ineffective assignments
- unused: Reports unused code
- misspell: Finds commonly misspelled words

## CI/CD Pipeline

The template includes a GitHub Actions workflow that:
1. Runs on push to main and pull requests
2. Sets up Go environment
3. Installs dependencies
4. Runs tests
5. Runs linters
6. Ensures code quality

## Best Practices

1. **Project Structure**
   - Use `internal/` for private application code
   - Use `pkg/` for code that can be used by external applications
   - Keep `cmd/` as thin as possible

2. **Code Organization**
   - Follow Go standard project layout
   - Use dependency injection
   - Write tests for your code
   - Use environment variables for configuration

3. **Development**
   - Write clear commit messages
   - Keep dependencies up to date
   - Document public APIs
   - Follow Go idioms and best practices

## Customization

1. **Adding New Features**
   - Add new directories in `internal/` for new features
   - Update `Makefile` with new commands
   - Add new environment variables to `.env.example`

2. **Modifying CI/CD**
   - Edit `.github/workflows/ci.yml` for CI changes
   - Add new GitHub Actions workflows as needed

3. **Adding Dependencies**
   - Use `go get` to add new dependencies
   - Run `go mod tidy` to clean up
   - Update documentation if needed

## Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## License

MIT License - feel free to use this template for any project.

## Support

If you find this template helpful, please consider:
- Starring the repository
- Reporting issues
- Contributing improvements
- Sharing with others 