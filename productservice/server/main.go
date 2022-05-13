package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/evgeniy-dammer/httpbasicauthgrpc/productservice/handlers"
	"github.com/evgeniy-dammer/httpbasicauthgrpc/productservice/interceptors"
	productservice "github.com/evgeniy-dammer/httpbasicauthgrpc/productservice/proto"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
)

func main() {
	creds, err := credentials.NewServerTLSFromFile("../../keys/server-cert.pem", "../../keys/server-key.pem")

	if err != nil {
		log.Fatalf("failed to setup tls: %v", err)
	}

	listen, err := net.Listen("tcp", ":1111")

	if err != nil {
		fmt.Println(err)
	}

	defer listen.Close()

	productServ := handlers.ProductServiceServer{}

	grpcServer := grpc.NewServer(
		grpc.Creds(creds),
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				interceptors.BasicAuthInterceptor,
			),
		),
	)

	productservice.RegisterProductServiceServer(grpcServer, &productServ)

	if err := grpcServer.Serve(listen); err != nil {
		fmt.Println(err)
	}
}
