/*
Copyright (c) 2018 SAP SE or an SAP affiliate company. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
