package grpc

import (
	"context"

	taskpb "github.com/soso4ek99/project-protos/proto/task"
	userpb "github.com/soso4ek99/project-protos/proto/user"
	"github.com/soso4ek99/tasks-service/task"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)



type Handler struct {
	svc       	task.TaskService
	userClient userpb.UserServiceClient
	taskpb.UnimplementedTaskServiceServer
}

// NewHandler создает новый обработчик.
func NewHandler(svc task.TaskService, uc userpb.UserServiceClient) *Handler {
	return &Handler{svc: svc, userClient: uc}
}

// CreateTask создает новую задачу.
func (h *Handler) CreateTask(ctx context.Context, req *taskpb.CreateTaskRequest) (*taskpb.CreateTaskResponse, error) {
    // 1. Проверить пользователя:
    if _, err := h.userClient.GetUser(ctx, &userpb.GetUserRequest{Id: req.UserId}); err != nil { // Явное указание пакета
        return nil, status.Errorf(codes.NotFound, "user %d not found: %v", req.UserId, err)
    }

	// 2. Внутренняя логика:
	 t, err := h.svc.CreateTask(ctx, req.UserId, req.Title, req.Description)
    if err != nil {
        return nil, status.Errorf(codes.Internal, "failed to create task: %v", err)
    }

    // 3. Ответ:  Преобразуем model.Task в taskpb.Task
    taskPb := &taskpb.Task{
        Id:          t.ID,
        UserId:      t.UserID,
        Title:       t.Title,
        Description: t.Description,
        IsDone:      t.IsDone,
    }

    return &taskpb.CreateTaskResponse{Task: taskPb}, nil
}

// GetTask возвращает задачу по ID.
func (h *Handler) GetTask(ctx context.Context, req *taskpb.GetTaskRequest) (*taskpb.Task, error) {
	t, err := h.svc.GetTask(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "task not found: %v", err)
	}

	return &taskpb.Task{
		Id:          t.ID,
		UserId:      t.UserID,
		Title:       t.Title,
		Description: t.Description,
		IsDone:      t.IsDone,
	}, nil
}

// ListTasks возвращает список всех задач.
func (h *Handler) ListTasks(ctx context.Context, req *taskpb.ListTasksRequest) (*taskpb.ListTasksResponse, error) {
	tasks, err := h.svc.ListTasks(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list tasks: %v", err)
	}

	var taskPb []*taskpb.Task
	for _, t := range tasks {
		taskPb = append(taskPb, &taskpb.Task{
			Id:          t.ID,
			UserId:      t.UserID,
			Title:       t.Title,
			Description: t.Description,
			IsDone:      t.IsDone,
		})
	}

	return &taskpb.ListTasksResponse{Tasks: taskPb}, nil
}

// ListTasksByUser возвращает список задач для конкретного пользователя.
func (h *Handler) ListTasksByUser(ctx context.Context, req *taskpb.ListTasksByUserRequest) (*taskpb.ListTasksResponse, error) {
	tasks, err := h.svc.ListTasksByUser(ctx, req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list tasks by user: %v", err)
	}

	var taskPb []*taskpb.Task
	for _, t := range tasks {
		taskPb = append(taskPb, &taskpb.Task{
			Id:          t.ID,
			UserId:      t.UserID,
			Title:       t.Title,
			Description: t.Description,
			IsDone:      t.IsDone,
		})
	}

	return &taskpb.ListTasksResponse{Tasks: taskPb}, nil
}

// UpdateTask обновляет задачу.
func (h *Handler) UpdateTask(ctx context.Context, req *taskpb.UpdateTaskRequest) (*taskpb.UpdateTaskResponse, error) {
	// TODO: Добавить проверку пользователя, если это требуется

	taskToUpdate := &task.Task{
		ID:          req.Id,
		UserID:      0, // Заполни userID, если это необходимо и если у тебя есть такая логика
		Title:       req.Title,
		Description: req.Description,
		IsDone:      req.IsDone,
	}

	updatedTask, err := h.svc.UpdateTask(ctx, taskToUpdate)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update task: %v", err)
	}

	taskPb := &taskpb.Task{
		Id:          updatedTask.ID,
		Title:       updatedTask.Title,
		Description: updatedTask.Description,
		IsDone:      updatedTask.IsDone,
	}

	return &taskpb.UpdateTaskResponse{Task: taskPb}, nil
}

// DeleteTask удаляет задачу.
func (h *Handler) DeleteTask(ctx context.Context, req *taskpb.DeleteTaskRequest) (*taskpb.DeleteTaskResponse, error) {
	// TODO: Добавить проверку пользователя, если это требуется

	err := h.svc.DeleteTask(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete task: %v", err)
	}
	return &taskpb.DeleteTaskResponse{}, nil
}
