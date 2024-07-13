package main

import (
	"awesomeProject/accounts/models"
	"awesomeProject/proto"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"net"
	"sync"
)

type Server struct {
	proto.UnimplementedHandlerServer

	accounts map[string]*models.Account
	guard    *sync.RWMutex
	db       *sql.DB
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

		return nil, fmt.Errorf("account already exists")
	}

	_, err := s.db.ExecContext(ctx, "INSERT INTO accounts(name, balance) VALUES($1, $2)", req.Name, req.Amount)
	if err != nil {
		return nil, fmt.Errorf("error while inserting into db")
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

	_, err := s.db.ExecContext(ctx, "DELETE FROM accounts WHERE name=$1", req.Name)
	if err != nil {
		return nil, fmt.Errorf("error while deleting from db")
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

	_, err := s.db.ExecContext(ctx, "UPDATE accounts SET balance=$1 WHERE name=$2", req.NewAmount, req.Name)
	if err != nil {
		return nil, fmt.Errorf("error while changing account amount in db")
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

	_, err := s.db.ExecContext(ctx, "UPDATE accounts SET name=$1 WHERE name=$2", req.NewName, req.Name)
	if err != nil {
		return nil, fmt.Errorf("error while changing account name in db")
	}

	s.guard.Unlock()
	return &emptypb.Empty{}, nil
}

func main() {
	connectionString := "host=localhost port=5432 dbname=postgres user=postgres password=mysecretpassword"
	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		panic(err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 4567))
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()

	myServer := New()
	myServer.db = db
	proto.RegisterHandlerServer(s, myServer)

	ctx := context.Background()
	rows, err := db.QueryContext(ctx, "SELECT name, balance FROM accounts")
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		var account models.Account
		if err := rows.Scan(&account.Name, &account.Amount); err != nil {
			panic(err)
		}

		myServer.accounts[account.Name] = &account
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}

	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
