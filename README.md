
# Thoth

**Thoth** is a project built with a Go backend and a modern Svelte frontend, integrated with Meilisearch to experiment with search functionality for databases.

## 🚀 **Getting Started Locally**

### Prerequisites

Before you begin, ensure you have the following tools installed:

- [Go](https://golang.org/dl/) (v1.24.1)
- [Node.js](https://nodejs.org/) (v20.9)
- [Yarn](https://classic.yarnpkg.com/en/docs/install/)
- [Meilisearch](https://www.meilisearch.com/) (for production search)

### 1. **Clone the Repository**

```bash
git clone https://github.com/yourusername/thoth.git
cd thoth
```

### 2. **Backend (Go)**

#### Install Go Dependencies

Navigate to the `/backend` folder and install dependencies:

```bash
cd backend
go mod tidy
```

#### Run the Go Backend

To start the backend server locally:

```bash
go run main.go
```

The Go server will run on `http://localhost:8080`.

### 3. **Frontend (Svelte)**

#### Install Frontend Dependencies with Yarn

Navigate to the `/frontend` folder and install the necessary dependencies using Yarn:

```bash
cd frontend
yarn install
```

#### Run the Frontend

To start the frontend development server:

```bash
yarn dev
```

The Svelte app will be available at `http://localhost:5173`.

## 🛠️ **Tech Stack**

- **Go** for backend API development.
- **Svelte** for building the frontend UI.
- **Meilisearch** for fast and efficient search functionality.
