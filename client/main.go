package main

import (
	"context"
	"log"
	"time"

	pb "github.com/alexandreelkhoury/grpc-go/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Function definitions
func getLastBlock(client pb.BlockchainClient) (*pb.BlockInfo, error) {
	req := &pb.Empty{}
	res, err := client.GetLastBlock(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func createAndBroadcastBlock(client pb.BlockchainClient, uuid string) error {
	req := &pb.BakeRequest{
		Uuid: uuid,
	}
	_, err := client.BakeBlock(context.Background(), req)
	if err != nil {
		return err
	}
	log.Println("Block successfully created and broadcasted.")
	return nil
}
func main() {

	conn, err := grpc.NewClient("35.241.224.46:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewBlockchainClient(conn)
	log.Printf("Connected to server '35.241.224.46:50051'")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	res, err := client.Register(ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("could not register: %v", err)
	}
	log.Printf("Registered with UUID: %s, Reputation: %d", res.GetUuid(), res.GetReputation())

	// Get the last block to see pending transactions
	blockInfo, err := getLastBlock(client)
	if err != nil {
		log.Fatalf("could not get last block: %v", err)
	}
	log.Printf("Last Block Info: %v", blockInfo)

	// Example UUID for baking
	uuid := "example-uuid"

	// Create and broadcast a new block
	err = createAndBroadcastBlock(client, uuid)
	if err != nil {
		log.Fatalf("could not create and broadcast block: %v", err)
	}
}
