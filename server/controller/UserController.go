package controller

import (
	"log"
	"net/http"
	"server/common"
	"server/model"

	"server/util"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func isTelephoneExist(db *gorm.DB, tel string) bool {
	var user model.User
	db.Where("telephone=?", tel).First(&user)
	if user.ID == 0 {
		return false
	}
	return true
}
func Register(ctx *gin.Context) {

	// 验证数据
	db := common.GetDB()

	// 这种获取不到
	// username := ctx.PostForm("username")
	// password := ctx.PostForm("password")
	// telephone := ctx.PostForm("telephone")

	// 用这种方式双向绑定
	var request = model.User{}
	ctx.Bind(&request)
	username := request.Username
	password := request.Password
	telephone := request.Telephone

	if len(telephone) != 11 {
		util.Response{}.ResponsFmt(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	if len(password) < 6 {
		util.Response{}.ResponsFmt(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不少于6位")

		return
	}
	// 若用户名为空则随机创建
	if username == "" {
		username = util.RandomUsername(10)

	}
	log.Println(username, password, telephone)
	// 在数据库查询手机号是否注册
	// fmt.Printf("在数据库查询手机号是否注册")
	if isTelephoneExist(db, telephone) {
		util.Response{}.ResponsFmt(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号已存在")
		return
	}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		util.Response{}.ResponsFmt(ctx, http.StatusInternalServerError, 500, nil, "密码加密错误")
		return
	}
	newUser := model.User{
		Username:  username,
		Password:  string(hashedPassword),
		Telephone: telephone,
	}
	db.Create(&newUser)

	// util.Response{}.Success(ctx, nil, "注册成功")
	// 发放token

	token, err := common.ReleaseToken(newUser)
	if err != nil {
		util.Response{}.ResponsFmt(ctx, http.StatusInternalServerError, 500, nil, "系统异常")
		log.Printf("token generate error: %v", err.Error())
		return
	}

	// 返回结果
	util.Response{}.Success(ctx, gin.H{"token": token}, "注册成功")

}
func Login(ctx *gin.Context) {

	// 验证数据
	db := common.GetDB()

	// 用这种方式双向绑定
	var request = model.User{}
	ctx.Bind(&request)
	password := request.Password
	telephone := request.Telephone
	if len(telephone) != 11 {
		util.Response{}.ResponsFmt(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")

		return
	}
	// 在数据库查询手机号是否注册
	// fmt.Printf("在数据库查询手机号是否注册")
	var user model.User
	db.Where("telephone=?", telephone).First(&user)
	if user.ID == 0 {
		util.Response{}.ResponsFmt(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号不存在")

		return
	}
	// 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		util.Response{}.Error(ctx, nil, "密码错误")
		return
	}
	// 发放token

	token, err := common.ReleaseToken(user)
	if err != nil {
		util.Response{}.ResponsFmt(ctx, http.StatusInternalServerError, 500, nil, "系统异常")
		log.Printf("token generate error: %v", err.Error())
		return
	}

	// 返回结果
	util.Response{}.Success(ctx, gin.H{"token": token}, "登录成功")

}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	util.Response{}.Success(ctx, gin.H{
		"code": 200,
		"data": gin.H{
			"user": user,
		},
	}, "获取用户信息成功")
}
