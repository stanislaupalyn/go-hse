package main

import (
	"awesomeProject/accounts/models"
	"awesomeProject/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"net"
	"sync"
)

type Server struct {
	proto.UnimplementedHandlerServer

	accounts map[string]*models.Account
	guard    *sync.RWMutex
}

func New() *Server {
	return &Server{
		accounts: make(map[string]*models.Account),
		guard:    &sync.RWMutex{},
	}
}

func (s *Server) CreateAccount(ctx context.Context, req *proto.CreateAccountRequest) (*emptypb.Empty, error) {
	if len(req.Name) == 0 {
		return &emptypb.Empty{}, fmt.Errorf("empty name")
	}

	s.guard.Lock()

	if _, ok := s.accounts[req.Name]; ok {
		s.guard.Unlock()

		return &emptypb.Empty{}, fmt.Errorf("account already exists")
	}

	s.accounts[req.Name] = &models.Account{
		Name:   req.Name,
		Amount: int(req.Amount),
	}

	s.guard.Unlock()
	fmt.Println("Account created:", req.Name, req.Amount)

	return &emptypb.Empty{}, nil
}

func (s *Server) GetAccount(ctx context.Context, req *proto.GetAccountRequest) (*proto.GetAccountResponse, error) {
	s.guard.RLock()

	account, ok := s.accounts[req.Name]

	s.guard.RUnlock()

	if !ok {
		return nil, fmt.Errorf("account not found")
	}

	return &proto.GetAccountResponse{
		Name:   account.Name,
		Amount: int32(account.Amount),
	}, nil
}

func (s *Server) DeleteAccount(ctx context.Context, req *proto.DeleteAccountRequest) (*emptypb.Empty, error) {
	s.guard.Lock()

	if _, ok := s.accounts[req.Name]; !ok {
		s.guard.Unlock()

		return nil, fmt.Errorf("account doesn't exist")
	}

	delete(s.accounts, req.Name)

	s.guard.Unlock()
	return &emptypb.Empty{}, nil
}

func (s *Server) ChangeAccountAmount(ctx context.Context, req *proto.ChangeAmountRequest) (*emptypb.Empty, error) {
	s.guard.Lock()

	if _, ok := s.accounts[req.Name]; !ok {
		s.guard.Unlock()

		return nil, fmt.Errorf("account doesn't exist")
	}

	s.accounts[req.Name].Amount = int(req.NewAmount)

	s.guard.Unlock()
	return &emptypb.Empty{}, nil
}

func (s *Server) ChangeAccountName(ctx context.Context, req *proto.ChangeNameRequest) (*emptypb.Empty, error) {
	s.guard.Lock()

	if _, ok := s.accounts[req.Name]; !ok {
		s.guard.Unlock()

		return nil, fmt.Errorf("account doesn't exist")
	}

	if _, ok := s.accounts[req.NewName]; ok {
		s.guard.Unlock()

		return nil, fmt.Errorf("account with new name already exists")
	}

	account := s.accounts[req.Name]
	account.Name = req.NewName
	s.accounts[req.NewName] = account
	delete(s.accounts, req.Name)

	s.guard.Unlock()
	return &emptypb.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 4567))
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	proto.RegisterHandlerServer(s, New())
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
