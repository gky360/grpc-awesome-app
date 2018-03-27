package server

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/izumin5210/grapi/pkg/grapiserver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api_pb "github.com/gky360/grpc-awesome-app/api"
)

// NewBookServiceServer creates a new BookServiceServer instance.
func NewBookServiceServer() interface {
	api_pb.BookServiceServer
	grapiserver.Server
} {
	return &bookServiceServerImpl{}
}

type bookServiceServerImpl struct {
}

// RegisterWithServer implements grapiserver.Server.RegisterWithServer.
func (s *bookServiceServerImpl) RegisterWithServer(grpcSvr *grpc.Server) {
	api_pb.RegisterBookServiceServer(grpcSvr, s)
}

// RegisterWithHandler implements grapiserver.Server.RegisterWithHandler.
func (s *bookServiceServerImpl) RegisterWithHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return api_pb.RegisterBookServiceHandler(ctx, mux, conn)
}

func (s *bookServiceServerImpl) ListBooks(ctx context.Context, req *api_pb.ListBooksRequest) (*api_pb.ListBooksResponse, error) {
	// TODO: Not yet implemented.
	// return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
	resp := &api_pb.ListBooksResponse{}
	resp.Books = []*api_pb.Book{
		&api_pb.Book{BookId: "1", Title: "Star Wars I", Author: "George Lucas"},
		&api_pb.Book{BookId: "2", Title: "Star Wars II", Author: "George Lucas"},
	}
	return resp, nil
}

func (s *bookServiceServerImpl) GetBook(ctx context.Context, req *api_pb.GetBookRequest) (*api_pb.Book, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *bookServiceServerImpl) CreateBook(ctx context.Context, req *api_pb.CreateBookRequest) (*api_pb.Book, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *bookServiceServerImpl) UpdateBook(ctx context.Context, req *api_pb.UpdateBookRequest) (*api_pb.Book, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *bookServiceServerImpl) DeleteBook(ctx context.Context, req *api_pb.DeleteBookRequest) (*empty.Empty, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}
