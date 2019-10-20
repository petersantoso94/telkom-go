package service

import(
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "telkom-go/pkg/proto/user-service"
	"telkom-go/service/user-service/user"
)

// Server is a struct to implement "user" of proto
type Server struct {
	userServer user.Users
}

// CreateServer is a function to get reference address of "Server" struct
func CreateServer(userService user.Users) *Server {
	return &Server{userServer: userService}
}


func (s *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.EmptyResponse, error) {
	return nil,nil
}
func (s *Server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	Users, err := s.userServer.GetUser(&user.GetUserOption{
		Email:req.Email,
	})

	if err != nil {
		switch err.(type) {
		case *user.NotFoundError:
			return nil, status.Error(codes.NotFound, err.Error())
		case *user.InternalError:
			return nil, status.Error(codes.Internal, err.Error())
		default:
			return nil, status.Error(codes.Unimplemented, err.Error())
		}
	}

	var resp []*pb.User

	for _,us := range Users.Users {
		resp = append(resp, &pb.User{
			Email: us.Email,
			LockIp: us.LockIP,
			Password: us.Password,
			Position: us.Position,
		})
	}

	return &pb.GetUserResponse{
		User: resp,
	}, nil
}
func (s *Server) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.EmptyResponse, error) {
	return nil,nil
}
func (s *Server) DeleteUser(ctx context.Context, req *pb.DeleteRequest) (*pb.EmptyResponse, error) {
	return nil,nil
}
