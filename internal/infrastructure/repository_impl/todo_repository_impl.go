package repositoryimpl

import (
	"todo-golang-example/internal/domain/entity"
	"todo-golang-example/internal/infrastructure/model"

	"gorm.io/gorm"
)

type TodoRepositoryImpl struct {
	database *gorm.DB
}

func NewTodoRepositoryImpl(database *gorm.DB) *TodoRepositoryImpl {
	return &TodoRepositoryImpl{
		database: database,
	}
}

func (repository *TodoRepositoryImpl) Create(todoEntity *entity.TodoEntity) error {
	todoModel := model.TodoModel{
		Title:       todoEntity.Title,
		Description: todoEntity.Description,
		UserId:      todoEntity.UserId,
	}
	error := repository.database.Create(&todoModel).
		Error
	if error != nil {
		return error
	}
	todoEntity.Id = todoModel.Id
	todoEntity.CreatedAt = todoModel.CreatedAt
	todoEntity.UpdatedAt = todoModel.UpdatedAt
	return nil
}

func (repository *TodoRepositoryImpl) Update(todoEntity *entity.TodoEntity) error {
	todoModel := model.TodoModel{
		Title:       todoEntity.Title,
		Description: todoEntity.Description,
	}
	result := repository.database.
		Model(&model.TodoModel{}).
		Where("id = ?", todoEntity.Id).
		Updates(&todoModel)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	var updatedTodo model.TodoModel
	error := repository.database.
		Where("id = ?", todoEntity.Id).
		First(&updatedTodo).Error
	if error != nil {
		return error
	}
	*todoEntity = *updatedTodo.ToEntity()
	return nil
}

func (repository *TodoRepositoryImpl) DeleteById(id int64) error {
	result := repository.database.Delete(&model.TodoModel{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (repository *TodoRepositoryImpl) FindById(id int64) (*entity.TodoEntity, error) {
	var todoModel model.TodoModel
	err := repository.database.
		Where("id = ?", id).
		First(&todoModel).
		Error
	if err != nil {
		return nil, err
	}
	return todoModel.ToEntity(), nil
}

func (repository *TodoRepositoryImpl) List(userId int64, page int64, size int64) ([]*entity.TodoEntity, error) {
	var todoModels []model.TodoModel
	error := repository.database.
		Where("user_id = ?", userId).
		Limit(int(size)).
		Offset(int((page - 1) * size)).
		Find(&todoModels).Error
	if error != nil {
		return nil, error
	}
	todoEntities := make([]*entity.TodoEntity, len(todoModels))
	for index, todoModel := range todoModels {
		todoEntities[index] = todoModel.ToEntity()
	}
	return todoEntities, nil
}
