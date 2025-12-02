package main

import (
	"context"
	"log"
	"time"

	"github.com/fatih/color"
	"github.com/lva100/go-grpc/pkg/note_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:50051"
	noteID  = 12
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Failed to connect to server:", err)
	}
	defer conn.Close()

	c := note_v1.NewNoteV1Client(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Get(ctx, &note_v1.GetRequest{Id: noteID})
	if err != nil {
		log.Fatal("Failed to get note by id:", err)
	}
	log.Printf(color.RedString("Note info:\n"), color.GreenString("%+v", r.GetNote()))
}
