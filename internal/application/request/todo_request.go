package request

type CreateTodoRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type GetTodoListRequest struct {
	Page int64 `form:"page" binding:"gt=0"`
	Size int64 `form:"size" binding:"gt=0,lte=50"`
}

type UpdateTodoByIdRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}
