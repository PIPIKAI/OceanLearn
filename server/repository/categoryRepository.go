package repository

import (
	"server/common"
	"server/model"

	"github.com/jinzhu/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository() CategoryRepository {
	return CategoryRepository{DB: common.GetDB()}
}
func (c CategoryRepository) Create(name string) (*model.Category, error) {
	categpryModel := model.Category{
		Name: name,
	}
	if err := c.DB.Create(&categpryModel).Error; err != nil {
		return nil, err
	}
	return &categpryModel, nil
}
func (c CategoryRepository) Update(categoryModel model.Category, name string) (*model.Category, error) {

	if err := c.DB.Model(&categoryModel).Update("name", name).Error; err != nil {
		return nil, err
	}
	return &categoryModel, nil
}
func (c CategoryRepository) SelectById(id int) (*model.Category, error) {
	var category model.Category
	if err := c.DB.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}
func (c CategoryRepository) DeleteById(id int) error {
	if err := c.DB.Delete(model.Category{}, id).Error; err != nil {
		return err
	}
	return nil
}
