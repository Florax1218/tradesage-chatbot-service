package main

import (
	"log"
	"net"
	"os"

	pb "github.com/Ayush10/tradesage-chatbot-service/internal/pb"
	"github.com/Ayush10/tradesage-chatbot-service/internal/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
    port := os.Getenv("CHATBOT_SERVICE_PORT")
    if port == "" {
        port = "50051"
    }

    lis, err := net.Listen("tcp", ":"+port)
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer()

    chatbotServer, err := services.NewChatbotServiceServer()
    if err != nil {
        log.Fatalf("Failed to create chatbot server: %v", err)
    }
    pb.RegisterChatbotServiceServer(grpcServer, chatbotServer)

    // Register reflection service on gRPC server
    reflection.Register(grpcServer)

    log.Printf("TradeSage AI - Chatbot Integration Service is running on port %s.", port)

    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}
