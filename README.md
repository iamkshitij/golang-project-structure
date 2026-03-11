# 🚀 Go Fiber API with Oracle Database

A modern, high-performance REST API built with Go Fiber framework and Oracle Database integration.

## ✨ Features

- **⚡ Fast & Lightweight**: Built with Fiber v3 for maximum performance
- **🗄️ Oracle Database**: Robust enterprise database integration
- **🏗️ Clean Architecture**: Well-structured project with separation of concerns
- **🔧 Environment Configuration**: Easy configuration management with .env files
- **🛡️ Graceful Shutdown**: Proper signal handling for clean server shutdown
- **📦 Modular Design**: Resource-based API structure for scalability

## 🛠️ Tech Stack

- **Framework**: [Fiber v3](https://gofiber.io/) - Express-inspired web framework
- **Database**: Oracle Database with [go-ora](https://github.com/sijms/go-ora) driver
- **ORM**: [SQLx](https://github.com/jmoiron/sqlx) for database operations
- **Language**: Go 1.25.1

## 📁 Project Structure

```
├── api/
│   └── resources/
│       └── user/
│           ├── model.go      # User data models
│           ├── repository.go # Database operations
│           ├── resource.go   # Business logic
│           └── routes.go     # HTTP route handlers
├── config/
│   └── config.go            # Configuration management
├── database/
│   └── oracle.go            # Database connection setup
├── .env                     # Environment variables
├── main.go                  # Application entry point
└── go.mod                   # Go module dependencies
```

## 🚀 Quick Start

### Prerequisites

- Go 1.25.1 or higher
- Oracle Database (local or remote)
- Git

### Installation

1. **Clone the repository**
   ```bash
   git clone <your-repo-url>
   cd golang.project.structure
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Configure environment variables**
   
   Copy the example environment file and update with your settings:
   ```bash
   cp .env.example .env
   ```
   
   Update `.env` with your Oracle database credentials:
   ```env
   DB_HOST=localhost
   DB_PORT=1521
   DB_USER=your_username
   DB_PASSWORD=your_password
   DB_NAME=your_database_name
   
   API_PORT=8080
   BASE_URL=http://localhost:8080
   ```

4. **Run the application**
   ```bash
   go run main.go
   ```
   
   Or use the provided script:
   ```bash
   ./run.sh
   ```

The API will be available at `http://localhost:8080`

## 🔧 Configuration

The application uses environment variables for configuration. All settings are defined in the `.env` file:

| Variable | Description | Default |
|----------|-------------|---------|
| `DB_HOST` | Oracle database host | `localhost` |
| `DB_PORT` | Oracle database port | `1521` |
| `DB_USER` | Database username | - |
| `DB_PASSWORD` | Database password | - |
| `DB_NAME` | Database name/service | - |
| `API_PORT` | API server port | `8080` |
| `BASE_URL` | Base URL for the API | `http://localhost:8080` |

## 📊 Database Schema

### User Table
```sql
CREATE TABLE users (
    ID NUMBER PRIMARY KEY,
    NAME VARCHAR2(100) NOT NULL,
    EMAIL VARCHAR2(255) UNIQUE NOT NULL,
    PHONE VARCHAR2(20),
    STATUS VARCHAR2(20) DEFAULT 'active',
    CREATED_AT TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

## 🔌 API Endpoints

### User Management

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/api/users` | Get all users |
| `GET` | `/api/users/:id` | Get user by ID |
| `POST` | `/api/users` | Create new user |
| `PUT` | `/api/users/:id` | Update user |
| `DELETE` | `/api/users/:id` | Delete user |

### Example Request/Response

**Create User**
```bash
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "phone": "+1234567890"
  }'
```

**Response**
```json
{
  "id": 1,
  "name": "John Doe",
  "email": "john@example.com",
  "phone": "+1234567890",
  "status": "active",
  "created_at": "2024-01-15T10:30:00Z"
}
```

## 🏗️ Architecture

This project follows a clean architecture pattern:

- **Models**: Data structures and database schemas
- **Repository**: Database access layer with CRUD operations
- **Resource**: Business logic and data processing
- **Routes**: HTTP handlers and request/response management

## 🔒 Security Features

- Environment-based configuration
- Database connection pooling
- Graceful error handling
- Input validation (to be implemented)

## 🚀 Deployment

### Docker (Coming Soon)
```dockerfile
# Dockerfile will be added for containerized deployment
```

### Production Checklist
- [ ] Set up proper logging
- [ ] Add authentication middleware
- [ ] Implement rate limiting
- [ ] Add input validation
- [ ] Set up monitoring
- [ ] Configure HTTPS

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- [Fiber](https://gofiber.io/) - Amazing Go web framework
- [go-ora](https://github.com/sijms/go-ora) - Pure Go Oracle driver
- [SQLx](https://github.com/jmoiron/sqlx) - Extensions to Go's database/sql

## 📞 Support

If you have any questions or need help, please:
- Open an issue on GitHub
- Contact the maintainers

---

⭐ **Star this repository if you find it helpful!**