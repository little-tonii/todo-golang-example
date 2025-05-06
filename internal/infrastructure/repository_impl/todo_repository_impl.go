package repositoryimpl

// import (
// 	"errors"
// 	"fmt"
// 	"sync"
// 	"todo-golang-example/internal/domain/entity"
// 	"todo-golang-example/internal/domain/repository"
// 	"todo-golang-example/internal/infrastructure/config"
// 	"todo-golang-example/internal/infrastructure/model"

// 	"gorm.io/gorm"
// )

// type TodoRepositoryImpl struct {
// 	database *gorm.DB
// }

// var (
// 	todoRepository repository.TodoRepository
// 	once           sync.Once
// )

// func NewTodoRepositoryImpl(database *gorm.DB) *TodoRepositoryImpl {
// 	return &TodoRepositoryImpl{
// 		database: database,
// 	}
// }

// func GetTodoRepository() repository.TodoRepository {
// 	once.Do(func() {
// 		todoRepository = NewTodoRepositoryImpl(config.GetDatabase())
// 	})
// 	return todoRepository
// }

// func (todoRepositoryImpl *TodoRepositoryImpl) Create(todoEntity *entity.TodoEntity) error {
// 	todoModel := model.TodoModel{
// 		Description: todoEntity.Description,
// 		Title:       todoEntity.Title,
// 	}
// 	result := todoRepositoryImpl.database.Create(&todoModel)
// 	if result.Error != nil {
// 		return errors.New("Có lỗi xảy ra khi tạo mới todo")
// 	}
// 	return nil
// }

// func (todoRepositoryImpl *TodoRepositoryImpl) Update(todoEntity *entity.TodoEntity) error {
// 	todoModel := model.TodoModel{
// 		Id:          todoEntity.Id,
// 		Description: todoEntity.Description,
// 		Title:       todoEntity.Title,
// 	}
// 	result := todoRepositoryImpl.database.Save(&todoModel)
// 	if result.Error != nil {
// 		return errors.New("Có lỗi xảy ra khi cập nhật todo")
// 	}
// 	return nil
// }

// func (todoRepositoryImpl *TodoRepositoryImpl) DeleteById(todoEntity *entity.TodoEntity) error {
// 	todoModel := model.TodoModel{
// 		Id: todoEntity.Id,
// 	}
// 	result := todoRepositoryImpl.database.Delete(&todoModel)
// 	if result.Error != nil {
// 		return errors.New("Có lỗi xảy ra khi xóa todo")
// 	}
// 	return nil
// }

// func (todoRepositoryImpl *TodoRepositoryImpl) GetById(id int64) (*entity.TodoEntity, error) {
// 	todoModel := model.TodoModel{Id: id}
// 	result := todoRepositoryImpl.database.First(&todoModel)
// 	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
// 		return nil, errors.New("Todo không tồn tại")
// 	}
// 	if result.Error != nil {
// 		return nil, fmt.Errorf("Có lỗi xảy ra khi lấy todo: %w", result.Error)
// 	}
// 	return todoModel.ToEntity(), nil
// }

// func (todoRepositoryImpl *TodoRepositoryImpl) GetAll(page int, size int) ([]*entity.TodoEntity, error) {
// 	offset := (page - 1) * size
// 	todoModels := make([]model.TodoModel, 0)
// 	result := todoRepositoryImpl.database.
// 		Offset(offset).
// 		Limit(size).
// 		Find(&todoModels)
// 	if result.Error != nil {
// 		return nil, errors.New("Có lỗi xảy ra khi lấy danh sách todo")
// 	}
// 	todoEntities := make([]*entity.TodoEntity, len(todoModels))
// 	for i, todoModel := range todoModels {
// 		todoEntities[i] = todoModel.ToEntity()
// 	}
// 	return todoEntities, nil
// }
