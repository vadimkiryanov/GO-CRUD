package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// createList обрабатывает HTTP запрос на создание нового списка
func (h *Handler) createList(ctx *gin.Context) {
    // Получаем ID пользователя из контекста запроса
    // который был ранее установлен middleware аутентификации
    // второй возвращаемый параметр (exists) игнорируется с помощью _
    idUser, _ := ctx.Get(userIdCtx)

    // Отправляем успешный ответ клиенту
    // со статусом 200 OK и JSON объектом, содержащим ID пользователя
    ctx.JSON(http.StatusOK, map[string]interface{}{
        "id": idUser,
    })
}
func (h *Handler) getAllLists(ctx *gin.Context) {
	log.Println("createList")

}
func (h *Handler) getListById(ctx *gin.Context) {

}
func (h *Handler) updateList(ctx *gin.Context) {

}
func (h *Handler) deleteList(ctx *gin.Context) {

}
