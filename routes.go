package main

import (
	"tulisaja/controller"
	"tulisaja/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func setupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	router.Use(gin.ErrorLogger())
	router.Use(gin.Recovery())
	router.SetTrustedProxies([]string{"127.0.0.1"})
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	baseRouter := router.Group("/api")
	{
		authRoute := baseRouter.Group("/auth")
		{
			authRoute.POST("/login", controller.Login)
			authRoute.POST("/register", controller.Register)
		}

		profileRoute := baseRouter.Group("/profile")
		{
			profileRoute.GET("/", middleware.VerifyJwtToken(), controller.GetProfile)
			profileRoute.POST("/change-password", controller.ChangePassword)
		}

		articleRoute := baseRouter.Group("/articles")
		{
			articleRoute.GET("/", controller.GetRandomArticles)
			articleRoute.GET("/:username", controller.GetArticles)
			articleRoute.GET("/:username/:id", controller.GetArticle)
			articleRoute.POST("/", controller.CreateArticle)
			articleRoute.POST("/:username/:id/comment", controller.CommentArticle)
			articleRoute.POST("/:username/:id/like", controller.LikeArticle)
			articleRoute.PUT("/:id", controller.UpdateArticle)
			articleRoute.DELETE("/:id", controller.DeleteArticle)
		}

		followingRouter := baseRouter.Group("/following")
		{
			followingRouter.GET("/", controller.GetFollowingUser)
			followingRouter.POST("/", controller.FollowUser)
			followingRouter.DELETE("/:id", controller.DeleteFollowingUser)
		}

		baseRouter.GET("/followers", controller.GetFollowers)
	}

	router.GET("/api/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
