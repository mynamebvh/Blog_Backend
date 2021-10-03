package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"mynamebvh.com/blog/database"
	"mynamebvh.com/blog/models"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func GetUser(c *fiber.Ctx) error{
	id := c.Params("id")
	db := database.DB

	var user models.User
	db.Find(&user, id)

	if user.Fullname == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Người dùng không tồn tại"})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "success", "data": user})
}

func SignUp(c *fiber.Ctx) error{
  user := new(models.User)
	db := database.DB

	if err:= c.BodyParser(user); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status":"error", "message": "Kiểm tra lại thông tin", "data": err})
	}

	hash, err := hashPassword(user.Password)
	if err != nil{
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status":"error", "message": "Lỗi máy chủ", "data": err})
	}
  
	user.Password = hash
	fmt.Println(user)

	if err := db.Create(&user); err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status":"error", "message": "Lỗi máy chủ", "data": err})
	}
	return c.Status(201).JSON(fiber.Map{"status":"success", "message": "Tạo tài khoản thành công"})
}

func UpdateUser(c *fiber.Ctx){

}

func DeleteUser(c *fiber.Ctx) error{
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(float64)
	_ = userId

	db := database.DB
	
	if err := db.First(&models.User{}, userId).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Người dùng không tồn tại"})
	}

	if err := db.Delete(&models.User{}, userId).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Lỗi máy chủ"})
	}
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"status": "success", "message": "Xoá thành công"})
}