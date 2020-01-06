// shippy-service-consignment/main.go
package main

// --------------------------------------------------------
// --------------------------------------------------------
// IMPORTS
// --------------------------------------------------------
import (
	"log"

	// import the generated protobuf code
	pb "github.com/athirah-yahya/shippy-consignment-service/proto/consignment"
	"github.com/micro/go-micro"
)

// --------------------------------------------------------
// --------------------------------------------------------
// FUNCTIONS
// --------------------------------------------------------
func main() {
	// register micro-service
	server := micro.NewService(micro.Name("shippy.consignment.service"))
	server.Init()

	// register service to gRPC server
	repo := &Repository{}
	serv := &Service{repo}
	pb.RegisterShippingServiceHandler(server.Server(), serv)

	if err := server.Run(); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
