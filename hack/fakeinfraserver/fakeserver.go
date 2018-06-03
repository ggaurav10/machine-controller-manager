package main

import (
	"log"

	"net"

	"golang.org/x/net/context"

	pb "github.com/gardener/machine-controller-manager/pkg/grpc/infrapb"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// Server represents the gRPC server
type Server struct {
}

// Create cretes an instance
func (s *Server) Create(ctx context.Context, in *pb.CreateParams) (*pb.CreateResp, error) {
	log.Printf("Receive message %s", in.Name)
	return &pb.CreateResp{
		ProviderID: "Create",
	}, nil
}

// Delete deletes provided instance
func (s *Server) Delete(ctx context.Context, in *pb.DeleteParams) (*pb.DeleteResp, error) {
	log.Printf("Receive message %s", in.Name)
	return &pb.DeleteResp{
		Error: 0,
	}, nil
}

// List returns list of instances
func (s *Server) List(ctx context.Context, in *pb.ListParams) (*pb.ListResp, error) {
	log.Printf("Receive message %s", in.Name)
	return &pb.ListResp{
		Message: "List",
	}, nil
}

// ShareMeta takes metadata and caches
func (s *Server) ShareMeta(ctx context.Context, in *pb.Metadata) (*pb.ErrorResp, error) {
	log.Printf("Receive message %s", in.Type)
	return &pb.ErrorResp{
		Err:     0,
		Message: "Cached",
	}, nil

}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// Creates a new gRPC server
	s := grpc.NewServer()
	pb.RegisterInfragrpcServer(s, &Server{})
	s.Serve(lis)
}
