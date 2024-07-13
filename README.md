
# Distributed Key-Value Store in Go

This project demonstrates a simple distributed key-value store implemented in Go. It uses TCP networking to allow multiple nodes to communicate and manage a shared key-value store.

## Features

- Basic key-value operations (`GET`, `SET`, `DELETE`)
- TCP networking with Go's `net` package
- Concurrency control with goroutines and channels
- Simple replication (to be implemented)

## Getting Started

### Prerequisites

- Go (version 1.15 or later recommended)

### Running the Server

1. Clone the repository or download the source code.
2. Navigate to the source directory.
3. Build the project:
   ```bash
   go build
   ```
4. Run the server:
   ```bash
   ./main
   ```

The server will start listening on `localhost:3000`. You can modify the port by changing the `address` parameter in the `main` function.

### Implementing Further

- **Client Implementation**: Implement a client to interact with the server using the key-value operations.
- **Replication**: Add a basic replication strategy to synchronize data across nodes for increased data availability.
- **Fault Tolerance**: Enhance the system to handle node failures gracefully.

## Architecture

The system is designed with simplicity in mind, primarily to demonstrate basic distributed systems concepts in Go. It uses a single-threaded approach per connection but can be expanded to a more robust architecture using Go's concurrency features.

## License

This project is open-sourced under the MIT License. See the LICENSE file for more details.
