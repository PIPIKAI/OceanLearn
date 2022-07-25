package controller

import (
	"server/model"
	"server/repository"
	"server/util"
	"server/vo"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	repository repository.CategoryRepository
}

func NewCategoryController() CategoryController {
	repository := repository.NewCategoryRepository()
	repository.DB.AutoMigrate(model.Category{})
	return CategoryController{repository}
}

type ICategoryController interface {
	RestController
}

func (c CategoryController) Create(ctx *gin.Context) {
	var requestCategory vo.CreateCategoryRequest

	if ctx.ShouldBind(&requestCategory) != nil {
		util.Response{}.Error(ctx, nil, "数据验证错误，分类名必填")
		return
	}
	// var category *model.Category
	category, err := c.repository.Create(requestCategory.Name)
	if err != nil {
		util.Response{}.Error(ctx, nil, "创建失败，已存在该类")
		return
	}
	util.Response{}.Success(ctx, gin.H{"category": category}, "")
}

func (c CategoryController) Update(ctx *gin.Context) {
	// 绑定body中的参数
	var requestCategory vo.CreateCategoryRequest
	if ctx.ShouldBind(&requestCategory) != nil {
		util.Response{}.Error(ctx, nil, "数据验证错误，分类名必填")
		return
	}
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	selectedCategory, err := c.repository.SelectById(categoryId)
	if err != nil {
		util.Response{}.Error(ctx, nil, "分类不存在")
		panic(err)
	}
	// 更新
	updateCategory, err := c.repository.Update(*selectedCategory, requestCategory.Name)
	if err != nil {
		util.Response{}.Error(ctx, nil, "更新失败，已存在该类")
		return
	}
	util.Response{}.Success(ctx, gin.H{"category": updateCategory}, "修改成功")

}

func (c CategoryController) Show(ctx *gin.Context) {
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	selectedCategory, err := c.repository.SelectById(categoryId)
	if err != nil {
		util.Response{}.Error(ctx, nil, "分类不存在")
		return
	}

	util.Response{}.Success(ctx, gin.H{"category": selectedCategory}, "修改成功")
}

func (c CategoryController) Delete(ctx *gin.Context) {
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	_, err := c.repository.SelectById(categoryId)
	if err != nil {
		panic(err)

	}
	if c.repository.DeleteById(categoryId) != nil {
		util.Response{}.Error(ctx, nil, "删除失败")
		return
	}
	util.Response{}.Success(ctx, nil, "删除成功")

}
