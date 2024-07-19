// client/double_spending.go
package main

import (
	"context"
	"log"

	pb "github.com/alexandreelkhoury/grpc-go/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("35.241.224.46:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewBlockchainClient(conn)

	err = performDoubleSpendingAttack(client)
	if err != nil {
		log.Fatalf("Error performing double spending attack: %v", err)
	}
}

func performDoubleSpendingAttack(client pb.BlockchainClient) error {
	// Create two conflicting transactions
	err := addTransaction(client, "8940082882196709507", "8940082882196709507", 50, "First transaction")
	if err != nil {
		return err
	}

	err = addTransaction(client, "8940082882196709507", "8940082882196709507", 50, "Second conflicting transaction")
	if err != nil {
		return err
	}

	return nil
}

func addTransaction(client pb.BlockchainClient, sender string, receiver string, amount int32, data string) error {
	tx := &pb.Transaction{
		Sender:   sender,
		Receiver: receiver,
		Amount:   amount,
		Data:     data,
	}
	_, err := client.AddTransaction(context.Background(), tx)
	return err
}
