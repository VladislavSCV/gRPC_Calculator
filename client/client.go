package main

import (
    "context"
    "log"
    
    pb "github.com/VladislavSCV/gRPC_Calculator/grpc" // Путь к сгенерированному коду нашего протобуфа
    "google.golang.org/grpc"
)

func main() {
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure()) // Устанавливаем соединение с сервером
    if err != nil {
        log.Fatalf("could not connect: %v", err)
    }
    defer conn.Close()

    client := pb.NewCalculatorClient(conn)

    req := &pb.ReqNums{
        Num1: 10,
        Num2: 0,
    }
    res, err := client.Multip(context.Background(), req)
    if err != nil {
        log.Fatalf("could not add: %v", err)
    }
	
    log.Printf("Result: %d", res.Result)
}
