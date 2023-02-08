package request

type AddRole struct {
	Name string `json:"name" validate:"required"`
}
