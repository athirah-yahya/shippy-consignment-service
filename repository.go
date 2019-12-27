// shippy-service-consignment/repository.go
package main

// --------------------------------------------------------
// --------------------------------------------------------
// IMPORTS
// --------------------------------------------------------
import (
	"sync"

	// import the generated protobuf code
	pb "consignment-service/proto/consignment"
)

// --------------------------------------------------------
// --------------------------------------------------------
// INTERFACES
// --------------------------------------------------------
type repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
}

// --------------------------------------------------------
// --------------------------------------------------------
// STRUCTS
// --------------------------------------------------------
type Repository struct {
	mu           sync.RWMutex
	consignments []*pb.Consignment
}

// --------------------------------------------------------
// --------------------------------------------------------
// FUNCTIONS
// --------------------------------------------------------

/**
 * Create new consignment
 *
 * @param	consignment *pb.Consignment
 *
 * @return	*pb.Consignment
 * @return	error
 */
func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	updated := append(repo.consignments, consignment)
	repo.consignments = updated

	return consignment, nil
}
