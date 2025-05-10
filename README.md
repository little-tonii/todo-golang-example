# Todo Golang Example

A simple Todo REST API built with Golang for practice purposes.

## 🚀 Technologies Used

<p align="center" style="margin-bottom: 20px;">
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" alt="Go" width="48" height="48" style="margin-right: 16px;"/>
  <img src="https://raw.githubusercontent.com/gin-gonic/logo/master/color.png" alt="Gin" width="48" height="48" style="margin-right: 16px;"/>
  <img src="https://avatars.githubusercontent.com/u/17219288?s=200&v=4" alt="GORM" width="48" height="48" style="margin-right: 16px;"/>
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/postgresql/postgresql-original.svg" alt="PostgreSQL" width="48" height="48" style="margin-right: 16px;"/>
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/docker/docker-original.svg" alt="Docker" width="48" height="48" style="margin-right: 16px;"/>
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/swagger/swagger-original.svg" alt="Swagger" width="48" height="48" style="margin-right: 16px;"/>
  <img src="https://img.icons8.com/ios-filled/50/000000/jwt.png" alt="JWT" width="48" height="48" style="margin-right: 16px;"/>
</p>

## Features

- User registration and login with JWT authentication
- CRUD operations for Todo items
- Pagination for Todo list
- Swagger API documentation
- Dockerized for easy deployment

## Project Architecture

This project follows the [Go Project Layout Standard](https://github.com/golang-standards/project-layout) for organizing the codebase, which is widely adopted in the Go community for scalable and maintainable projects.

Additionally, the application is designed using Domain-Driven Design (DDD) principles. The codebase is structured to separate domain logic, application services, infrastructure, and interfaces, making it easier to maintain, test, and extend.

## Getting Started

### Prerequisites

- [Go](https://golang.org/) >= 1.18
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

### Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/yourusername/todo-golang-example.git
   cd todo-golang-example
   ```

2. **Copy environment variables:**

   ```bash
   cp .env.example .env
   ```

   Edit `.env` if needed.

3. **Start the application with Docker Compose:**

   ```bash
   docker-compose -f deployments/docker-compose.yml up --build
   ```

   The API will be available at `http://localhost:8080`.

4. **Access Swagger UI:**

   Visit [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html) to explore the API documentation.

## Project Structure

```
.
├── cmd/                # Main application entrypoint
├── internal/           # Application code (domain, service, handler, etc.) following DDD principles
├── pkg/                # Shared packages (middleware, utils, etc.)
├── deployments/        # Docker Compose and deployment files
├── build/              # Dockerfile and build scripts
├── docs/               # Swagger docs
├── .env.example        # Example environment variables
└── README.md
```

- The `internal/` directory is organized according to DDD, separating domain models, repositories, services, and handlers.
- The overall structure is based on the [golang-standards/project-layout](https://github.com/golang-standards/project-layout) repository.

## License

This project is for practice and educational purposes.
