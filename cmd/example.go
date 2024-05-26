package main

import (
	"context"
	gen "github.com/breathbath/chat-protos/gen/protos/v1"
	"github.com/breathbath/chat-protos/service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// start the gRPC server
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	gen.RegisterChatServiceServer(grpcServer, &service.ExampleService{})
	go grpcServer.Serve(lis)

	conn, err := grpc.NewClient(":9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	// create an HTTP router using the client connection above
	// and register it with the service client
	rmux := runtime.NewServeMux()
	client := gen.NewChatServiceClient(conn)
	err = gen.RegisterChatServiceHandlerClient(ctx, rmux, client)
	if err != nil {
		log.Fatal(err)
	}

	// create a standard HTTP router
	mux := http.NewServeMux()

	// mount the gRPC HTTP gateway to the root
	mux.Handle("/", rmux)

	// mount a path to expose the generated OpenAPI specification on disk
	mux.HandleFunc("/swagger-ui/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./gen/docs/protos/v1/chat.swagger.json")
	})

	// mount the Swagger UI that uses the OpenAPI specification path above
	mux.Handle("/swagger-ui/", http.StripPrefix("/swagger-ui/", http.FileServer(http.Dir("./swagger-ui"))))

	// start a standard HTTP server with the router
	err = http.ListenAndServe(":8282", mux)
	if err != nil {
		log.Fatal(err)
	}
}
