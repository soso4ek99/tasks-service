package task

import (
	"context"
	"fmt"

)
type TaskService interface {
CreateTask(ctx context.Context, userId uint32, title string, description string) (*Task, error) // Добавлено description
	GetTask(ctx context.Context, id uint32) (*Task, error)
	ListTasks(ctx context.Context) ([]*Task, error)
	ListTasksByUser(ctx context.Context, userID uint32) ([]*Task, error)
	UpdateTask(ctx context.Context, task *Task) (*Task, error)
	DeleteTask(ctx context.Context, id uint32) error
}

// taskServiceImpl implements the TaskService interface.
type taskServiceImpl struct {
	repo TaskRepository
}

// NewTaskService creates a new TaskService.
func NewTaskServiceImpl(repo TaskRepository) TaskService { // Принимаем TaskRepository
	return &taskServiceImpl{repo: repo}
}

// CreateTask создает новую задачу.
func (s *taskServiceImpl) CreateTask(ctx context.Context, userId uint32, title string, description string) (*Task, error) {  // Добавлено description
	newTask := &Task{
		UserID:      userId,
		Title:       title,
		Description: description,
		IsDone:      false, // Новые задачи по умолчанию не выполнены
	}
	createdTask, err := s.repo.Create(ctx, newTask)
	if err != nil {
		return nil, fmt.Errorf("failed to create task: %w", err)
	}
	return createdTask, nil
}

// GetTask возвращает задачу по ID.
func (s *taskServiceImpl) GetTask(ctx context.Context, id uint32) (*Task, error) {
	retrievedTask, err := s.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return retrievedTask, nil
}

// ListTasks возвращает список всех задач.
func (s *taskServiceImpl) ListTasks(ctx context.Context) ([]*Task, error) {
	tasks, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

// ListTasksByUser возвращает список задач для конкретного пользователя.
func (s *taskServiceImpl) ListTasksByUser(ctx context.Context, userID uint32) ([]*Task, error) {
	tasks, err := s.repo.ListByUser(ctx, userID)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

// UpdateTask обновляет задачу.
func (s *taskServiceImpl) UpdateTask(ctx context.Context, taskToUpdate *Task) (*Task, error) {
	// 1. Получи userId из контекста или запроса
	// !!! ВАЖНО: замени этот код на фактический код получения userId из запроса/контекста
	userId := taskToUpdate.UserID // Пример: если userId уже есть в taskToUpdate

	// 2. Установи userId в taskToUpdate
	taskToUpdate.UserID = userId

	updatedTask, err := s.repo.Update(ctx, taskToUpdate)
	if err != nil {
		return nil, err
	}
	return updatedTask, nil
}

// DeleteTask удаляет задачу.
func (s *taskServiceImpl) DeleteTask(ctx context.Context, id uint32) error {
	err := s.repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

