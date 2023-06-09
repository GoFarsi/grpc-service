package main

import (
	"context"
	"fmt"
	"github.com/GoFarsi/grpc-service/zarf/gen/greeting/v1"
	"google.golang.org/grpc"
	"log"
)

func main() {
	dialOpts := make([]grpc.DialOption, 0)
	dialOpts = append(dialOpts, grpc.WithInsecure())

	conn, err := grpc.Dial(":8080", dialOpts...)
	if err != nil {
		log.Fatalln(err)
	}

	g := greeting.NewGreetingServiceClient(conn)
	u := greeting.NewUserServiceClient(conn)

	resp, err := g.Greeting(context.Background(), &greeting.GreetingRequest{
		Name: "Javad",
	})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(resp.Message)

	respUser, err := u.CreateUser(context.Background(), &greeting.CreateUserRequest{
		FirstName: "Javad",
		LastName:  "Rajabzadeh",
		Age:       29,
	})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(respUser)
}
