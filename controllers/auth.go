package controllers

import (
	"time"

	jwt "github.com/form3tech-oss/jwt-go"

	"github.com/gofiber/fiber/v2"
	"github.com/nathakrit061103jnt/crud_fiber_go_gorm/database"
	"github.com/nathakrit061103jnt/crud_fiber_go_gorm/models"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	db := database.DBConn

	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return err
	}

	user.U_password, _ = hashPassword(user.U_password)

	if err := db.Create(&user).Error; err != nil {
		// error handling...
		return err
	}

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {

	user := new(models.User)

	db := database.DBConn

	if err := c.BodyParser(user); err != nil {
		return err
	}

	passwordInput := user.U_password
	user.U_password, _ = hashPassword(user.U_password)

	if err := db.First(&user, "u_email = ?", user.U_email).Error; err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	} else if checkPasswordHash(passwordInput, user.U_password) == false {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.U_email
	claims["admin"] = false
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"result": "ok", "token": t})

}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
