# NativeCloud_HR System Architecture and Workflow

This project is designed as the backend architecture for the company’s employee attendance and access control system. It consists of multiple modules and layers, aiming to achieve efficiency, scalability, and high availability.

## Table of Contents

1. [Backend System Architecture Overview](#backend-system-architecture-overview)  
2. [API Layer](#api-layer)  
3. [Message Queue Layer (Kafka / NATS)](#message-queue-layer-kafka--nats)  
4. [Database Layer (PostgreSQL / Redis)](#database-layer-postgresql--redis)  
5. [System Architecture](#system-architecture)  
6. [Error Handling and Reporting](#error-handling-and-reporting)  

---

## 1. Backend System Architecture Overview

```
📦 employee-access-system
│── 📂 cmd                  # Entry point
│   ├── 📝 main.go          # Main program
│── 📂 config               # Configuration
│   ├── 📝 config.go        # Config loader logic
│   ├── 📝 config.yaml      # Config file (GCP, DB, Kafka connections)
│   ├── 📝 .env             # Environment variables
│── 📂 internal             # Core business logic
│   ├── 📂 api              # API layer
│   │   ├── 📂 handlers.go  # HTTP request handlers
│   │   │   ├── 📝 auth.go      # Authentication
│   │   │   ├── 📝 clock.go     # Clock-in/out
│   │   │   ├── 📝 notify.go    # Notifications
│   │   │   ├── 📝 report.go    # Reports
│   │   ├── 📂 routes.go  # HTTP routes
│   │   │   ├── 📝 auth_routes.go      # Auth routes
│   │   │   ├── 📝 clock_routes.go     # Clock routes
│   │   │   ├── 📝 notify_routes.go    # Notification routes
│   │   │   ├── 📝 report_routes.go    # Report routes
│   │   ├── 📝 middleware.go# Authentication middleware
│   │   ├── 📝 router.go    # API router setup (Echo/Gin)
│   ├── 📂 db               # Database initialization
│   │   ├── 📝 postgres.go  # PostgreSQL connection
│   ├── 📂 messageQueue     # Message queue operations
│   │   ├── 📝 kafka.go     # Kafka config & producer
│   │   ├── 📝 producer.go  # Message producer
│   │   ├── 📝 consumer.go  # Message consumer
│   ├── 📂 repository       # Data access layer
│   │   │   ├── 📝 user_repo.go      # Auth data access
│   │   │   ├── 📝 clock_repo.go     # Clock data access
│   │   │   ├── 📝 notify_repo.go    # Notification data access
│   │   │   ├── 📝 report_repo.go    # Report data access
│   ├── 📂 service          # Service layer (business logic)
│   │   │   ├── 📝 auth.go      # Authentication service
│   │   │   ├── 📝 clock.go     # Clock service
│   │   │   ├── 📝 notify.go    # Notification service
│   │   │   ├── 📝 report.go    # Report service
│   ├── 📂 utils            # Utilities
│   │   ├── 📝 jwt.go       # Token management
├── 📂 deployments          # Deployment configurations
│── 📂 scripts              # DevOps scripts (in progress)
│   ├── 📝 migrate.sh       # DB migrations
│   ├── 📝 start.sh         # Startup script
│── 📂 docs                 # Documentation
│   ├── 📝 API.md           # API documentation
│   ├── 📝 architecture.md  # System architecture notes
│   ├── 📝 GCP_deployment.md# Cloud deployment guide
├── Dockerfile       # Containerization settings
├── docker-compose.yml # Local testing environment
│── go.mod                  # Go module dependencies
│── go.sum                  # Go module lock file
│── README.md               # Project README
```

---

## 2. Running the Project (Using Docker)

To quickly spin up the `NativeCloud_HR` development environment, we recommend using Docker for local setup. Follow the steps below.

### 2.1.Install Docker and Docker Compose

#### macOS / Windows:
If you are using macOS or Windows, you **must install and run Docker Desktop** to manage Docker containers.

Steps:
1. Download and install Docker Desktop from the [official website](https://www.docker.com/products/docker-desktop/).  
2. After installation, **ensure Docker Desktop is running** and can start containers.  
3. Docker Compose is bundled with Docker Desktop, so no additional installation is needed.  

> 💡 **Note**: If Docker Desktop is not running, executing `docker-compose up --build` may result in an error: *Cannot connect to the Docker daemon*.

#### Linux:
On Linux, this project has been tested without Docker Desktop. You only need Docker Engine and Docker Compose.  

- [Docker installation guide](https://docs.docker.com/engine/install/)  
- [Docker Compose installation guide](https://docs.docker.com/compose/install/)  

---


### 2.2. Install Go

Make sure [Go](https://go.dev/doc/install) is installed and your environment variables are configured.

⚠️ **Tip**: Run `go mod tidy` in your **local terminal**, not in VS Code’s integrated terminal, to avoid dependency fetching errors.


```bash
go version   # 確認 Go 已正確安裝
```

### 2.3. Clone the Repository

```bash
git clone https://github.com/4040www/NativeCloud_HR.git
cd NativeCloud_HR
```

### 2.4. Configure `.env` Variables

Create a `.env` file directly inside the `config/` folder (not inside a subfolder).

Use `config/.env.example` as reference:

```bash
# config/.env
DB_HOST = 35.221.151.72
DB_USER = (your_username)
DB_PASSWORD = (your_password)
DB_NAME = (your_db_name)
DB_PORT = 5432

JWT_SECRET = (your_secret)
```

⚠️ Make sure the DB connection information is up to date.

### 2.5. Install Dependencies and Start Docker Containers

```bash
go mod tidy
docker-compose up --build
```

This will start containers based on `docker-compose.yml`, including:

* API server
* PostgreSQL database
* Kafka message queue (if configured)

### 2.6. Access the Application

* Local API endpoint:：`http://localhost:8080`
* Health check API：`http://localhost:8080/api/status`

You can use Postman or similar tools to test endpoints such as:

* Authentication
* Clock-in/out
* Notifications
* Reports

### 2.7. Database Migration

```bash
docker-compose exec app ./scripts/migrate.sh
```

This updates the database schema to the latest version.

### 2.8. Stop Containers

```bash
docker-compose down
```

### 2.9. View Logs

```bash
docker-compose logs -f
```

This displays real-time logs for debugging.

