package main

import (
    "context"
    "log"
    "time"

    "google.golang.org/grpc"
    pb "powersupplyclient/power" 
)

func main() {
    conn, err := grpc.Dial("localhost:5248", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("Could not connect to server: %v", err)
    }
    defer conn.Close()

    client := pb.NewPowerSupplyClient(conn)

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    // 1. Open Connection
    res, err := client.OpenConnection(ctx, &pb.OpenRequest{Port: "COM5"})
    if err != nil {
        log.Fatalf("OpenConnection failed: %v", err)
    }
    log.Printf("OpenConnection: %s (code %d)", res.Message, res.Code)

    // 2. Set Voltage
    setRes, err := client.SetVolts(ctx, &pb.SetRequest{Channel: 0, Value: 7.0})
    if err != nil {
        log.Fatalf("SetVolts failed: %v", err)
    }
    log.Printf("SetVolts: %s (code %d)", setRes.Message, setRes.Code)
}
