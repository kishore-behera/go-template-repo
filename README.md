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
- 🗄️ **MongoDB Integration**: Pre-configured MongoDB with GORM
- 📊 **Structured Logging**: Zap logger for production-grade logging

## Using This Template

### Method 1: GitHub Actions (Recommended)

1. Go to the "Actions" tab in this repository
2. Select "Create Repository from Template" workflow
3. Click "Run workflow"
4. Fill in the required information:
   - Repository name
   - Repository description
   - Repository visibility (public/private)
   - GitHub username or organization
5. Click "Run workflow"

The workflow will:
- Create a new repository from this template
- Update the module name in go.mod
- Update repository name in README
- Set up branch protection rules
- Create standard issue labels
- Create a welcome issue

### Method 2: Manual Creation

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
│   ├── database/         # Database connection and models
│   ├── handler/          # HTTP handlers
│   ├── logger/           # Logging configuration
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
- Go 1.21 or later
- Make
- Docker (optional, for containerized development)
- MongoDB (if running locally)

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

### Docker Development

```bash
# Start all services (app, MongoDB, Redis)
docker-compose up

# Run in detached mode
docker-compose up -d

# Stop all services
docker-compose down

# View logs
docker-compose logs -f
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

## Environment Variables

Key environment variables:

```env
# Application
GIN_MODE=debug
PORT=8080

# MongoDB
MONGODB_URI=mongodb://localhost:27017
MONGODB_DB=appdb
MONGODB_USER=admin
MONGODB_PASSWORD=password

# Redis
REDIS_HOST=localhost
REDIS_PORT=6379
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
3. Sets up MongoDB service
4. Installs dependencies
5. Runs tests
6. Runs linters
7. Ensures code quality

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

## Kubernetes Deployment

### Prerequisites
- Kubernetes cluster
- kubectl configured
- Docker registry access

### Building and Pushing Docker Image

```bash
# Build the image
docker build -t your-registry/go-template:latest .

# Push to registry
docker push your-registry/go-template:latest
```

### Deploying to Kubernetes

1. Create secrets:
```bash
# Create secrets from template
kubectl create secret generic app-secrets \
  --from-literal=mongodb-uri="mongodb://admin:password@mongodb:27017" \
  --from-literal=mongodb-user="admin" \
  --from-literal=mongodb-password="password"
```

2. Apply Kubernetes manifests:
```bash
# Apply all manifests
kubectl apply -f k8s/

# Or apply individually
kubectl apply -f k8s/configmap.yaml
kubectl apply -f k8s/deployment.yaml
kubectl apply -f k8s/service.yaml
```

3. Verify deployment:
```bash
# Check pods
kubectl get pods

# Check services
kubectl get svc

# Check logs
kubectl logs -f deployment/app
```

### Kubernetes Resources

The template includes the following Kubernetes manifests:
- `k8s/deployment.yaml`: Application deployment with 3 replicas
- `k8s/service.yaml`: LoadBalancer service
- `k8s/configmap.yaml`: Non-sensitive configuration
- `k8s/secrets.yaml`: Sensitive data template

Features:
- Health checks (readiness and liveness probes)
- Resource limits and requests
- ConfigMap and Secret management
- Load balancing
- High availability (3 replicas)

### Production Considerations

1. **Security**
   - Use a secrets management solution (e.g., HashiCorp Vault)
   - Enable network policies
   - Use RBAC for service accounts
   - Enable pod security policies

2. **Monitoring**
   - Add Prometheus metrics
   - Configure Grafana dashboards
   - Set up alerting

3. **Scaling**
   - Configure HPA (Horizontal Pod Autoscaling)
   - Set appropriate resource limits
   - Use node affinity/anti-affinity

4. **CI/CD**
   - Add Kubernetes deployment to CI pipeline
   - Use Helm for package management
   - Implement blue-green deployments 