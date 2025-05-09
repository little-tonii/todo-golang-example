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
	return nil
}

func (repository *TodoRepositoryImpl) Update(todoEntity *entity.TodoEntity) error {
	result := repository.database.
		Model(&model.TodoModel{}).
		Where("id = ?", todoEntity.Id).
		Updates(map[string]any{
			"title":       todoEntity.Title,
			"description": todoEntity.Description,
		})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
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

func (repository *TodoRepositoryImpl) List(page int64, size int64) ([]*entity.TodoEntity, error) {
	var todoModels []model.TodoModel
	error := repository.database.
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
