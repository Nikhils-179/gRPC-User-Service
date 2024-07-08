package main

import (
	"log"
	"net"

	"context"

	pb "github.com/Nikhils-179/grpc-user-service/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

const socket = "localhost:50051"

var userList []*pb.User

type Server struct {
	pb.UnimplementedUserServiceServer
}

func main() {
	userList = []*pb.User{
		{Id: "1", FirstName: "Steve", LastName: "Joe", City: "LA", Phone: "1234567890", Height: 5.8, Married: true},
		{Id: "2", FirstName: "John", LastName: "Wick", City: "NY", Phone: "0987654321", Height: 6.1, Married: false},
		{Id: "3", FirstName: "Kajol", LastName: "Kapoor", City: "Bangalore", Phone: "45326784", Height: 5.2, Married: true},
		{Id: "4", FirstName: "Sharukh", LastName: "Wadiya", City: "New Delhi", Phone: "239865432", Height: 6.5, Married: false},
	}
	lis, err := net.Listen("tcp", socket)
	if err != nil {
		log.Fatalf("Error Occured while listening on %s: %v", socket, err)
	}
	log.Printf("Server started at %s ", socket)

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &Server{})
	reflection.Register(s)
	if err = s.Serve(lis); err != nil {
		log.Fatalf("Error Occured while serving: %v", err)
	}
}

func (s *Server) GetUser(ctx context.Context, userID *pb.UserID) (*pb.User, error) {
	log.Printf("Hitting GetUser API with the user ID %s\n", userID.Id)
	for _, u := range userList {
		if u.Id == userID.Id {
			return u, nil
		}
	}
	return nil, status.Errorf(codes.NotFound, "Given user ID not found")
}

func (s *Server) SetUser(ctx context.Context, user *pb.User) (*pb.UserID, error) {
	log.Printf("Hitting SetUser API with the user ID %s\n", user.Id)
	for _, u := range userList {
		if u.Id == user.Id {
			return nil, status.Errorf(codes.AlreadyExists, "Given user ID already exists . Use UpdateUser API to update the user")
		}
		userList = append(userList, user)
	}
	return &pb.UserID{Id: user.Id}, nil
}

func (s *Server) ListUsers(ctx context.Context, empty *emptypb.Empty) (*pb.UserList, error) {
	log.Println("Hitting ListUsers API")
	return &pb.UserList{Users: userList}, nil
}

func (s *Server) UpdateUser(ctx context.Context, user *pb.User) (*emptypb.Empty, error) {
	log.Printf("Hitting UpdateUser with the user ID %s\n", user.Id)
	for i, u := range userList {
		if u.Id == user.Id {
			userList[i] = user
			log.Printf("Updated user %v\n", user)
			return &emptypb.Empty{}, nil
		}
	}
	return nil, status.Errorf(codes.NotFound, "Given user ID not found to update")
}

func (s *Server) DeleteUser(ctx context.Context, userID *pb.UserID) (*emptypb.Empty, error) {
	log.Printf("Hitting DeleteUser API with the user ID %s\n", userID.Id)
	for i, u := range userList {
		if u.Id == userID.Id {
			userList = append(userList[:i], userList[i+1:]...)
			log.Printf("Deleted user:%v\n", userID)
			return &emptypb.Empty{}, nil
		}
	}
	return nil, status.Errorf(codes.NotFound, "Given user ID not found to delete")
}

func (s *Server) SearchUsers(ctx context.Context, criteria *pb.SearchCriteria) (*pb.UserList, error) {
	log.Printf("Hiiting SearchUsers API with the criteria: %v\n", criteria)
	var filteredUsers []*pb.User
	for _, u := range userList {
		if (criteria.City == "" || u.City == criteria.City) && (criteria.Phone == "" || u.Phone == criteria.Phone) &&
			(criteria.Married == u.Married) {
			filteredUsers = append(filteredUsers, u)
		}
	}
	return &pb.UserList{Users: filteredUsers}, nil
}

func (s *Server) GetUsersByID(ctx context.Context, usersIDs *pb.UserIDs) (*pb.UserList, error) {
	log.Printf("Hitting GetUsersByID with the IDs: %v\n", usersIDs.Ids)
	var foundUsers []*pb.User
	for _, id := range usersIDs.Ids {
		for _, u := range userList {
			if u.Id == id {
				foundUsers = append(foundUsers, u)
			}
		}
	}
	if len(foundUsers) == 0 {
		return nil, status.Errorf(codes.NotFound, "No users found for given IDs")
	}
	return &pb.UserList{Users: foundUsers}, nil
}
