package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"sync"

	pb "awesomeProject/proto"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

var (
	port = flag.Int("port", 1323, "The server port")
)

type server struct {
	pb.AwesomeProjectServer
	accounts map[string]int64
	guard    *sync.RWMutex
}

func (s *server) CreateAccount(ctx context.Context, in *pb.CreateAccountRequest) (*pb.Empty, error) {
	s.guard.Lock()
	defer s.guard.Unlock()
	if _, found := s.accounts[in.GetName()]; found {
		return &pb.Empty{}, fmt.Errorf("user %s already exists", in.GetName())
	}
	s.accounts[in.GetName()] = in.GetAmount()
	return &pb.Empty{}, nil
}

func (s *server) GetAccount(ctx context.Context, in *pb.GetAccountRequest) (*pb.GetAccountResponse, error) {
	s.guard.RLock()
	defer s.guard.RUnlock()
	if amount, found := s.accounts[in.GetName()]; found {
		return &pb.GetAccountResponse{
			Name:   proto.String(in.GetName()),
			Amount: proto.Int64(amount),
		}, nil
	}
	return nil, fmt.Errorf("account \"%s\" not found", in.GetName())
}

func (s *server) DeleteAccount(ctx context.Context, in *pb.DeleteAccountRequest) (*pb.Empty, error) {
	s.guard.Lock()
	defer s.guard.Unlock()
	if _, found := s.accounts[in.GetName()]; found {
		delete(s.accounts, in.GetName())
		return &pb.Empty{}, nil
	}
	return &pb.Empty{}, fmt.Errorf("user %s not found", in.GetName())
}

func (s *server) RenameAccount(ctx context.Context, in *pb.RenameAccountRequest) (*pb.Empty, error) {
	s.guard.Lock()
	defer s.guard.Unlock()
	if val, found := s.accounts[in.GetName()]; found {
		_, other := s.accounts[in.GetNewName()]
		if other {
			return &pb.Empty{}, fmt.Errorf("user %s already exists", in.GetNewName())
		}
		s.accounts[in.GetNewName()] = val
		delete(s.accounts, in.GetName())
		return &pb.Empty{}, nil
	}
	return &pb.Empty{}, fmt.Errorf("user %s not found", in.GetName())
}

func (s *server) SetBalance(ctx context.Context, in *pb.SetBalanceRequest) (*pb.Empty, error) {
	s.guard.Lock()
	defer s.guard.Unlock()
	if _, found := s.accounts[in.GetName()]; found {
		s.accounts[in.GetName()] = in.GetAmount()
		return &pb.Empty{}, nil
	}
	return &pb.Empty{}, fmt.Errorf("user %s not found", in.GetName())
}

func main() {
	flag.Parse()
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAwesomeProjectServer(s, &server{accounts: make(map[string]int64), guard: &sync.RWMutex{}})
	log.Printf("server listening at %v", listener.Addr())
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
