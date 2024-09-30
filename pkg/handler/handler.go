package handler

import (
	"github.com/Chazarov/rest-app/pkg/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {

	router := gin.New()

	auth := router.Group("/auth")
	{
		// signUp обрабатывает регистрацию нового пользователя
		// @Summary User Sign Up
		// @Description Register a new user and return the user ID
		// @Tags auth
		// @Accept json
		// @Produce json
		// @Param input body project.User true "User registration input"
		// @Success 201 {object} map[string]interface{}{"id": "int"}
		// @Failure 400 {object} gin.H{"error": "Bad request"}
		// @Failure 500 {object} gin.H{"error": "Internal server error"}
		// @Router /auth/sign-up [post]
		auth.POST("/sign-up", h.signUp)

		// signIn обрабатывает вход пользователя и генерирует токен
		// @Summary User Sign In
		// @Description Sign in an existing user and receive a JWT token
		// @Tags auth
		// @Accept json
		// @Produce json
		// @Param input body signInInput true "Sign In Input"
		// @Success 200 {object} map[string]interface{}{"token": "string"}
		// @Failure 400 {object} gin.H{"error": "Bad request"}
		// @Failure 500 {object} gin.H{"error": "Internal server error"}
		// @Router /auth/sign-in [post]
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		items := api.Group(":id/items")
		{
			// @Summary Create Item
			// @Description Create a new advertisement item for the user
			// @Tags items
			// @Accept json
			// @Produce json
			// @Param input body project.AdvertItem true "Advertisement item input"
			// @Success 201 {object} map[string]interface{}{"id": "int"}
			// @Failure 400 {object} gin.H{"error": "Bad request"}
			// @Failure 500 {object} gin.H{"error": "Internal server error"}
			// @Router /api/{id}/items [post]
			items.POST("/", h.createItem)

			items.GET("/", h.getAllItems)

			items.GET("/:item_id", h.getItemById)

			items.PUT("/:item_id", h.updateItem)

			items.DELETE("/:item_id", h.deleteItem)
		}
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
