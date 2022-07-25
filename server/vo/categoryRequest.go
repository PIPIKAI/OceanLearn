package vo

type CreateCategoryRequest struct {
	Name string `json:"name" gorm:"type:varchar(50); not null;unique"`
}
