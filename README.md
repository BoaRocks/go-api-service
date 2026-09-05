# Go API Service

A lightweight HTTP API written in Go demonstrating backend service development, structured logging, health monitoring, unit testing, and technical troubleshooting.

## Overview

This project implements a small HTTP service using Go's standard library.

The service includes a health-check endpoint, JSON responses, structured request logging, HTTP method validation, and automated unit tests.

## Technologies

- Go
- HTTP
- JSON
- REST-style API concepts
- Structured logging
- Unit testing

## Features

- `/health` health-check endpoint
- JSON API responses
- HTTP status-code handling
- Method validation
- Structured JSON request logging
- Unit tests
- No third-party dependencies

## Repository Files

`main.go`  
Contains the HTTP service, routing, JSON responses, and structured logging.

`main_test.go`  
Contains automated tests for the health endpoint.

`go.mod`  
Defines the Go module.

`.gitignore`  
Excludes generated binaries and local files.

## Run the Service

```bash
go run .
```

The service listens on:

```text
http://localhost:8080
```

## Test the Health Endpoint

Open:

```text
http://localhost:8080/health
```

Or use curl:

```bash
curl http://localhost:8080/health
```

Expected response:

```json
{
  "status": "ok",
  "message": "service is running"
}
```

## Run Unit Tests

```bash
go test ./...
```

## Example Structured Log

Requests generate structured JSON log output containing information such as:

```text
method
path
remote_addr
duration_ms
```

Structured logs make it easier to investigate application behavior and troubleshoot service issues.

## Troubleshooting Approach

Useful troubleshooting steps include:

1. Confirming the service is listening on port 8080
2. Checking HTTP response codes
3. Reviewing structured request logs
4. Testing the `/health` endpoint
5. Running automated tests
6. Checking whether another application is already using the port
7. Reviewing request methods and paths

## Skills Demonstrated

- Go programming
- HTTP fundamentals
- JSON
- Structured logging
- Unit testing
- Debugging
- Service monitoring concepts
- Technical troubleshooting
- Technical documentation

## Purpose

This repository is part of my technical portfolio focused on systems, networking, software fundamentals, cybersecurity, and troubleshooting.