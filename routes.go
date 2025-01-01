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
			profileRoute.POST("/change-password", middleware.VerifyJwtToken(), controller.ChangePassword)
		}

		articleRoute := baseRouter.Group("/articles")
		{
			articleRoute.GET("/", controller.GetRandomArticles)
			articleRoute.GET("/:username", controller.GetArticles)
			articleRoute.GET("/:username/:id", controller.GetArticle)
			articleRoute.POST("/", middleware.VerifyJwtToken(), controller.CreateArticle)
			articleRoute.POST("/:username/:id/comment", middleware.VerifyJwtToken(), controller.CommentArticle)
			articleRoute.POST("/:username/:id/like", middleware.VerifyJwtToken(), controller.LikeArticle)
			articleRoute.PUT("/:id", middleware.VerifyJwtToken(), controller.UpdateArticle)
		}

		followingRouter := baseRouter.Group("/following")
		{
			followingRouter.GET("/", middleware.VerifyJwtToken(), controller.GetFollowingUser)
			followingRouter.POST("/", middleware.VerifyJwtToken(), controller.FollowUser)
			followingRouter.DELETE("/:id", middleware.VerifyJwtToken(), controller.DeleteFollowingUser)
		}

		baseRouter.GET("/followers", middleware.VerifyJwtToken(), controller.GetFollowers)
	}

	router.GET("/api/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
