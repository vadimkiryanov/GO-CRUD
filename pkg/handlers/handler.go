package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/vadimkiryanov/GO-CRUD/pkg/service"
)

type Handler struct {
	services *service.Service
}

// Инициализация обработчиков
func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

// Инициализация роутеров
func (h *Handler) InitRouters() *gin.Engine {
	router := gin.New() // создание роутера

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp) // регистрация
		auth.POST("/sign-in", h.signIn) // авторизация
	}

	api := router.Group("/api")
	{

		lists := api.Group("/lists") // эндпоинты для работы с списками | /api/lists
		{
			lists.POST("/", h.createList)      // создание списка | /api/lists
			lists.GET("/", h.getAllLists)      // получение всех списков | /api/lists
			lists.GET("/:id", h.getListById)   // получение конкретного списка по id | /api/lists/:id
			lists.PUT("/:id", h.updateList)    // обновление списка по id | /api/lists/:id
			lists.DELETE("/:id", h.deleteList) // удаление списка по id | /api/lists/:id

			items := lists.Group(":id/items") // эндпоинты для работы с элементами списка | /api/lists/:id/items
			{

				items.POST("/", h.createItem)           // создание элемента списка | /api/lists/:id/items
				items.GET("/", h.getAllItems)           // получение всех элементов списка | /api/lists/:id/items
				items.GET("/:item_id", h.getItemById)   // получение конкретного элемента списка по id | /api/lists/:id/items/:item_id
				items.PUT("/:item_id", h.updateItem)    // обновление элемента списка по id | /api/lists/:id/items/:item_id
				items.DELETE("/:item_id", h.deleteItem) // удаление элемента списка по id | /api/lists/:id/items/:item_id
			}
		}
	}

	return router
}
