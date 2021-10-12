package services

import (
	"errors"
	"log"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"mynamebvh.com/blog/internal/entities"
	utils "mynamebvh.com/blog/internal/utils"
	"mynamebvh.com/blog/src/user/dto"
	"mynamebvh.com/blog/src/user/repositories"
)

type UserService interface {
	FindByID(id uint) entities.User
	Login(data *dto.UserLogin) (dto.JwtResponse, error)
	Signup(data *dto.UserRequest) (entities.User, error)
	Delete(id uint) error
	Update(id uint, userUpdate dto.UserUpdate) (entities.User, error)
}

type userService struct {
	userRepository repositories.UserRepositoryInterface
	jwtAuth        utils.JwtTokenInterface
}

func NewUserService(
	userRepository repositories.UserRepositoryInterface,
	jwtAuth utils.JwtTokenInterface,
) UserService {
	return &userService{
		userRepository: userRepository,
		jwtAuth:        jwtAuth,
	}
}

func (c *userService) FindByID(id uint) entities.User {
	return c.userRepository.FindByID(id)
}

func (c *userService) Login(data *dto.UserLogin) (dto.JwtResponse, error) {

	user := c.userRepository.FindByEmail(data.Email)

	if user.ID == 0 {
		return dto.JwtResponse{}, errors.New("tài khoản hoặc mật khẩu sai")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))

	if err != nil {
		return dto.JwtResponse{}, err
	}

	userToken := utils.Sign(jwt.MapClaims{
		"id": user.ID,
	})

	token := dto.JwtResponse(userToken)

	return token, nil
}

func (c *userService) Signup(data *dto.UserRequest) (entities.User, error) {

	user := entities.User{
		Fullname: data.Fullname,
		Email:    data.Email,
		Gender:   data.Gender,
		Password: data.Password,
	}

	hash, err := utils.HashPassword(user.Password)

	if err != nil {
		log.Fatal(err)
	}

	user.Password = hash

	c.userRepository.Save(user)
	return user, nil
}

func (c *userService) Delete(id uint) error {
	if err := c.userRepository.Delete(id); err != nil {
		return err
	}
	return nil
}

func (c *userService) Update(id uint, userUpdate dto.UserUpdate) (entities.User, error) {
	if user, err := c.userRepository.Update(id, userUpdate); err != nil {
		return entities.User{}, err
	} else {
		return user, nil
	}

}
