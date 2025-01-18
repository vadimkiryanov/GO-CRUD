package todo

type User struct {
	Id       int    `json:"-"`
	Name     string `json:"name" binding:"required"`     // Когда поле помечено как binding:"required",
	Username string `json:"username" binding:"required"` // это означает что
	Password string `json:"password" binding:"required"` // это поле является обязательным
}

/*

 */
