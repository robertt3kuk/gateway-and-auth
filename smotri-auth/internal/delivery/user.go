package delivery

import (
	"context"
	"net/http"

	"smotri-auth/internal/entity"
	"smotri-auth/internal/pb"
)

func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	var user entity.User
	user.Email = req.Email
	user.Password = req.Password
	_, err := s.T.CreateUser(ctx, user)
	if err != nil {
		return &pb.RegisterResponse{
			Status: http.StatusBadRequest,
		}, nil

	}
	return &pb.RegisterResponse{
		Status: http.StatusCreated,
	}, nil
}

func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	email, password := req.Email, req.Password
	token, err := s.T.GenerateToken(ctx, email, password)
	if err != nil {
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "User not found or incorrect password",
		}, nil
	}

	return &pb.LoginResponse{
		Status: http.StatusOK,
		Token:  token,
	}, nil
}

func (s *Server) Validate(ctx context.Context, req *pb.TokenValidateRequest) (*pb.TokenValidateResponse, error) {
	ID, err := s.T.ParseToken(req.Token)
	if err != nil {
		return &pb.TokenValidateResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	}

	return &pb.TokenValidateResponse{
		Status: http.StatusOK,
		UserId: int64(ID),
	}, nil
}
