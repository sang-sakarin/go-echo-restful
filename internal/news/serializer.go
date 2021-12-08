package news

type SerializerForm struct {
	Title       string `json:"title" form:"title" validate:"required"`
	Description string `json:"description" form:"description" validate:"required"`
	IsPublish   bool   `json:"is_publish" form:"is_publish" validate:"required"`
}
