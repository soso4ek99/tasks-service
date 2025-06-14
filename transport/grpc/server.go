package grpc

import (
	"log"
	"net"

	taskpb "github.com/soso4ek99/project-protos/proto/task"
	userpb "github.com/soso4ek99/project-protos/proto/user"
	"github.com/soso4ek99/tasks-service/task"
	"google.golang.org/grpc"
)
func RunGRPCServer(lis net.Listener, svc task.TaskService, uc userpb.UserServiceClient) error {
	grpcSrv := grpc.NewServer()
	handler := NewHandler(svc, uc) // Используем конструктор NewHandler
	taskpb.RegisterTaskServiceServer(grpcSrv, handler)

	log.Printf("gRPC server listening on %s", lis.Addr()) // Используем lis.Addr()

	if err := grpcSrv.Serve(lis); err != nil {
		log.Printf("failed to serve: %v", err) // Используем Logf, чтобы не завершать программу
		return err                               // Возвращаем ошибку
	}

	return nil
}