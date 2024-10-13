package security

import (
	"fmt"
	"go-restapi/internal/database"
	"go-restapi/internal/model"
	"go-restapi/internal/util"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

var (
	lg        = util.NewLogger()
	secretKey = []byte(os.Getenv("SECRETKEY"))
)

type AuthConfig struct {
	User *model.Employee
}

func Auth(db *gorm.DB, u *model.Employee) (*AuthConfig, error) {
	// Cross check with database
	repo := database.InitRepo(db)
	employee, err := repo.FindByID(u.Empid, u.Password)
	if err != nil {
		lg.Fatal(err)
		return &AuthConfig{}, err
	}

	switch {
	case (u.Empid == employee.Empid) && (u.Password == employee.Password):
		return &AuthConfig{
			User: u,
		}, nil

	default:
		return &AuthConfig{
			User: u,
		}, nil
	}

	return &AuthConfig{}, nil
}

func (a *AuthConfig) CreateToken() (string, error) {
	// Create JWT Token
	claims := jwt.MapClaims{
		"empid": a.User.Empid,
		"exp":   time.Now().Add(time.Minute * 5).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func ValidateToken(t string) error {
	// Validate user JWT Token
	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	// Check if Token still valid
	switch {
	case err != nil:
		return err
	case !token.Valid:
		return fmt.Errorf("Invalid Token")
	default:
		return nil
	}
}
