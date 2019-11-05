package internal

import (
	"context"
	{{ .service.Package }} "{{ .service.Package }}/generated"
	"fmt"
	"flag"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/timestamp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

var (
	MongoAddress = flag.String("{{ .service.EnvVar }}_MONGO_ADDR", "mongodb://{{ .mongo.Credentials }}{{ .mongo.Address }}", "mongo address")
	MongoPingTimeout = flag.Duration("{{ .service.EnvVar }}_MONGO_PING_TIMEOUT", "2", "mongo ping timeout")
)

type Server struct {
	{{ .service.Package }}.AppServer
	{{ .service.Resource }}Collection *mongo.Collection
}

type {{ .service.Resource }} struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	Enabled   bool               `bson:"enabled"`
	Type      int32              `bson:"type"`
	Name      string             `bson:"name"`
	CreatedAt *time.Time         `bson:"created_at,omitempty"`
	UpdatedAt *time.Time         `bson:"updated_at,omitempty"`
}

func (s *Server) Create(ctx context.Context, req *{{ .service.Package }}.{{ .service.Resource }}) (*{{ .service.Package }}.Id, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	var now = time.Now()
	var doc = {{ .service.Resource }}{
		Enabled:   req.GetEnabled(),
		Type:      req.GetType(),
		Name:      req.GetName(),
		CreatedAt: &now,
	}
	result, err := s.{{ .service.Resource }}Collection.InsertOne(context.Background(), doc)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &{{ .service.Package }}.Id{
		Id: result.InsertedID.(primitive.ObjectID).Hex(),
	}, nil
}

func (s *Server) Get(ctx context.Context, req *{{ .service.Package }}.Id) (*{{ .service.Package }}.{{ .service.Resource }}, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	oid, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}
	result := s.{{ .service.Resource }}Collection.FindOne(context.Background(), bson.M{"_id": oid})
	// Create an empty BlogItem to write our decode result to
	data := {{ .service.Resource }}{}
	// decode and write to data
	if err := result.Decode(&data); err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find {{ .service.Resource }} with Object Id %s: %v", req.GetId(), err))
	}
	var createdAt *timestamp.Timestamp
	var updatedAt *timestamp.Timestamp
	createdAt, err = ptypes.TimestampProto(*data.CreatedAt)
	if err != nil {
		return nil, err
	}
	if data.UpdatedAt != nil {
		updatedAt, err = ptypes.TimestampProto(*data.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}

	response := &{{ .service.Package }}.{{ .service.Resource }}{
		Id:        oid.Hex(),
		Enabled:   data.Enabled,
		Type:      data.Type,
		Name:      data.Name,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	return response, nil
}

func (s *Server) List(ctx context.Context, req *{{ .service.Package }}.Criteria) (*{{ .service.Package }}.{{ .service.Resources }}, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	// collection.Find returns a cursor for our (empty) query
	cursor, err := s.{{ .service.Resource }}Collection.Find(context.Background(), bson.M{
		"enabled": req.Enabled,
		"type":    req.Type,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Unknown internal error: %v", err))
	}
	// An expression with defer will be called at the end of the function
	defer cursor.Close(context.Background())

	var items = make([]*{{ .service.Package }}.{{ .service.Resource }}, 0)
	var createdAt *timestamp.Timestamp
	var updatedAt *timestamp.Timestamp
	// cursor.Next() returns a boolean, if false there are no more items and loop will break
	for cursor.Next(context.Background()) {
		var data = &{{ .service.Resource }}{}
		// Decode the data at the current pointer and write it to data
		err := cursor.Decode(data)
		// check error
		if err != nil {
			return nil, status.Errorf(codes.Unavailable, fmt.Sprintf("Could not decode data: %v", err))
		}
		createdAt, err = ptypes.TimestampProto(*data.CreatedAt)
		if err != nil {
			return nil, status.Errorf(codes.Unavailable, fmt.Sprintf("Could not decode createdAt: %v", err))
		}
		if data.UpdatedAt != nil {
			updatedAt, err = ptypes.TimestampProto(*data.UpdatedAt)
			if err != nil {
				return nil, status.Errorf(codes.Unavailable, fmt.Sprintf("Could not decode updatedAt: %v", err))
			}
		}

		items = append(items, &{{ .service.Package }}.{{ .service.Resource }}{
			Id:        data.Id.Hex(),
			Enabled:   data.Enabled,
			Type:      data.Type,
			Name:      data.Name,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		})
	}
	return &{{ .service.Package }}.{{ .service.Resources }}{Items: items}, nil
}

func (s *Server) Update(ctx context.Context, req *{{ .service.Package }}.{{ .service.Resource }}) (*{{ .service.Package }}.{{ .service.Resource }}, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	// Convert the Id string to a MongoDB ObjectId
	oid, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Could not convert the supplied blog id to a MongoDB ObjectId: %v", err),
		)
	}

	// Convert the data to be updated into an unordered Bson document
	update := bson.M{
		"enabled":    req.GetEnabled(),
		"type":       req.GetType(),
		"name":       req.GetName(),
		"updated_at": time.Now(),
	}

	// Convert the oid into an unordered bson document to search by id
	filter := bson.M{"_id": oid}

	// Result is the BSON encoded result
	// To return the updated document instead of original we have to add options.
	result := s.{{ .service.Resource }}Collection.FindOneAndUpdate(ctx, filter, bson.M{"$set": update}, options.FindOneAndUpdate().SetReturnDocument(1))

	// Decode result and write it to 'data'
	data := {{ .service.Resource }}{}
	err = result.Decode(&data)
	if err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Could not find {{ .service.Resource }} with supplied ID: %v", err),
		)
	}
	var createdAt *timestamp.Timestamp
	var updatedAt *timestamp.Timestamp
	createdAt, err = ptypes.TimestampProto(*data.CreatedAt)
	if err != nil {
		return nil, err
	}
	if data.UpdatedAt != nil {
		updatedAt, err = ptypes.TimestampProto(*data.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}
	return &{{ .service.Package }}.{{ .service.Resource }}{
		Id:        oid.Hex(),
		Enabled:   data.Enabled,
		Type:      data.Type,
		Name:      data.Name,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}

func (s *Server) Delete(ctx context.Context, req *{{ .service.Package }}.Id) (*empty.Empty, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	oid, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}
	_, err = s.{{ .service.Resource }}Collection.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find/delete blog with id %s: %v", req.GetId(), err))
	}
	return &empty.Empty{}, nil
}

func (s *Server) connectToMongo() {
	fmt.Print("Connecting to mongo... ")
	client, err := mongo.NewClient(options.Client().ApplyURI(*MongoAddress))
	if err != nil {
		fmt.Println("error")
		panic(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), *MongoPingTimeout*time.Second)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println("error... ping failed")
		panic(err)
	}

	fmt.Println("success")
	s.{{ .service.Resource }}Collection = client.Database("{{ .mongo.Database }}").Collection("{{ .mongo.Collection }}")
}

func NewServer() {{ .service.Package }}.AppServer {
	server := Server{}
	server.connectToMongo()
	return &server
}
