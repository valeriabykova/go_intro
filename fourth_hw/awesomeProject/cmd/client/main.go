package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	pb "awesomeProject/api/proto"

	"github.com/labstack/gommon/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	addr := flag.String("addr", "localhost:1323", "the address to connect to")
	cmd := flag.String("cmd", "", "command to execute")
	nameVal := flag.String("name", "", "name of account")
	newNameVal := flag.String("new_name", "", "new name of account")
	amountVal := flag.Int64("amount", 0, "amount of account")
	flag.Parse()

	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect %v", err)
		return
	}
	defer conn.Close()

	client := pb.NewAwesomeProjectClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	switch *cmd {
	case "create":
		_, err = client.CreateAccount(ctx, &pb.CreateAccountRequest{Name: nameVal, Amount: amountVal})
	case "get":
		var resp *pb.GetAccountResponse
		resp, err = client.GetAccount(ctx, &pb.GetAccountRequest{Name: nameVal})
		if err == nil {
			fmt.Printf("Got account \"%s\" with balance %d\n", resp.GetName(), resp.GetAmount())
		}
	case "update":
		_, err = client.SetBalance(ctx, &pb.SetBalanceRequest{Name: nameVal, Amount: amountVal})
	case "rename":
		_, err = client.RenameAccount(ctx, &pb.RenameAccountRequest{Name: nameVal, NewName: newNameVal})
	case "delete":
		_, err = client.DeleteAccount(ctx, &pb.DeleteAccountRequest{Name: nameVal})
	default:
		err = fmt.Errorf("unknown command %s", *cmd)
	}
	if err != nil {
		log.Fatalf("Error doing request: %v", err)
		return
	}
}
