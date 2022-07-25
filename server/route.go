package main

import (
	"server/controller"
	"server/middleware"

	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware(), middleware.RecoveryMiddleware())
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)

	categoryRouter := r.Group("/category")
	categoryController := controller.NewCategoryController()
	categoryRouter.POST("", categoryController.Create)
	categoryRouter.PUT("/:id", categoryController.Update)
	categoryRouter.GET("/:id", categoryController.Show)
	categoryRouter.DELETE("/:id", categoryController.Delete)

	postRouter := r.Group("/post")
	postRouter.Use(middleware.AuthMiddleware())
	postController := controller.NewPostController()
	postRouter.POST("", postController.Create)
	postRouter.PUT("/:id", postController.Update)
	postRouter.GET("/:id", postController.Show)
	postRouter.DELETE("/:id", postController.Delete)
	postRouter.POST("/page/list", postController.PageList)

	return r
}
