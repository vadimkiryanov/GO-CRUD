package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	todo "github.com/vadimkiryanov/GO-CRUD"
)

// createList обрабатывает HTTP запрос на создание нового списка
func (h *Handler) createList(ctx *gin.Context) {
	// Получаем ID пользователя из контекста запроса
	// который был ранее установлен middleware аутентификации
	// второй возвращаемый параметр (exists) игнорируется с помощью _
	idUser, ok := getUserId(ctx)

	if ok != nil {
		return
	}

	var input todo.TodoList
	if err := ctx.BindJSON(&input); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.TodoList.Create(idUser, input)

	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllListsReponse struct {
	Data []todo.TodoList `json:"data"`
}

func (h *Handler) getAllLists(ctx *gin.Context) {
	idUser, ok := getUserId(ctx)

	if ok != nil {
		return
	}

	lists, err := h.services.TodoList.GetAll(idUser)

	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, getAllListsReponse{
		Data: lists,
	})

}

func (h *Handler) getListById(ctx *gin.Context) {
	idUser, ok := getUserId(ctx)

	if ok != nil {
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, "invalid id param")
		return
	}

	list, err := h.services.TodoList.GetById(idUser, id)

	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, list)

}
func (h *Handler) updateList(ctx *gin.Context) {

}
func (h *Handler) deleteList(ctx *gin.Context) {

}
