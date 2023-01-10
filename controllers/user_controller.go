package controllers

import (
	"context"
	"go_fiber/main/configs"
	"go_fiber/main/models"
	"go_fiber/main/responses"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var validate = validator.New()

func CreateUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	var user models.User
	defer cancel()

	//validate req body
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}
	//validate required field using validator library
	if validationErr := validate.Struct(&user); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}
	newUser := models.User{
		Id:             primitive.NewObjectID(),
		ProfileCode:    time.Now().UTC().Unix(),
		WantedJobTitle: user.WantedJobTitle,
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		Email:          user.Email,
		Phone:          user.Phone,
		Country:        user.Country,
		City:           user.City,
		Address:        user.Address,
		PostalCode:     user.PostalCode,
		DrivingLicense: user.DrivingLicense,
		Nationality:    user.Nationality,
		PlaceOfBirth:   user.PlaceOfBirth,
		DateOfBirth:    user.DateOfBirth,
	}

	result, err := userCollection.InsertOne(ctx, newUser)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}, ProfileCode: newUser.ProfileCode})
}
