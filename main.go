package main

import (
	"context"
	_ "embed"
	"flag"
	"fmt"
	"github.com/GoFarsi/grpc-service/zarf/gen/greeting/v1"
	"github.com/Ja7ad/swaggerui"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"math/rand"
	"net"
	"net/http"
)

//go:embed zarf/gen/api/apidocs.swagger.json
var Swagger []byte

type GreetingService struct {
	greeting.UnimplementedGreetingServiceServer
}

type UserService struct {
	greeting.UnimplementedUserServiceServer
}

func (g *GreetingService) Greeting(ctx context.Context, in *greeting.GreetingRequest) (*greeting.GreetingResponse, error) {
	return &greeting.GreetingResponse{
		Message: fmt.Sprintf("greeting %s", in.Name),
	}, nil
}

func (u *UserService) CreateUser(ctx context.Context, in *greeting.CreateUserRequest) (*greeting.User, error) {
	return &greeting.User{
		Id:        rand.Int31(),
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Age:       in.Age,
	}, nil
}

func main() {
	development := flag.Bool("development", false, "dev mode")
	address := flag.String("address", ":8080", "server address")
	flag.Parse()

	listener, err := net.Listen("tcp", *address)
	if err != nil {
		log.Fatalln(err)
	}

	srv := grpc.NewServer()

	if *development {
		reflection.Register(srv)
	}

	greeting.RegisterGreetingServiceServer(srv, &GreetingService{})
	greeting.RegisterUserServiceServer(srv, &UserService{})

	mux := http.NewServeMux()
	mux.Handle("/api-docs/", http.StripPrefix("/api-docs", swaggerui.Handler(Swagger)))

	httpServer := &http.Server{
		Handler: mux,
		Addr:    ":8081",
	}

	go func() {
		log.Println("http server on :8081 has been started")
		log.Fatalln(httpServer.ListenAndServe())
	}()

	log.Printf("grpc server on %s has been started", *address)
	log.Fatalln(srv.Serve(listener))
}
