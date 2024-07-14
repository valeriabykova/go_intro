package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net"

	_ "github.com/jackc/pgx/v5/stdlib"

	pb "awesomeProject/api/proto"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

var (
	port = flag.Int("port", 1323, "The server port")
)

type server struct {
	pb.AwesomeProjectServer
	db *sql.DB
}

func (s *server) get(user string) (int64, bool, error) {
	rows, err := s.db.Query(fmt.Sprintf("SELECT amount FROM accounts WHERE name = '%s'", user))
	if err != nil {
		return 0, false, err
	}
	defer rows.Close()

	var amount int
	found := false
	for rows.Next() {
		found = true
		err := rows.Scan(&amount)
		if err != nil {
			return 0, false, fmt.Errorf("unable to scan row: %w", err)
		}
	}
	if !found {
		return 0, found, nil
	}
	return int64(amount), true, nil
}

func (s *server) insert(user string, amount int64) error {
	query := fmt.Sprintf("INSERT INTO accounts (name, amount) VALUES ('%s', %d)", user, amount)
	_, err := s.db.Exec(query)
	if err != nil {
		return fmt.Errorf("unable to insert: %w", err)
	}

	return nil
}

func (s *server) update(user string, amount int64) error {
	query := fmt.Sprintf("UPDATE accounts SET amount = %d WHERE name = '%s'", amount, user)
	_, err := s.db.Exec(query)
	if err != nil {
		return fmt.Errorf("unable to update: %w", err)
	}
	return nil
}

func (s *server) delete(user string) error {
	query := fmt.Sprintf("DELETE FROM accounts WHERE name = '%s'", user)
	_, err := s.db.Exec(query)
	if err != nil {
		return fmt.Errorf("unable to delete: %w", err)
	}

	return nil
}

func (s *server) CreateAccount(ctx context.Context, in *pb.CreateAccountRequest) (*pb.Empty, error) {
	_, found, err := s.get(in.GetName())
	if err != nil {
		return nil, err
	}
	if found {
		return &pb.Empty{}, fmt.Errorf("user %s already exists", in.GetName())
	}
	err = s.insert(in.GetName(), in.GetAmount())
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}

func (s *server) GetAccount(ctx context.Context, in *pb.GetAccountRequest) (*pb.GetAccountResponse, error) {
	amount, found, err := s.get(in.GetName())
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, fmt.Errorf("account \"%s\" not found", in.GetName())
	}
	return &pb.GetAccountResponse{
		Name:   proto.String(in.GetName()),
		Amount: proto.Int64(amount),
	}, nil
}

func (s *server) DeleteAccount(ctx context.Context, in *pb.DeleteAccountRequest) (*pb.Empty, error) {
	_, found, err := s.get(in.GetName())
	if err != nil {
		return nil, err
	}
	if !found {
		return &pb.Empty{}, fmt.Errorf("user %s not found", in.GetName())
	}
	err = s.delete(in.GetName())
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}

func (s *server) RenameAccount(ctx context.Context, in *pb.RenameAccountRequest) (*pb.Empty, error) {
	amount, found, err := s.get(in.GetName())
	if err != nil {
		return nil, err
	}
	if !found {
		return &pb.Empty{}, fmt.Errorf("user %s not found", in.GetName())
	}
	_, found, err = s.get(in.GetNewName())
	if err != nil {
		return nil, err
	}
	if found {
		return &pb.Empty{}, fmt.Errorf("user %s already exists", in.GetNewName())
	}
	err = s.insert(in.GetNewName(), amount)
	if err != nil {
		return nil, err
	}
	err = s.delete(in.GetName())
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}

func (s *server) SetBalance(ctx context.Context, in *pb.SetBalanceRequest) (*pb.Empty, error) {
	_, found, err := s.get(in.GetName())
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, fmt.Errorf("user %s not found", in.GetName())
	}
	err = s.update(in.GetName(), in.GetAmount())
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}

func main() {
	psqlconn := "host=0.0.0.0 port=5432 user=postgres password=mypassword dbname=postgres sslmode=disable"

	db, err := sql.Open("pgx", psqlconn)
	if err != nil {
		log.Fatal("error connecting to db", err)
	}
	defer db.Close()

	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS ACCOUNTS(
			NAME	VARCHAR(20) NOT NULL UNIQUE,
			AMOUNT	INT NOT NULL
	);`)

	if err != nil {
		log.Fatal("Error creating accounts db", err)
	}

	flag.Parse()
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAwesomeProjectServer(s, &server{db: db})
	log.Printf("server listening at %v", listener.Addr())
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
