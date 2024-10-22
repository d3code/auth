# Auth Service
[![Build Status](https://img.shields.io/github/actions/workflow/status/d3code/auth/build.yaml?branch=master&label=Build&logo=github)](https://github.com/d3code/auth/actions/workflows/build.yaml)
## Overview

This repository contains the implementation of an authentication service using Go. The service provides functionalities for user registration and JWT token generation.

## Project Structure

- `proto/v1/auth/auth.proto`: Protocol Buffers definitions for the authentication service.
- `internal/service`: Implementation of the authentication services.
- `internal/model`: Data models used in the service.
- `pkg/encrypt`: Encryption utilities.
- `config`: Configuration.
- `go.mod`: Go module dependencies.

## Getting Started

### Prerequisites

- Go 1.22+
- MySQL 5.7+
- `buf` for proto generation

### Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/d3code/auth.git
    cd auth
    ```

2. Install Go dependencies:
    ```sh
    go mod tidy
    ```
3. Install `buf`:
    ```sh
    brew tap bufbuild/buf
    brew install buf
    ```
   
4. Compile Protocol Buffers:
    ```sh
    make proto
    ```

5. Setup database:
    ```sh
    ./database.sh
    ```

### Running the Service

To run the authentication service, use the following command:
```sh
go run main.go