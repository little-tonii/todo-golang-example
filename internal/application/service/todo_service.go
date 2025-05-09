package service

import (
	"errors"
	"fmt"
	"net/http"
	"todo-golang-example/internal/application/request"
	"todo-golang-example/internal/application/response"
	"todo-golang-example/internal/domain/entity"
	"todo-golang-example/internal/domain/repository"
	"todo-golang-example/internal/shared/common"

	"gorm.io/gorm"
)

type TodoService struct {
	todoRepository repository.TodoRepository
}

func NewTodoService(todoRepostiroy repository.TodoRepository) *TodoService {
	return &TodoService{
		todoRepository: todoRepostiroy,
	}
}

func (service *TodoService) GetTodoById(id int64) (*response.GetTodoByIdResponse, *common.ApplicationError) {
	todoEntity, error := service.todoRepository.FindById(id)
	if error != nil {
		if errors.Is(error, gorm.ErrRecordNotFound) {
			return nil, common.NewApplicationError(
				http.StatusNotFound,
				errors.New(fmt.Sprintf("Todo không tồn tại")),
			)
		} else {
			return nil, common.NewApplicationError(http.StatusInternalServerError, error)
		}
	}
	return &response.GetTodoByIdResponse{
		Id:          todoEntity.Id,
		Title:       todoEntity.Title,
		Description: todoEntity.Description,
		CreatedAt:   todoEntity.CreatedAt,
		UpdatedAt:   todoEntity.UpdatedAt,
	}, nil
}

func (service *TodoService) CreateTodo(userId int64, request *request.CreateTodoRequest) (*response.CreateTodoResponse, *common.ApplicationError) {
	todoEntity := &entity.TodoEntity{
		Title:       request.Title,
		Description: request.Description,
		UserId:      userId,
	}
	error := service.todoRepository.Create(todoEntity)
	if error != nil {
		return nil, common.NewApplicationError(http.StatusInternalServerError, error)
	}
	return &response.CreateTodoResponse{
		Id:          todoEntity.Id,
		Title:       todoEntity.Title,
		Description: todoEntity.Description,
		CreatedAt:   todoEntity.CreatedAt,
		UpdatedAt:   todoEntity.UpdatedAt,
	}, nil
}

func (service *TodoService) DeleteTodoById(id int64) *common.ApplicationError {
	error := service.todoRepository.DeleteById(id)
	if error != nil {
		if errors.Is(error, gorm.ErrRecordNotFound) {
			return common.NewApplicationError(
				http.StatusNotFound,
				errors.New(fmt.Sprintf("Todo không tồn tại")),
			)
		} else {
			return common.NewApplicationError(http.StatusInternalServerError, error)
		}
	}
	return nil
}

func (service *TodoService) GetTodoList(userId int64, request *request.GetTodoListRequest) (*response.GetTodoListResponse, *common.ApplicationError) {
	todoEntites, error := service.todoRepository.List(userId, request.Page, request.Size)
	if error != nil {
		return nil, common.NewApplicationError(http.StatusInternalServerError, error)
	}
	todoResposes := make([]*response.GetTodoByIdResponse, len(todoEntites))
	for index, todoEntity := range todoEntites {
		todoResposes[index] = &response.GetTodoByIdResponse{
			Id:          todoEntity.Id,
			Title:       todoEntity.Title,
			Description: todoEntity.Description,
			CreatedAt:   todoEntity.CreatedAt,
			UpdatedAt:   todoEntity.UpdatedAt,
		}
	}
	return &response.GetTodoListResponse{
		Page:  request.Page,
		Size:  request.Size,
		Todos: todoResposes,
	}, nil
}

func (service *TodoService) UpdateTodoById(todoId int64, request *request.UpdateTodoByIdRequest) (*response.UpdateTodoByIdResponse, *common.ApplicationError) {
	todoEntity := &entity.TodoEntity{
		Id:          todoId,
		Title:       request.Title,
		Description: request.Description,
	}
	error := service.todoRepository.Update(todoEntity)
	if error != nil {
		if errors.Is(error, gorm.ErrRecordNotFound) {
			return nil, common.NewApplicationError(
				http.StatusNotFound,
				errors.New(fmt.Sprintf("Todo không tồn tại")),
			)
		} else {
			return nil, common.NewApplicationError(http.StatusInternalServerError, error)
		}
	}
	return &response.UpdateTodoByIdResponse{
		Id:          todoEntity.Id,
		Title:       todoEntity.Title,
		Description: todoEntity.Description,
		CreatedAt:   todoEntity.CreatedAt,
		UpdatedAt:   todoEntity.UpdatedAt,
	}, nil
}
