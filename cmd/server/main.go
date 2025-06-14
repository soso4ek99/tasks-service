package main

import (
	"log"
	"net"

	"github.com/soso4ek99/tasks-service/internal/database"
	"github.com/soso4ek99/tasks-service/task"
	"github.com/soso4ek99/tasks-service/transport/grpc"
)
func main() {
	// 1. Инициализация БД
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			log.Printf("failed to get sql db: %v", err)
			return
		}
		if err := sqlDB.Close(); err != nil {
			log.Printf("failed to close db: %v", err)
		}
	}()

	// 2. Репозиторий и сервис задач
	repo := task.NewRepository(db)
	svc := task.NewTaskServiceImpl(repo)

	// 3. Клиент к Users-сервису
	userClient, conn, err := grpc.NewUserClient("localhost:50051") // Замени на правильный импорт
	if err != nil {
		log.Fatalf("failed to connect to users: %v", err)
	}
	defer conn.Close()

	// 4. Создание Listener'а
	lis, err := net.Listen("tcp", ":50052") // Слушаем на порту 50052
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Task service listening on %s", lis.Addr())

	// 5. Запуск gRPC Tasks-сервиса
	if err := grpc.RunGRPCServer(lis, svc, userClient); err != nil { // Передаем listener в RunGRPCServer
		log.Fatalf("Tasks gRPC server error: %v", err)
	}
}