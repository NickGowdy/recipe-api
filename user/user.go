package user

import (
	"github.com/recipe-api/repository"
	"github.com/recipe-api/security"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	repo *repository.UserRepository
}

func NewUser(repo *repository.UserRepository) *User {
	return &User{
		repo: repo,
	}
}

func (u User) Register(firstname string, lastname string, email string, password string) (int64, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		return 0, err
	}
	hashedPasswordStr := string(hashedPassword)

	userId, err := u.repo.InsertRecipeUser(firstname, lastname, email, hashedPasswordStr)
	if err != nil {
		return 0, err
	}
	return userId, nil
}

func (u User) Login(email string, password string) (*string, error) {
	user, err := u.repo.GetRecipeUserPwd(email)
	if err != nil {
		return nil, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}

	tokenString, err := security.GenerateToken(int64(user.ID))
	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}
