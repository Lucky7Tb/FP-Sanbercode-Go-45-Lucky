package main

import (
	"tulisaja/controller"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func setupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	router.Group("/api")
	{
		router.Group("/auth")
		{
			router.POST("/login", controller.Login)
			router.POST("/register", controller.Register)
		}

		router.Group("/profile")
		{
			router.GET("/", controller.GetProfile)
			router.POST("/change-password", controller.ChangePassword)
		}

		router.Group("/articles")
		{
			router.GET("/", controller.GetRandomArticles)
			router.GET("/:username", controller.GetArticles)
			router.GET("/:username/:id", controller.GetArticle)
			router.POST("/", controller.CreateArticle)
			router.POST("/:username/:id/comment", controller.CommentArticle)
			router.POST("/:username/:id/like", controller.LikeArticle)
			router.PUT("/:id", controller.UpdateArticle)
			router.DELETE("/:id", controller.DeleteArticle)
		}

		router.Group("/following")
		{
			router.GET("/", controller.GetFollowingUser)
			router.POST("/", controller.FollowUser)
			router.DELETE("/:id", controller.DeleteFollowingUser)
		}

		router.GET("/followers", controller.GetFollowers)
	}

	router.GET("/api/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
