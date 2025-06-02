# Go REST API

A high-performance RESTful API service built with Go, PostgreSQL, and Redis. This project demonstrates clean architecture, proper error handling, and efficient database operations.

## Features

- **RESTful API endpoints** for managing items
- **PostgreSQL** for persistent data storage
- **Redis** for caching and session management
- **Gorilla Mux** for routing
- **Environment-based configuration**
- **Health check endpoint**
- **Proper error handling and logging**

## Prerequisites

- Go 1.16 or higher
- PostgreSQL 13 or higher
- Redis 6 or higher
- Git

## Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/yourusername/go-rest-api.git
cd go-rest-api
```

### 2. Set Up the Database

1. Start PostgreSQL and Redis services
2. Create a new database and user (or use the default credentials):
   ```sql
   CREATE USER sampleuser WITH PASSWORD 'samplepass';
   CREATE DATABASE sampledb OWNER sampleuser;
   GRANT ALL PRIVILEGES ON DATABASE sampledb TO sampleuser;
   ```

3. Run the database initialization script:
   ```bash
   sudo -u postgres ./setup_db.sh
   ```

### 3. Configure Environment Variables

Create a `.env` file in the root directory:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=sampleuser
DB_PASSWORD=samplepass
DB_NAME=sampledb
REDIS_ADDR=localhost:6379
REDIS_PASSWORD=
SERVER_PORT=8080
```

### 4. Install Dependencies

```bash
go mod download
```

### 5. Run the Application

```bash
go run main.go
```

The server will start on `http://localhost:8080`

## API Endpoints

### Health Check

```http
GET /health
```

**Response:**
```json
{
  "status": "ok"
}
```

### Items

#### Get All Items

```http
GET /api/items
```

**Response:**
```json
[
  {
    "id": 1,
    "name": "Item 1"
  },
  {
    "id": 2,
    "name": "Item 2"
  }
]
```

#### Create New Item

```http
POST /api/items
Content-Type: application/json

{
  "name": "New Item"
}
```

**Response (201 Created):**
```json
{
  "id": 3,
  "name": "New Item"
}
```

## Project Structure

```
.
├── cache/               # Redis cache implementation
│   └── redis.go
├── db/                  # Database related code
│   ├── postgres.go
│   └── init.sql
├── handlers/            # HTTP request handlers
│   └── item_handler.go
├── models/              # Data models and database operations
│   └── item.go
├── .env.example        # Example environment variables
├── go.mod             # Go module file
├── go.sum             # Go module checksums
├── main.go            # Application entry point
└── README.md          # This file
```

## Environment Variables

| Variable        | Default Value    | Description                          |
|----------------|----------------|--------------------------------------|
| DB_HOST        | localhost       | PostgreSQL host                      |
| DB_PORT        | 5432           | PostgreSQL port                      |
| DB_USER        | sampleuser     | PostgreSQL user                      |
| DB_PASSWORD    | samplepass     | PostgreSQL password                  |
| DB_NAME        | sampledb       | PostgreSQL database name             |
| REDIS_ADDR    | localhost:6379 | Redis server address                |
| REDIS_PASSWORD|                | Redis password (if any)              |
| SERVER_PORT   | 8080           | Port the server will listen on       |

## Running Tests

```bash
go test -v ./...
```

## Contributing

1. Fork the repository
2. Create a new branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [Gorilla Mux](https://github.com/gorilla/mux) - Request router and dispatcher
- [Go PostgreSQL Driver](https://github.com/lib/pq) - PostgreSQL driver for Go
- [Go Redis](https://github.com/go-redis/redis) - Redis client for Go
