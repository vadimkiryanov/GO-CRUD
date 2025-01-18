package service

import (
	"crypto/sha1"
	"fmt"

	todo "github.com/vadimkiryanov/GO-CRUD"
	"github.com/vadimkiryanov/GO-CRUD/pkg/repository"
)

// AuthService структура для работы с аутентификацией
type AuthService struct {
	repository repository.Authorization // Интерфейс для работы с хранилищем данных
}

// NewAuthService создает новый экземпляр сервиса аутентификации
func NewAuthService(repository repository.Authorization) *AuthService {
	return &AuthService{repository: repository}
}

// CreateUser создает нового пользователя
func (service *AuthService) CreateUser(user todo.User) (int, error) {
	// Генерируем хеш пароля перед сохранением
	user.Password = generatePasswordHash(user.Password)
	// Делегируем создание пользователя репозиторию
	return service.repository.CreateUser(user)
}

// generatePasswordHash создает хеш пароля с использованием SHA1 и соли
func generatePasswordHash(password string) string {
	// Соль для усиления безопасности хеша
	salt := "1a2b3c4dasdasdasd78921928"

	// Создаем новый SHA1 хеш
	hash := sha1.New()
	// Записываем пароль в хеш
	hash.Write([]byte(password))

	// Возвращаем хеш в виде шестнадцатеричной строки,
	// добавляя соль к результату
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
