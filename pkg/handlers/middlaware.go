package handlers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Константы для работы с авторизацией
const (
	// Название заголовка для авторизации
	authorizationHeader = "Authorization"
	// Ключ контекста для хранения ID пользователя
	userIdCtx string = "userId"
)

// userIdentity - промежуточное ПО (middleware) для аутентификации пользователя
// Проверяет токен в заголовке запроса и извлекает ID пользователя
func (handler *Handler) userIdentity(ctx *gin.Context) {
	// Получаем значение заголовка Authorization
	header := ctx.GetHeader(authorizationHeader)
	if header == "" {
		// Если заголовок пустой, возвращаем ошибку
		NewErrorResponse(ctx, http.StatusUnauthorized, "empty auth header")
		return
	}

	// Разделяем заголовок на части (обычно формат: "Bearer token")
	headerParts := strings.Split(header, " ")

	// Проверяем, что заголовок состоит из двух частей
	const validHeaderLength = 2
	if len(headerParts) != validHeaderLength {
		// Если формат неверный, возвращаем ошибку
		NewErrorResponse(ctx, http.StatusUnauthorized, "invalid auth header")
		return
	}

	// Парсим токен и получаем ID пользователя
	// headerParts[1] содержит сам токен (после "Bearer")
	userId, err := handler.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		// Если токен недействительный, возвращаем ошибку
		NewErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	// Сохраняем ID пользователя в контексте для использования в следующих обработчиках
	ctx.Set(userIdCtx, userId)
}

func getUserId(ctx *gin.Context) (int, error) {
	id, ok := ctx.Get(userIdCtx)

	if !ok {
		NewErrorResponse(ctx, http.StatusInternalServerError, "user id not found")
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)

	if !ok {
		NewErrorResponse(ctx, http.StatusInternalServerError, "user id is of invalid type")
		return 0, errors.New("user id is of invalid type")
	}
	return idInt, nil

}
