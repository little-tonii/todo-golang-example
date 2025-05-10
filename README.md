# Todo Golang Example

A simple Todo REST API built with Golang for practice purposes.

## 🚀 Technologies Used

<p>
<img src="https://img.shields.io/badge/Go-00ADD8?logo=go&logoColor=white&style=for-the-badge" alt="Go" />
<img src="https://img.shields.io/badge/Gin-00B386?logo=gin&logoColor=white&style=for-the-badge" alt="Gin" />
<img src="https://img.shields.io/badge/GORM-FF7043?logo=go&logoColor=white&style=for-the-badge" alt="GORM" />
<img src="https://img.shields.io/badge/PostgreSQL-4169E1?logo=postgresql&logoColor=white&style=for-the-badge" alt="PostgreSQL" />
<img src="https://img.shields.io/badge/Docker-2496ED?logo=docker&logoColor=white&style=for-the-badge" alt="Docker" />
<img src="https://img.shields.io/badge/Swagger-85EA2D?logo=swagger&logoColor=white&style=for-the-badge" alt="Swagger" />
<img src="https://img.shields.io/badge/JWT-000000?logo=jsonwebtokens&logoColor=white&style=for-the-badge" alt="JWT" />
</p>

## Features

- User registration and login with JWT authentication
- CRUD operations for Todo items
- Pagination for Todo list
- Swagger API documentation
- Dockerized for easy deployment

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
├── internal/           # Application code (domain, service, handler, etc.)
├── pkg/                # Shared packages (middleware, utils, etc.)
├── deployments/        # Docker Compose and deployment files
├── build/              # Dockerfile and build scripts
├── docs/               # Swagger docs
├── .env.example        # Example environment variables
└── README.md
```

## License

This project is for practice and educational purposes.
