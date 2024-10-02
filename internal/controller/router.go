package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/nurzzaat/news/middleware"
	"github.com/nurzzaat/news/pkg"

	_ "github.com/nurzzaat/news/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/nurzzaat/news/internal/controller/auth"
	"github.com/nurzzaat/news/internal/controller/categories"
	"github.com/nurzzaat/news/internal/controller/news"
	"github.com/nurzzaat/news/internal/controller/user"
	"github.com/nurzzaat/news/internal/repository"
)

func Setup(app pkg.Application, router *gin.Engine) {
	env := app.Env
	db := app.Pql

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Static("/news_images", "./news_images")

	loginController := &auth.AuthController{
		UserRepository: repository.NewUserRepository(db),
		Env:            env,
	}

	userController := &user.UserController{
		UserRepository: repository.NewUserRepository(db),
	}
	newsController := &news.NewsController{
		NewsRepository: repository.NewNewsRepository(db),
	}
	categoryController := &categories.CategoryController{
		CategoryRepository: repository.NewCategoryRepository(db),
	}

	router.POST("/register", loginController.Signup)
	router.POST("/login", loginController.Signin)

	newsRouter := router.Group("/news")
	{
		newsRouter.GET("", newsController.GetAll)
		newsRouter.GET("/:id", newsController.GetByID)
	}

	categoriesRouter := router.Group("/categories")
	{
		categoriesRouter.GET("", categoryController.GetAll)
		categoriesRouter.GET("/:id", categoryController.GetByID)
	}

	router.Use(middleware.JWTAuth(env.AccessTokenSecret))

	userRouter := router.Group("/user")
	{
		userRouter.GET("/profile", userController.GetProfile)
	}

	newsAdminRouter := router.Group("/news")
	{
		newsAdminRouter.POST("", newsController.Create)
		newsAdminRouter.PUT("/:id", newsController.Update)
		newsAdminRouter.DELETE("/:id", newsController.Delete)
	}
	categoriesAdminRouter := router.Group("/category")
	{
		categoriesAdminRouter.POST("", categoryController.Create)
		categoriesAdminRouter.PUT("/:id", categoryController.Update)
		categoriesAdminRouter.DELETE("/:id", categoryController.Delete)
	}

}
