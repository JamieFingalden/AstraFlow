# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

AstraFlow is a comprehensive financial management system for SMEs and individual users, supporting multi-tenant architecture, OCR invoice recognition, and reimbursement workflows. The system consists of three main components:

1. **Frontend**: Vue 3 application with Vite for build tooling
2. **Backend**: Go application using Gin framework with clean architecture
3. **AI Service**: Flask-based OCR service (planned integration)

## Development Commands

### Frontend (Vue.js)

```bash
# Install dependencies
cd astraflow
npm install

# Start development server
npm run dev

# Build for production
npm run build

# Preview production build
npm run preview
```

Key frontend technologies:
- Vue 3 with Composition API
- Vite for fast development and build
- Pinia for state management
- Vue Router for client-side routing
- Tailwind CSS for styling
- Axios for HTTP requests
- ECharts and Recharts for data visualization

### Backend (Go)

```bash
# Navigate to backend directory
cd AstraFlow-go

# Initialize database (required for first-time setup)
go run cmd/migrate/main.go

# Start the development server
go run cmd/server/main.go

# Run with live reload (if using air)
air

# Build the project
go build ./...

# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests in a specific package
go test ./internal/service/
```

## Architecture Overview

### Frontend Structure
```
astraflow/
├── src/
│   ├── views/              # Page components
│   │   ├── HomeView.vue        # Landing page
│   │   ├── InvoiceUpload.vue   # File upload & OCR
│   │   ├── BillManagement.vue  # Bill listing & management
│   │   └── dashboard/
│   │       └── Dashboard.vue   # Analytics & visualization
│   ├── router/             # Vue Router configuration
│   ├── services/           # API client (Axios)
│   ├── composables/        # Vue composables
│   └── stores/             # Pinia state management
├── package.json            # Dependencies & scripts
└── vite.config.js          # Vite build configuration
```

### Backend Structure
```
AstraFlow-go/
├── cmd/                    # Application entry points
│   ├── server/             # Main HTTP server
│   └── migrate/            # Database migration tool
├── internal/               # Private application code
│   ├── api/                # HTTP routing configuration
│   ├── handler/            # HTTP request handlers (controllers)
│   ├── service/            # Business logic layer
│   ├── repository/         # Data access layer
│   ├── model/              # Data models and entities
│   └── database/           # Database connection setup
├── pkg/                    # Reusable packages
│   ├── config/             # Configuration management
│   ├── jwt/                # JWT token utilities
│   ├── logger/             # Logging utilities
│   └── middleware/         # HTTP middleware (auth, CORS, etc.)
└── go.mod                  # Go module definition
```

## Multi-Tenant Architecture

The system supports both enterprise tenants and individual users:

### Data Isolation Strategy
- **Individual Users**: `tenant_id = NULL`
- **Tenant Users**: `tenant_id` references tenant table
- **Soft Links**: No foreign key constraints, managed in service layer
- **Indexing**: Optimized queries with tenant-scoped indexes

### User Roles
- **admin**: Tenant administrator (can create/manage users)
- **normal**: Regular tenant user
- **personal**: Individual user (no tenant)

## Key Architectural Patterns

### 1. Clean Architecture (Backend)
- Clear separation of concerns: handlers → services → repositories → models
- Services contain business logic and orchestrate repository calls
- Repositories handle all database operations
- Models define data structures and database schema

### 2. JWT Authentication Flow
- Access tokens: 24-hour expiration
- Refresh tokens: 7-day expiration
- Middleware protects endpoints
- Tokens contain user ID, username, and role claims

### 3. Repository Pattern
- `repository/*.go` handles all database operations
- Services depend on repository interfaces, not directly on database
- Enables easy testing and data source switching

### 4. Service Layer Pattern
- `service/*.go` contains business logic
- Orchestrates multiple repository calls
- Handles validation and business rules

## Database Schema

The system uses MySQL with the following key tables:

1. **tenant** - Enterprise organization data
2. **user** - User accounts with tenant associations
3. **invoice** - Invoice and bill information
4. **ocr_result** - AI processing results
5. **attachment** - File storage metadata
6. **reimbursement** - Reimbursement requests
7. **reimbursement_item** - Line items for reimbursements

Database management uses GORM's AutoMigrate feature:
- Go structs in `internal/model/` define the database schema
- Running migrations syncs these definitions to MySQL
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

Key endpoints include:
- Authentication: `/api/v1/auth/register`, `/api/v1/auth/login`, `/api/v1/auth/refresh`, `/api/v1/auth/me`
- Tenants: CRUD operations on `/api/v1/tenants`
- Invoices: CRUD operations on `/api/v1/invoices`
- OCR Results: Creation and retrieval on `/api/v1/ocr`
- Attachments: Upload and management on `/api/v1/attachments`
- Reimbursements: Full workflow on `/api/v1/reimbursements`

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

### Frontend Development
1. Create new views in `src/views/`
2. Add routes in `src/router/index.js`
3. Implement API services in `src/services/`
4. Use Pinia stores for state management when needed
5. Follow existing component patterns for consistency