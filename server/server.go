package main

import (
    "context"
    "log"
    "net"

    pb "github.com/VladislavSCV/gRPC_Calculator/grpc" // Путь к сгенерированному коду нашего протобуфа
    "google.golang.org/grpc"
)

type server struct{
    pb.UnimplementedCalculatorServer
}

func (s *server) Add(ctx context.Context, req *pb.ReqNums) (*pb.CalculResult, error) {
    result := req.Num1 + req.Num2
    return &pb.CalculResult{Result: result}, nil
}

func (s *server) Minus(ctx context.Context, req *pb.ReqNums) (*pb.CalculResult, error) {
    result := req.Num1 - req.Num2
    return &pb.CalculResult{Result: result}, nil
}

func (s *server) Devide(ctx context.Context, req *pb.ReqNums) (*pb.CalculResult, error) {
    if req.Num2 == 0 {
        log.Print("Devide by zero!")
        return nil, nil
    } else {
        result := req.Num1 / req.Num2
        return &pb.CalculResult{Result: result}, nil
    }
}

func (s *server) Multip(ctx context.Context, req *pb.ReqNums) (*pb.CalculResult, error) {
    result := req.Num1 * req.Num2
    return &pb.CalculResult{Result: result}, nil
}

// Добавляем метод mustEmbedUnimplementedCalculatorServer для удовлетворения интерфейса
func (s *server) mustEmbedUnimplementedCalculatorServer() {}

func main() {
    var lis net.Listener
    var err error

    if lis, err = net.Listen("tcp", ":50051"); err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    s := grpc.NewServer()
    pb.RegisterCalculatorServer(s, &server{}) // Регистрируем наш сервер

    defer func() {
        if r := recover(); r != nil {
            log.Println("Recovered in f", r)
        }
    }()

    log.Println("Server is running on port 50051...")

    if err = s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}

