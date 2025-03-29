
# Go Client for Power Supply gRPC Service

This is a simple Go client that connects to a .NET gRPC server for controlling a programmable power supply. It sends commands `OpenConnection`, `SetVolts`, and more over gRPC.

---

## 📁 Project Structure

```
go-client-for-gRPC/
├── device.proto             # gRPC protocol definition (copied from .NET server)
├── powersupplyclient/                   # Generated Go code from .proto file
│   ├── device.pb.go
│   └── device_grpc.pb.go
├── main.go                  # The actual Go client
├── go.mod                   # Go module definition
└── README.md                # You're here!
```

---

## ✅ Prerequisites

- Go 1.18+ installed
- `protoc` (Protocol Buffers compiler) installed
- Go plugins for protoc:
  ```bash
  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
  ```

---

## 🚀 Setup Instructions

### 1. Initialize Go Module

```bash
go mod init go-client-for-gRPC
```

### 2. Update `.proto` File

Make sure `device.proto` contains:

```proto
option go_package = "go-client-for-gRPC/power";
```

And matches your `go.mod` module name.

### 3. Generate Go Code from Proto

```bash
protoc --go_out=. --go-grpc_out=. device.proto
```

This creates the `power/` folder with client code.

---

## 🧪 Run the Client

Make sure your .NET gRPC server is running on `localhost:5248`, then run:

```bash
go run main.go
```

You should see logs like:

```
OpenConnection: Connection opened. (code 0)
SetVolts: Voltage set. (code 0)
```

---

## 📞 gRPC Calls Used

- `OpenConnection(Port)`
- `SetVolts(Channel, Value)`


You can extend this by modifying `main.go` to call more gRPC methods defined in `device.proto`.

---

