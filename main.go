// shippy-service-consignment/main.go
package main

// --------------------------------------------------------
// --------------------------------------------------------
// IMPORTS
// --------------------------------------------------------
import (
	"fmt"
	"log"
	"net"

	// import the generated protobuf code
	pb "github.com/athirah-yahya/shippy-consignment-service/proto/consignment"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// --------------------------------------------------------
// --------------------------------------------------------
// CONSTANTS
// --------------------------------------------------------
const (
	host = "0.0.0.0"
	port = "50051"
)

// --------------------------------------------------------
// --------------------------------------------------------
// FUNCTIONS
// --------------------------------------------------------
func main() {
	// setup gRPC server
	address := fmt.Sprintf("%s:%s", host, port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen to: %v", err)
	}

	server := grpc.NewServer()

	// register service to gRPC server
	repo := &Repository{}
	serv := &Service{repo}
	pb.RegisterShippingServiceServer(server, serv)

	// register reflection service on gRPC server
	reflection.Register(server)

	log.Println("Running on port:", port)
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
