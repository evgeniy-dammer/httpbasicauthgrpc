package main

import (
	"context"
	"fmt"
	"log"

	productservice "github.com/evgeniy-dammer/httpbasicauthgrpc/productservice/proto"
	"github.com/evgeniy-dammer/httpbasicauthgrpc/security"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	creds, err := credentials.NewClientTLSFromFile("../../keys/server-cert.pem", "../../keys/server-key.pem")

	if err != nil {
		log.Fatalf("Failed to staup tls: %v", err)
	}

	connection, err := grpc.Dial(
		"localhost:1111",
		grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(security.BasicAuth{
			Username: "acc1",
			Password: "123",
		}),
	)

	if err != nil {
		fmt.Println(err)
	}

	defer connection.Close()

	productServ := productservice.NewProductServiceClient(connection)

	response1, err1 := productServ.FindAll(context.Background(), &productservice.FindAllRequest{})

	if err1 != nil {
		fmt.Println(err1)
	} else {
		products := response1.Products

		fmt.Println("Product List")

		for _, product := range products {
			fmt.Println("Id: ", product.Id)
			fmt.Println("Name: ", product.Name)
			fmt.Println("Price: ", product.Price)
			fmt.Println("Quantity: ", product.Quantity)
			fmt.Println("Status: ", product.Status)
			fmt.Println("========================")
		}
	}

	response2, err2 := productServ.Search(context.Background(), &productservice.SearchRequest{Keyword: "vi"})

	if err2 != nil {
		fmt.Println(err2)
	} else {
		products := response2.Products

		fmt.Println("Search Result")

		for _, product := range products {
			fmt.Println("Id: ", product.Id)
			fmt.Println("Name: ", product.Name)
			fmt.Println("Price: ", product.Price)
			fmt.Println("Quantity: ", product.Quantity)
			fmt.Println("Status: ", product.Status)
			fmt.Println("========================")
		}
	}
}
