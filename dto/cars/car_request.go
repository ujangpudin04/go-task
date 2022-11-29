package carsdto

type CarCreateRequest struct {
	Price string `json:"price" form:"price" validate:"required" `
	Brand string `json:"brand" form:"brand" validate:"required"`
	Type  string `json:"type" form:"type" validate:"required"`
}

type CarUpdateRequest struct {
	Price string `json:"price" form:"price" gorm:"type: varchar(255)"`
	Brand string `json:"brand" gorm:"type:text" form:"brand"`
	Type  string `json:"type" form:"type" gorm:"type: int"`
}
