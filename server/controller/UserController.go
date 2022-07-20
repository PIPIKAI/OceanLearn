package controller

import (
	"log"
	"net/http"
	"server/common"
	"server/model"

	"server/util"

	"github.com/gin-gonic/gin"
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
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	telephone := ctx.PostForm("telephone")
	if len(telephone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为11位"})
		return
	}
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不少于位"})
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
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号已存在"})
		return
	}
	newUser := model.User{
		Username:  username,
		Password:  password,
		Telephone: telephone,
	}
	db.Create(&newUser)
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "注册成功",
	})

}
