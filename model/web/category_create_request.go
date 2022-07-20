package web

type CategoryCreateRequest struct {
	Name string `validate:"required" json:"name"`
}
