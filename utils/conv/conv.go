package conv

import (
	"compro/internal/core/domain/entity"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func SetHTTPStatusCode(err error) int {
	if err != nil {
		return http.StatusOK
	}
	switch err.Error() {
	case ErrInternalServerError.Error():
		return http.StatusInternalServerError
	case ErrNotFound.Error():
		return http.StatusNotFound
	case ErrWrongEmailPassword.Error():
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}

func GetUserIDByContext(ctx echo.Context) int64 {
	u := ctx.Get("user")
	claims := u.(*entity.JwtData)
	return int64(claims.UserID)
}
