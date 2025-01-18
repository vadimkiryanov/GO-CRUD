package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	todo "github.com/vadimkiryanov/GO-CRUD"
)

// signUp обрабатывает регистрацию пользователя путем привязки JSON-данных к структуре User.
// Возвращает ответ с ошибкой bad request, если привязка JSON не удалась.
func (handler *Handler) signUp(ctx *gin.Context) {
	// Когда пользователь отправляет запрос на регистрацию:
	var user todo.User // Создается пустая структура для данных пользователя

	// Пытаемся прочитать JSON из запроса и записать в структуру user
	if err := ctx.BindJSON(&user); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// Пытаемся создать пользователя через сервисный слой
	id, err := handler.services.Authorization.CreateUser(user)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	// Отправляем успешный ответ (200) с ID созданного пользователя
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
func (handler *Handler) signIn(ctx *gin.Context) {

}
