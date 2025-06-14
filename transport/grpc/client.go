package grpc

import (
	"fmt"

	userpb "github.com/soso4ek99/project-protos/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


func NewUserClient(addr string) (userpb.UserServiceClient, *grpc.ClientConn, error) {
	// 1. grpc.Dial(addr, grpc.WithInsecure())
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials())) // Use insecure.NewCredentials() // TODO: Use TLS
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to Users service: %w", err)
	}

	// 2. userpb.NewUserServiceClient(conn)
	client := userpb.NewUserServiceClient(conn)

	// 3. вернуть client, conn, err
	return client, conn, nil
}