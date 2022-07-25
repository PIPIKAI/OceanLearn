package controller

import (
	"server/common"
	"server/model"
	"server/util"
	"server/vo"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type IPostController interface {
	RestController
	PageList(ctx *gin.Context)
}
type PostController struct {
	DB *gorm.DB
}

func NewPostController() IPostController {
	db := common.GetDB()
	db.AutoMigrate(model.Post{})
	return &PostController{DB: db}
}
func (p PostController) Create(ctx *gin.Context) {
	var requestPost vo.CreatePostRequest
	if ctx.ShouldBind(&requestPost) != nil {
		util.Response{}.Error(ctx, nil, "数据验证错误")
		return
	}
	// 获取登陆用户
	user, _ := ctx.Get("user")
	UserId := user.(model.User).ID
	// 创建post
	post := model.Post{
		UserId:     UserId,
		CategoryId: requestPost.CategoryId,
		Title:      requestPost.Title,
		HeadImg:    requestPost.HeadImg,
		Content:    requestPost.Content,
	}
	if err := p.DB.Create(&post).Error; err != nil {
		util.Response{}.Error(ctx, nil, "创建失败")
		return
		// panic(err)
	}
	util.Response{}.Success(ctx, gin.H{"post": post}, "创建成功")
}

func (p PostController) Update(ctx *gin.Context) {
	var requestPost vo.CreatePostRequest
	if ctx.ShouldBind(&requestPost) != nil {
		util.Response{}.Error(ctx, nil, "数据验证错误")
		return
	}
	// 获取登陆用户
	user, _ := ctx.Get("user")
	// 获取 path 中的id
	postId := ctx.Params.ByName("id")
	var post model.Post
	if p.DB.Where("id = ?", postId).First(&post).RecordNotFound() {
		util.Response{}.Error(ctx, nil, "文章不存在")
	}
	// 判断当前用户是否为文章的作者
	userId := user.(model.User).ID
	if userId != post.UserId {
		panic("文章不属于你，请勿非法操作！")
	}
	// 更新文章

	if err := p.DB.Model(&post).Update(requestPost).Error; err != nil {
		panic("文章更新失败！")
	}

	util.Response{}.Success(ctx, gin.H{"post": post}, "文章修改成功！")
}

func (p PostController) Show(ctx *gin.Context) {
	// 获取 path 中的id
	postId := ctx.Params.ByName("id")
	var post model.Post
	if p.DB.Preload("Category").Where("id = ?", postId).First(&post).RecordNotFound() {
		util.Response{}.Error(ctx, nil, "文章不存在")
		return
	}
	util.Response{}.Success(ctx, gin.H{"post": post}, "")
}

func (p PostController) Delete(ctx *gin.Context) {
	// 获取登陆用户
	user, _ := ctx.Get("user")
	// 获取 path 中的id
	postId := ctx.Params.ByName("id")
	var post model.Post
	if p.DB.Where("id = ?", postId).First(&post).RecordNotFound() {
		util.Response{}.Error(ctx, nil, "文章不存在")
		return
	}
	// 判断当前用户是否为文章的作者
	userId := user.(model.User).ID
	if userId != post.UserId {
		panic("文章不属于你，请勿非法操作！")
	}
	if p.DB.Delete(&post).Error != nil {
		panic("文章删除失败！")
	}
	util.Response{}.Success(ctx, nil, "文章删除成功！")

}

func (p PostController) PageList(ctx *gin.Context) {
	// 获取分页参数
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "20"))
	// 分页
	var posts []model.Post
	p.DB.Order("created_at desc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&posts)

	// 我们要知道记录的总数
	var total int
	p.DB.Model(model.Post{}).Count(&total)
	util.Response{}.Success(ctx, gin.H{"data": posts, "total": total}, "success")
}
