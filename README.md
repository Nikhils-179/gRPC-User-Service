# grpc User Service

This project demonstrates a client-server  gRPC-based User Service architecture service based on both a Go server and a Python client. The server provides various user management functionalities such as creating, listing , updating, searching by ids and id and deleting users. The client interacts with these services using gRPC.


## Table of Content

- Prerequisites
- Folder Structure
- Build and Run the Application
- Running Tests
- Accessing the gRPC Service Endpoints
- Configuration Details
- Screeshots


## prerequisites
Before running this application, ensure you have the following installed:

- Docker
- Go (1.22 or higher)
- Python (3.9 or higher)
- grpcurl (for testing the gRPC service)


## Folder Structure

Folder Structure :

```
└── grpc-user-service
    ├── Dockerfile
    ├── Readme.md
    ├── client
    │   ├── client.py
    │   ├── python_pb_files
    │   │   ├── user_pb2.py
    │   │   └── user_pb2_grpc.py
    │   └── requirements.txt
    ├── proto
    │   └── user.proto
    └── server
        ├── go.mod
        ├── go.sum
        ├── server.go
        ├── server_test.go
        └── user
            ├── user.pb.go
            └── user_grpc.pb.go
```
## Build and Run the Application

-  Generate Protbuf files 
    Ensure you have 'protoc installed. From the root directory run : 
```sh
protoc --go_out=server/user --go-grpc_out=server/user proto/user.proto

Python -m grpc_tools.protoc -Iprotofiles/ —python_out= client/python_pb_files/ —grpc_python_out=client/python_pb_files/ proto/user.proto
```
This generates protofiles for the .pb and grpc.pb files

- Build the Docker image:
```sh
docker build -t grpc-user-service .
```
- Run the Docker container:
```sh
docker run -d -p 50051:50051 grpc-user-service
```
- Verify the container is running:
```sh
docker ps
```
- check the logs of the application
```sh
docker logs <container-id>
```


## Running Tests

To run tests for the Go server 

- Navigate to the server directory:
```sh
cd server 
```
- Run the tests:
```sh
go test
```
## Accessing the Go Server Endpoints via grpcurl(without Python based client)

Use grpcurl to interact with the gRPC service. The following commands can be used to list services and call methods.

- List all services

```sh
grpcurl -plaintext localhost:50051 list
```

- List all methods in `UserService`
```sh
grpcurl -plaintext localhost:50051 list UserService
```
- Call GetUser endpoint
```sh
grpcurl -d '{"id": "1"}' -plaintext localhost:50051 UserService.GetUser
```

- Call SetUser endpoint

```sh
grpcurl -d '{
  "id": "1",
  "first_name": "John",
  "last_name": "Doe",
  "city": "Patna",
  "phone": "9873452",
  "height": 6.4,
  "married": true,
  "contact": {
    "home_addr": "123 Main St",
    "mob_num": "9876543210",
    "mail_id": "john.doe@example.com"
  }
}' -plaintext localhost:50051 UserService/SetUser
```
- Call UpdateUser endpoint

```sh 
grpcurl -d '{
  "id": "1",
  "first_name": "John",
  "last_name": "Doe",
  "city": "Boston",
  "phone": "9876543210",
  "height": 6.2,
  "married": false,
  "contact": {
    "home_addr": "456 Elm St",
    "mob_num": "1234567890",
    "mail_id": "john.doe@example.com"
  }
}' -plaintext localhost:50051 UserService.UpdateUser

```

- Call DeleteUser endpoint

```sh
grpcurl -d '{"id": "1"}' -plaintext localhost:50051 UserService.DeleteUser

```

- Call SearchUsers endpoint

```sh
grpcurl -d '{
  "city": "New York"
}' -plaintext localhost:50051 UserService/SearchUsers

```

- Call GetUsersByID
```sh
grpcurl -d '{
  "ids": ["1", "2", "3"]
}' -plaintext localhost:50051 UserService.GetUsersByID

```
## Configuration details

- Server Configuration :
    The server is written in Go and provides several user management gRPC endpoints.
    Ensure Go modules are correctly configured (go mod tidy is run during the build process).
    Client Configuration
    The client is written in Python and uses the generated protobuf files to interact with the server.
    Python dependencies are listed in requirements.txt.

- Docker Configuration :
    The Dockerfile is configured to build both the Go    server and Python client.
    The application exposes port 50051 for gRPC communications.
- Software and Compiler Versions
    - Go: 1.22 or higher
    - Python: 3.9 or higher
    - Docker: Latest version
    - grpcurl: Latest version
- Required Packages
    - Go
    Ensure the following Go packages are used:

       - google.golang.org/grpc
       - google.golang.org/protobuf
    - Python
    The requirements.txt should include:

       - grpcio
       - grpcio-tools

## See the Attachments folder for the Screenshot and video fil




