package domains

import "github.com/JooHyeongLee/clean-gin/models"

type AuthService interface {
	Authorize(tokenString string) (bool, error)
	CreateToken(models.User) string
}
