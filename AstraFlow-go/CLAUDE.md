# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

AstraFlow is a financial management system for SMEs and individual users, supporting multi-tenant architecture, OCR invoice recognition, and reimbursement workflows. The backend is built with Go using clean architecture principles.

## Development Commands

### Setup and Run
```bash
# Initialize database (required for first-time setup)
go run cmd/migrate/main.go

# Start the development server
go run cmd/server/main.go

# Build the project
go build ./...

# Run with live reload (if using air)
air
```

### Database Management
```bash
# Run database migrations (creates/updates tables)
go run cmd/migrate/main.go

# The migration tool uses GORM AutoMigrate to sync Go struct definitions to MySQL schema
```

### Testing
```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests in a specific package
go test ./internal/service/
```

## Architecture Overview

The project follows clean architecture with clear separation of concerns:

```
cmd/           - Application entry points
├── server/    - Main HTTP server
└── migrate/   - Database migration tool

internal/      - Private application code
├── api/       - HTTP routing configuration
├── handler/   - HTTP request handlers (controllers)
├── service/   - Business logic layer
├── repository/- Data access layer
├── model/     - Data models and entities
└── database/  - Database connection setup

pkg/           - Reusable packages
├── config/    - Configuration management
├── jwt/       - JWT token utilities
├── logger/    - Logging utilities
└── middleware/- HTTP middleware (auth, CORS, etc.)
```

## Key Architectural Patterns

### 1. Repository Pattern
- `repository/user_repository.go` handles all database operations
- Services depend on repository interfaces, not directly on database
- Enables easy testing and data source switching

### 2. Service Layer Pattern
- `service/auth_service.go` contains business logic
- Orchestrates multiple repository calls
- Handles validation and business rules

### 3. JWT Authentication Flow
- Access tokens: 24-hour expiration
- Refresh tokens: 7-day expiration
- Middleware `pkg/middleware/auth.go` protects endpoints
- Tokens contain user ID, username, and role claims

### 4. Multi-tenant Support
- Users can belong to a tenant (`tenant_id` field) or be individual users (`tenant_id = null`)
- All data models include `tenant_id` for data isolation

## Database Schema Management

The project uses GORM's AutoMigrate feature:
- Go structs in `internal/model/` define the database schema
- Running `go run cmd/migrate/main.go` syncs these definitions to MySQL
- Supports adding new fields without data loss
- Uses soft deletes (`gorm.DeletedAt`)

## API Structure

All APIs follow the pattern `/api/v1/{resource}` with consistent response format:

```json
{
  "code": 200,
  "message": "Success message",
  "data": { ... }
}
```

### Authentication Endpoints
- `POST /api/v1/auth/register` - User registration
- `POST /api/v1/auth/login` - User login
- `POST /api/v1/auth/refresh` - Token refresh (requires auth)
- `GET /api/v1/auth/me` - Get current user (requires auth)

### Request/Response Models
- Defined in handler files (e.g., `RegisterRequest`, `LoginRequest`)
- Use struct tags for validation (`binding:"required,min=3"`)
- Automatic JSON validation and error handling

## Configuration

Configuration is managed through `pkg/config/`:
- Uses YAML configuration files
- Database connection, server port, and JWT secret are configurable
- Environment-specific configs supported

## Security Implementation

- Passwords hashed with bcrypt
- JWT tokens signed with configurable secret
- Input validation through struct tags
- SQL injection protection via GORM
- Authentication middleware for protected routes

## Development Guidelines

### Adding New Features
1. Define models in `internal/model/`
2. Create repository in `internal/repository/`
3. Implement business logic in `internal/service/`
4. Add handlers in `internal/handler/`
5. Register routes in `internal/api/router.go`
6. Run migrations if schema changed

### Database Changes
1. Modify Go structs in `internal/model/`
2. Run `go run cmd/migrate/main.go`
3. Test the changes

### Adding New API Endpoints
1. Define request/response structs in handler
2. Implement handler method
3. Add service layer logic if needed
4. Register route in router.go
5. Apply authentication middleware if needed

### My rules
1. Don't repeat the start of the project
2. Frontend and backend projects I will start, when I need to reboot, prompt "Please restart frontend/backend"
3. Write the code strictly according to the original architecture, and only change the architecture when I request it.