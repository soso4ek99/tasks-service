package task

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type TaskRepository interface {
	Create(ctx context.Context, task *Task) (*Task, error)
	Get(ctx context.Context, id uint32) (*Task, error)
	List(ctx context.Context) ([]*Task, error)
	ListByUser(ctx context.Context, userID uint32) ([]*Task, error)
	Update(ctx context.Context, task *Task) (*Task, error)
	Delete(ctx context.Context, id uint32) error
}

// taskRepository реализует интерфейс TaskRepository.
type taskRepository struct {
	db *gorm.DB
}



// NewRepository создает новый TaskRepository.
func NewRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

// Get implements TaskRepository.
func (r *taskRepository) Get(ctx context.Context, id uint32) (*Task, error) {
	var task Task
	result := r.db.WithContext(ctx).First(&task, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("task not found")
		}
		return nil, fmt.Errorf("failed to get task: %w", result.Error)
	}
	return &task, nil
}

// ListByUser implements TaskRepository.
func (r *taskRepository) ListByUser(ctx context.Context, userID uint32) ([]*Task, error) {
	var tasks []*Task
	result := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&tasks)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to list tasks by user: %w", result.Error)
	}
	return tasks, nil
}


// Create implements TaskRepository.
func (r *taskRepository) Create(ctx context.Context, task *Task) (*Task, error) {
	result := r.db.WithContext(ctx).Create(task)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to create task: %w", result.Error)
	}
	return task, nil
}

// Delete implements TaskRepository.
func (r *taskRepository) Delete(ctx context.Context, id uint32) error {
	result := r.db.WithContext(ctx).Delete(&Task{}, id)
	if result.Error != nil {
		return fmt.Errorf("failed to delete task: %w", result.Error)
	}
	return nil
}

// List implements TaskRepository.
func (r *taskRepository) List(ctx context.Context) ([]*Task, error) {
	var tasks []*Task
	result := r.db.WithContext(ctx).Find(&tasks)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to list tasks: %w", result.Error)
	}
	return tasks, nil
}

// Update implements TaskRepository.
func (r *taskRepository) Update(ctx context.Context, task *Task) (*Task, error) {
	result := r.db.WithContext(ctx).Save(task)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to update task: %w", result.Error)
	}

	return task, nil
}
