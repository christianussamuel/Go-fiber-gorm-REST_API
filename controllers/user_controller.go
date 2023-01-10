package controllers

import (
	"context"
	"go_fiber/main/configs"
	"go_fiber/main/models"
	"go_fiber/main/responses"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var masterCollection *mongo.Collection = configs.GetCollection(configs.DB, "master")

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
		ProfileCode:    time.Now().UTC().UnixMilli(),
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
	newMaster := models.Master{
		ProfileCode: time.Now().UTC().Unix(),
	}

	result, err := userCollection.InsertOne(ctx, newUser)
	_, err_master := masterCollection.InsertOne(ctx, newMaster)

	if err != nil || err_master != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}, ProfileCode: newUser.ProfileCode})
}

func GetAUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	profileCode := c.Params("profileCode")
	var user models.User
	defer cancel()

	//get profile code
	i, _ := strconv.ParseInt(profileCode, 10, 64)
	err := userCollection.FindOne(ctx, bson.M{"profilecode": i}).Decode(&user)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.GetUserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": user}})
}

func EditAUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	profileCode := c.Params("profileCode")
	var user models.User
	defer cancel()

	i, _ := strconv.ParseInt(profileCode, 10, 64)

	//validate the request body
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&user); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	update := bson.M{
		"wantedJobTitle": user.WantedJobTitle,
		"firstName":      user.FirstName,
		"lastName":       user.LastName,
		"email":          user.Email,
		"phone":          user.Phone,
		"country":        user.Country,
		"city":           user.City,
		"address":        user.Address,
		"postalCode":     user.Address,
		"drivingLicense": user.DrivingLicense,
		"nationality":    user.Nationality,
		"placeOfBirth":   user.PlaceOfBirth,
		"dateOfBirth":    user.DateOfBirth,
	}

	result, err := userCollection.UpdateOne(ctx, bson.M{"profilecode": i}, bson.M{"$set": update})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}
	//get updated user details
	var updatedUser models.User
	if result.MatchedCount == 1 {
		err := userCollection.FindOne(ctx, bson.M{"profilecode": i}).Decode(&updatedUser)

		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}
	}

	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": updatedUser}})
}
