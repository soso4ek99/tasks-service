package database

import (
	"fmt"
	"log"
	"github.com/soso4ek99/tasks-service/task"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
    dsn := "host=localhost user=postgres password=password dbname=user_db port=5432 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, fmt.Errorf("failed to connect to database: %w", err)
    }

    // Автомиграция для таблицы Task
    err = db.AutoMigrate(&task.Task{})
    if err != nil {
       log.Fatalf("failed to auto migrate: %v", err)
       return nil, fmt.Errorf("failed to auto migrate: %w", err) // Возвращаем ошибку!
    }

    return db, nil
}

