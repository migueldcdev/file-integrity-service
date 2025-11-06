# File Integrity Monitor

A Go-based file monitoring and integrity verification project.  
Tracks file changes in specified directories, computes file hashes, stores them in SQLite, and exposes a REST API for querying and updating file integrity.


## Project Overview

This project consists of three main components:

1. **Monitor Service**  
   - Watches one directory for file changes.  
   - Logs file creation, modification, and deletion events.
    
2. **Deep File Walker & Hasher**  
   - Recursively walks the specified directories.  
   - Computes cryptographic hashes (SHA-256) for each file.  
   - Saves file path, hash, and timestamps into an SQLite database.

3. **HTTP REST API**  
   Provides endpoints to:
   - List all hashed files
   - Check integrity of a specific file
   - Update file hashes manually


## Folder Structure

```

/cmd/
service/   → contains main binary for running the service
/internal/
monitor/   → file monitoring
db/        → SQLite database access and queries
api/       → HTTP API handlers
hash/      → hashing logic
service/   → recursive file hashing
go.mod
go.sum
.gitignore
README.md
````


## Key Features

- **Directory Monitoring**: Detects changes in real-time using `fsnotify`.
- **Integrity Verification**: Computes file hashes and compares them with stored values.
- **Database Storage**: Stores file metadata and hashes in SQLite for persistence.
- **REST API**: Query file list, check integrity, update hashes via HTTP requests.


## Usage

### Run the Service

```bash
go build -o service ./cmd/service
./service
````

### REST API Endpoints

| Endpoint            | Method | Description                                                             |
| ------------------- | ------ | ----------------------------------------------------------------------- |
| `/hashedfiles`            | GET    | List all tracked files                                                  |
| `/check-file-integrity?path=<file>` | GET    | Check integrity of a file                                               |
| `/update-file-hash`             | POST   | Recompute and update the hash of a specific file (send path in JSON body)|

---

## Database Schema

```sql
CREATE TABLE hash (
   id INTEGER PRIMARY KEY AUTOINCREMENT,
   path TEXT UNIQUE NOT NULL,
   hash TEXT NOT NULL,
   size INTEGER NOT NULL,
   created_at DATETIME NOT NULL,
   last_checked DATETIME NOT NULL
)
```

* `id` hash id.
* `path` path to the hashed file.
* `hash` the latest computed value.
* `size` lenght in bytes.
* `created_at` hash creation time.
* `last_checked` last time the file was processed.


## Planned Enhancements

* Add a **CLI tool** to interact with the API (scan, status, update commands).
* Support **recursive directory watching**, including new subdirectories.
* Add **unit tests** for hashing and database functions.
* Add **logging and observability** for better debugging and monitoring.


## Technologies

* **Language**: Go
* **Database**: SQLite3
* **API**: net/http (RESTful)
* **Concurrency**: Goroutines for monitoring and hash computation

