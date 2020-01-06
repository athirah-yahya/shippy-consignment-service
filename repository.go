// shippy-service-consignment/repository.go
package main

// --------------------------------------------------------
// --------------------------------------------------------
// IMPORTS
// --------------------------------------------------------
import (
	"sync"

	// import the generated protobuf code
	pb "github.com/athirah-yahya/shippy-consignment-service/proto/consignment"
)

// --------------------------------------------------------
// --------------------------------------------------------
// INTERFACES
// --------------------------------------------------------
type repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
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

	repo.consignments = append(repo.consignments, consignment)

	return consignment, nil
}

// --------------------------------------------------------

/**
 * Get all consignments
 *
 * @return	[]*pb.Consignment
 */
func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}
