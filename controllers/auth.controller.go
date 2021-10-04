package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"mynamebvh.com/blog/config"
	"mynamebvh.com/blog/database"
	"mynamebvh.com/blog/models"
)

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func getUserByEmail(e string) (*models.User, error) {
	db := database.DB
	var user models.User
	if err := db.Where(&models.User{Email: e}).Find(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func Login(c *fiber.Ctx) error{
	type UserLogin struct{
		Email string `json:"email"`
		Password string `json:"password"`
	}

	var userLogin UserLogin

	if err := c.BodyParser(&userLogin); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Lỗi xác thực"})
	}

	user, err := getUserByEmail(userLogin.Email)

	if(err != nil){
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Lỗi xác thực"})
	}

	if user == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Sai email hoặc mật khẩu"})
		
	}

	if !CheckPasswordHash(userLogin.Password, user.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Sai mật khẩu"})
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	t, err := token.SignedString([]byte(config.GetEnv("SECRET")))

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Đăng nhập thành công", "token": t})
}