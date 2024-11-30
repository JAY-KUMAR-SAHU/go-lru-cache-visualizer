# go-lru-cache-visualizer

A simple React-based frontend to visualize and interact with an LRU (Least Recently Used) Cache implemented in Go.

---

## Features

- View the current state of the cache.
- Dynamically add words to the cache.
- Adjust the cache size with validation.
- Warning dialog when decreasing cache size.

---

## Prerequisites

Make sure you have the following installed:

1. [Node.js](https://nodejs.org/) (v16 or above recommended)
2. [npm](https://www.npmjs.com/) (comes with Node.js)
3. [Go](https://golang.org/) (to run the backend)

---

## Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/JAY-KUMAR-SAHU/go-lru-cache-visualizer.git
cd go-lru-cache-visualizer
```

### 2. Backend Setup

```bash
cd backend
go mod tidy
go run main.go
```

Run the Go backend on localhost:8080. Ensure it provides the following APIs:

1. GET /cache: Fetch the current cache state.
2. POST /add/{word}: Add a word to the cache.
3. POST /set-cache-size/{size}: Update the cache size.

### 3. Frontend Setup

```bash
cd frontend
npm install
npm start
```

This will launch the app at http://localhost:3000.

## File Structure

```bash
/go-lru-cache-visualizer
├── backend/        # Go backend
│ ├── main.go       # Backend implementation
│ ├── go.mod        # Go module configuration (defines dependencies)
│ ├── go.sum        # Go module checksum file (ensures dependency integrity)
├── frontend/       # React frontend
│ ├── public/       # Public assets
│ ├── src/
│ │ ├── components/
│ │ │ ├── Cache.js  # Cache component
│ │ ├── App.js      # Main application
│ │ ├── App.css     # General styles
│ │ └── index.js    # React entry point
│ ├── package.json  # npm configuration
├── README.md       # Setup guide (this file)
```
