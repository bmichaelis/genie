package internal

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/wrappers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	service "{{ .service.Package }}/generated"
	"log"
)

type Server struct {
	service.{{ .service.Resource }}Server
	{{ .service.Resource }}Collection *mongo.Collection
}

func (s *Server) SayHello(ctx context.Context, req *service.HelloRequest) (*service.HelloReply, error) {
	log.Printf("Received: %v", req.GetName())
	name := req.GetName()
	if name == "" {
		name = service.HelloRequest_world.String()
	}
	return &service.HelloReply{Message: "Hello, " + name}, nil
}

func (*Server) Create(ctx context.Context, req *service.User) (*wrappers.StringValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}

func (*Server) Get(ctx context.Context, req *service.Id) (*service.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func (*Server) Search(req *service.SearchRequest, srv service.{{ .service.Resource }}_SearchServer) error {
	return status.Errorf(codes.Unimplemented, "method Search not implemented")
}

func (*Server) Update(ctx context.Context, req *service.User) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}

func (*Server) Delete(ctx context.Context, req *service.Id) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func (*Server) Restore(ctx context.Context, req *service.Id) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restore not implemented")
}

func (s *Server) connectToMongo() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		panic(err)
	}

	s.{{ .service.Resource  }}Collection = client.Database("{{ .mongo.Database }}").Collection("{{ .mongo.CollectionName }}")
}

func NewServer() service.{{ .service.Resource }}Server {
	server := Server{}
	server.connectToMongo()
	return &server
}
