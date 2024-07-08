package main

import (
	"context"
	"testing"

	pb "github.com/Nikhils-179/grpc-user-service/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)



func setupServer() *Server{
	s := &Server{}
	userList = []*pb.User{
		{Id: "1", FirstName: "Steve", LastName: "Joe", City: "LA", Phone: "1234567890", Height: 5.8, Married: true},
		{Id: "2", FirstName: "John" , LastName: "Wick" ,City: "NY", Phone: "0987654321", Height: 6.1, Married: false},
	}
	return s
}

func TestGetUser(t *testing.T){
	s := setupServer()

	tests := []struct{
		name        string
		userID      *pb.UserID
		wantUser    *pb.User
		wantErr     codes.Code
	}{
		{"User exists" , &pb.UserID{Id : "1"} , userList[0], codes.OK},
		{"User doesnot exist",&pb.UserID{Id: "3"} , nil , codes.NotFound}, 
	}


	for _ ,tt := range tests{
		t.Run(tt.name , func(t *testing.T) {
			user, err := s.GetUser(context.Background(), tt.userID)
			if status.Code(err) != tt.wantErr{
				t.Errorf("got error code %v, want %v",status.Code(err),tt.wantErr)
			}
			if tt.wantErr == codes.OK && user.Id != tt.wantUser.Id{
				t.Errorf("got user %v, want %v",user,tt.wantUser)
			}
		})
	}
}


func TestSetUser(t *testing.T){
	s := setupServer()

	tests := []struct{
		name string
		user *pb.User
		wantErr codes.Code
	}{
		{"New user" , &pb.User{Id: "3" , FirstName: "Nikhil" ,LastName: "Shrivastava",City: "Bangalore", Phone: "12345",Height: 5.9, Married: false},codes.OK},
		{"Duplicate user", &pb.User{Id: "1", FirstName: "Steve", LastName: "Joe", City: "LA", Phone: "1234567890", Height: 5.8, Married: true},codes.AlreadyExists},
	}

	for _ ,tt := range tests{
		t.Run(tt.name , func(t *testing.T){
			_ , err := s.SetUser(context.Background(),tt.user)
			if status.Code(err) != tt.wantErr{
				t.Errorf("got error code %v, want %v", status.Code(err),tt.wantErr)
			}
		})
	}
}

func TestUpdateUser(t *testing.T) {
	s := setupServer()

	tests := []struct {
		name    string
		user    *pb.User
		wantErr codes.Code
	}{
		{"Update existing user", &pb.User{Id: "1", FirstName: "Steve", LastName : "Joe",City: "SF", Phone: "1234567890", Height: 5.8, Married: true}, codes.OK},
		{"Update non-existing user", &pb.User{Id: "3", FirstName: "Nikhil", LastName: "Shrivastava",City: "SF", Phone: "1122334455", Height: 5.6, Married: false}, codes.NotFound},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := s.UpdateUser(context.Background(), tt.user)
			if status.Code(err) != tt.wantErr {
				t.Errorf("got error code %v, want %v", status.Code(err), tt.wantErr)
			}
		})
	}
}

func TestDeleteUser(t *testing.T) {
	s := setupServer()

	tests := []struct {
		name    string
		userID  *pb.UserID
		wantErr codes.Code
	}{
		{"Delete existing user", &pb.UserID{Id: "1"}, codes.OK},
		{"Delete non-existing user", &pb.UserID{Id: "3"}, codes.NotFound},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := s.DeleteUser(context.Background(), tt.userID)
			if status.Code(err) != tt.wantErr {
				t.Errorf("got error code %v, want %v", status.Code(err), tt.wantErr)
			}
		})
	}
}

func TestSearchUsers(t *testing.T) {
	s := setupServer()

	tests := []struct {
		name       string
		criteria   *pb.SearchCriteria
		wantResult int
	}{
		{"Search by city", &pb.SearchCriteria{City: "LA"}, 1},
		{"Search by phone", &pb.SearchCriteria{Phone: "0987654321"}, 1},
		{"Search by marital status", &pb.SearchCriteria{Married: false}, 1},
		{"Search by non-existing criteria", &pb.SearchCriteria{City: "Muzaffarpur"}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := s.SearchUsers(context.Background(), tt.criteria)
			if err != nil {
				t.Errorf("SearchUsers returned error: %v", err)
			}
			if len(resp.Users) != tt.wantResult {
				t.Errorf("got %v users, want %v", len(resp.Users), tt.wantResult)
			}
		})
	}
}

func TestGetUsersByID(t *testing.T) {
	s := setupServer()

	tests := []struct {
		name       string
		userIDs    *pb.UserIDs
		wantResult int
		wantErr    codes.Code
	}{
		{"Get users by IDs", &pb.UserIDs{Ids: []string{"1", "2"}}, 2, codes.OK},
		{"Get users by non-existing IDs", &pb.UserIDs{Ids: []string{"3"}}, 0, codes.NotFound},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := s.GetUsersByID(context.Background(), tt.userIDs)
			if status.Code(err) != tt.wantErr {
				t.Errorf("got error code %v, want %v", status.Code(err), tt.wantErr)
			}
			if tt.wantErr == codes.OK && len(resp.Users) != tt.wantResult {
				t.Errorf("got %v users, want %v", len(resp.Users), tt.wantResult)
			}
		})
	}
}

