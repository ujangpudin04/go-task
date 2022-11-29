package carsdto

type CarResponse struct {
	ID    int    `json:"id"`
	Price string `json:"price" form:"price" validate:"required"`
	Brand string `json:"brand" form:"brand" validate:"required"`
	Type  string `json:"type" form:"type"`
}
