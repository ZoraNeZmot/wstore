# WStore

A pet project for an online store: a fullstack application with a REST API in Go and a SPA in React. Currently implemented: user authentication, personal account, and catalog showcase (UI with mock data); the sales module and shopping cart are in development.

## Features

- **Authentication** — registration and login via email/password (Argon2), JWT access token + refresh token in HttpOnly cookie
- **Users** — user list, profile (`/user/me`), avatar
- **Frontend** — home page with catalog, login/registration, personal account, settings
- **i18n** — Russian and English interface
- **API documentation** — Swagger UI at `/swagger/`
- **Google OAuth** — code prepared, temporarily disabled (`//go:build ignore`)

## Tech Stack

| Layer | Technologies |
|-------|--------------|
| Backend | Go 1.25, `net/http`, PostgreSQL, JWT, Argon2, Logrus, Swaggo |
| Frontend | React 19, Vite 8, React Router 7, ESLint |
| Database | PostgreSQL (`accounts` schema, stored procedures) |
| Infrastructure | HTTPS (TLS / mkcert), env configuration |

## Project Structure

```
wstore/
├── backend/          # Go HTTP API
│   ├── cmd/wstore/   # entry point, Swagger docs
│   └── internal/     # handlers, storage, config, app
├── frontend/         # React SPA (Vite)
├── migrations/       # PostgreSQL SQL migrations
└── migrations.sql    # migration script
```

## Requirements

- Go 1.25+
- Node.js 20+
- PostgreSQL 14+

## Quick Start

### 1. Database

> The `migrations/roles.sql` file is needed to configure database roles and is not included in the repository — create it locally before running migrations.

### 2. Backend

Create a `.env` in the root or in `backend/` (loaded via `godotenv/autoload`):

```env
DATABASE_USER=postgres
DATABASE_PASSWORD=postgres
DATABASE_HOST=localhost
DATABASE_PORT=5432
DATABASE_DATABASE=wstore

SERVER_JWT_SECRET=your-secret-key
SERVER_TLS=true
SERVER_CERT_FILE=certificates/server.crt
SERVER_KEY_FILE=certificates/server.key

FRONTEND_URL=https://localhost:8000
LOG_LEVEL=info
```

```bash
cd backend
go run ./cmd/wstore
```

API: `https://localhost:8080/api/`  
Swagger: `https://localhost:8080/swagger/`

### 3. Frontend

```bash
cd frontend
npm install
npm run dev
```

Application: `https://localhost:8000`  
Requests to `/api` are proxied to the backend.

## API (main endpoints)

| Method | Path | Description |
|--------|------|-------------|
| GET | `/api/service/health` | Health check |
| POST | `/api/auth/email/signup` | Registration |
| POST | `/api/auth/email/login` | Login (returns access token) |
| GET | `/api/auth/refresh` | Refresh access token (cookie) |
| POST | `/api/auth/logout` | Logout |
| GET | `/api/users` | User list 🔒 |
| GET | `/api/user/{id}` | User profile (`me` — current user) 🔒 |
| GET | `/api/user/profile-image/{id}` | Avatar 🔒 |

🔒 — requires the `Authorization: Bearer <access_token>` header

## Backend Architecture

- HTTP server on the Go standard library without a framework
- Routes grouped by domain: `auth`, `users`, `service`
- PostgreSQL access via `database/sql` and stored functions
- Unified response format: `{ success, data, error }`
- JWT middleware for protected routes

## Frontend Architecture

- React SPA with React Router
- `AuthContext` — stores access token, automatic refresh
- `LanguageContext` — RU/EN switching
- API client with proxy through the Vite dev server

## Project Status

The project is in active development. The product catalog on the home page uses static mock data; integration with the sales API is planned.

## License

Apache 2.0