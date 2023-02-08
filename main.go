package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/GeeCeeCee/grpc-proto/person"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := person.NewSayHelloServiceClient(conn)

	response, err := c.SayHello(context.Background(), &person.Message{Body: "Hello From Client!"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Body)

}
