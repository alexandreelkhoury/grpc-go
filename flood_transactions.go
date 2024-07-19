// client/flood_transactions.go
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

	sender := "node1_uuid"
	receiver := "node2_uuid"
	amount := int32(50)
	data := "Flood transaction"

	floodNetworkWithTransactions(client, sender, receiver, amount, data, 100) // Flood with 100 transactions
}

func floodNetworkWithTransactions(client pb.BlockchainClient, sender string, receiver string, amount int32, data string, numTransactions int) {
	for i := 0; i < numTransactions; i++ {
		err := addTransaction(client, sender, receiver, amount, data)
		if err != nil {
			log.Printf("Error adding transaction: %v", err)
		}
	}
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
