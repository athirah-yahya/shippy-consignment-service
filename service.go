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
func (serv *Service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	// save consignment
	consignment, err := serv.repo.Create(req)
	if err != nil {
		return err
	}

	// set response
	res.Created = true
	res.Consignment = consignment

	return nil
}

// --------------------------------------------------------

/**
 * Get consignment list
 *
 * @param	ctx *context.Context
 * @param	req *pb.GetRequest
 *
 * @return	*pb.response
 * @return	error
 */
func (serv *Service) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	// get consignments
	consignments := serv.repo.GetAll()

	// set response
	res.Consignments = consignments

	return nil
}
