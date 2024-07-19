// transaction.go
package main

import (
	"context"
	"log"
	"time"

	pb "github.com/alexandreelkhoury/grpc-go/proto" // Importez le package généré par protoc
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Connexion au serveur gRPC
	conn, err := grpc.Dial("35.241.224.46:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Création du client gRPC
	client := pb.NewBlockchainClient(conn)
	log.Printf("Connected to server '35.241.224.46:50051'")

	// Création de la transaction avec des valeurs codées en dur
	tx := &pb.Transaction{
		Sender:   "9117360652631807282",
		Receiver: "8940082882196709507",
		Amount:   100, // montant en unités
		Data:     "FERDIOUI",
	}

	// Création du contexte avec un délai d'expiration
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Appel de la méthode AddTransaction
	_, err = client.AddTransaction(ctx, tx)
	if err != nil {
		log.Fatalf("could not add transaction: %v", err)
	}
	log.Println("Transaction added successfully")
}
