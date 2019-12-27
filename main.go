// shippy-service-consignment/main.go
package main

// --------------------------------------------------------
// --------------------------------------------------------
// IMPORTS
// --------------------------------------------------------
import (
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
	port = "localhost:50051"
)

// --------------------------------------------------------
// --------------------------------------------------------
// FUNCTIONS
// --------------------------------------------------------
func main() {
	// setup gRPC server
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen to: %v", err)
	}

	server := grpc.NewServer()

	// register service to gRPC server
	repo := &Repository{}
	serv := &Service{repo}
	pb.RegisterShippingServiceServer(server, serv)

	// register reflection ervice on gRPC server
	reflection.Register(server)

	log.Println("Running on port:", port)
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
