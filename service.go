// shippy-service-consignment/service.go
package main

// --------------------------------------------------------
// --------------------------------------------------------
// IMPORTS
// --------------------------------------------------------
import (
	"context"

	// import the generated protobuf code
	pb "github.com/athirah-yahya/shippy-consignment-service/proto/consignment"
)

// --------------------------------------------------------
// --------------------------------------------------------
// STRUCTS
// --------------------------------------------------------
type Service struct {
	repo *Repository
}

// --------------------------------------------------------
// --------------------------------------------------------
// FUNCTIONS
// --------------------------------------------------------
/**
 * Create new consignment
 *
 * @param	ctx *context.Context
 * @param	req *pb.Consignment
 *
 * @return	*pb.response
 * @return	error
 */
func (serv *Service) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error) {
	// save consignment
	consignment, err := serv.repo.Create(req)
	if err != nil {
		return nil, err
	}

	return &pb.Response{
		Created:     true,
		Consignment: consignment,
	}, nil
}
