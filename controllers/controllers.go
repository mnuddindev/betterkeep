package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/mnuddindev/betterkeep/db"
	"github.com/mnuddindev/betterkeep/models"
	"github.com/mnuddindev/betterkeep/utils"
)

func NotImplemented(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"message": "Hello there!!",
		"status":  fiber.StatusOK,
	})
}

func CreateUser(c *fiber.Ctx) error {
	user := new(models.Users)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  err.Error(),
			"status": fiber.StatusBadRequest,
		})
	}
	if utils.IsEmpty(user.FirstName) || utils.IsEmpty(user.LastName) || utils.IsEmpty(user.Email) || utils.IsEmpty(user.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":  "fields cannot be empty",
			"status": fiber.StatusUnauthorized,
		})
	}
	if !utils.IsEmail(user.Email) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":  "invalid email format",
			"status": fiber.StatusUnauthorized,
		})
	}
	uc, err := db.RegistrationHelper(*user)
	if err == nil {
		code := db.GetOTP(uc.ID)
		otp := strconv.FormatInt(code, 10)
		utils.ActiveUser(otp, uc.Email, uc.FirstName)
	} else {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error":  err.Error(),
			"status": fiber.StatusConflict,
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  fiber.StatusCreated,
		"message": "user created",
		"data": fiber.Map{
			"ID":           uc.ID,
			"Name":         uc.FirstName + " " + uc.LastName,
			"Email":        uc.Email,
			"Password":     "Your Password",
			"Verification": uc.Verified,
			"Message":      "Check your Email Box for Verification Code",
			"status":       fiber.StatusCreated,
		},
	})
}

func ActiveUser(c *fiber.Ctx) error {
	type Body struct {
		Otp int64 `json:"otp"`
	}
	demoUser := "00000000-0000-0000-0000-000000000000"
	b := new(Body)
	if err := c.BodyParser(b); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  err.Error(),
			"status": fiber.StatusBadRequest,
		})
	}
	userid, _ := uuid.Parse(c.Params("userid"))
	if userid.String() == demoUser {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":  "user not found",
			"status": fiber.StatusNotFound,
		})
	}
	user, err := db.UserById(userid)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":  err.Error(),
			"status": fiber.StatusNotFound,
		})
	}
	if user.ID.String() == demoUser {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":  "user not found",
			"status": fiber.StatusNotFound,
		})
	}
	if b.Otp != user.Verification {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "otp not matched",
			"status": fiber.StatusBadRequest,
		})
	} else {
		if user.Verified {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":  "otp expired",
				"status": fiber.StatusInternalServerError,
			})
		}
		err = db.UserActive(userid)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":  "failed to active account, internal server error",
				"status": fiber.StatusInternalServerError,
			})
		}
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "login success",
		"data": fiber.Map{
			"user_id":       user.ID,
			"name":          user.FirstName + " " + user.LastName,
			"email":         user.Email,
			"profile_photo": user.ProfilePhoto,
			"password":      "Your Password",
			"message":       "your account activated.please login now!!",
			"status":        fiber.StatusOK,
		},
	})
}

func Login(c *fiber.Ctx) error {

}
