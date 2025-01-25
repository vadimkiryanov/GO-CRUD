package todo

type TodoList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"` // Поле обязательное
	Description string `json:"description" db:"description"`
}

type TodoItem struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type UsersList struct {
	Id     int
	UserId int
	ListId int
}

type ListsItem struct {
	Id     int
	ListId int
	ItemId int
}
