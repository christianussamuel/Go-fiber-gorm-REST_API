package controllers

import (
	"context"
	"go_fiber/main/models"
	"go_fiber/main/responses"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"

)

func GetWorkingExperience(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	profileCode := c.Params("profileCode")
	var master models.Master
	defer cancel()

	//get profile code
	i, _ := strconv.ParseInt(profileCode, 10, 64)
	err := masterCollection.FindOne(ctx, bson.M{"profilecode": i}).Decode(&master)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.GetUserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": master}})
}

func EditWorkingExperience(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	profileCode := c.Params("profileCode")
	var master models.Master
	defer cancel()

	i, _ := strconv.ParseInt(profileCode, 10, 64)

	//validate the request body
	if err := c.BodyParser(&master); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&master); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	update := bson.M{
		"workingExperience": master.WorkingExperience,
	}

	result, err := masterCollection.UpdateOne(ctx, bson.M{"profilecode": i}, bson.M{"$set": update})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}
	var updatedMaster models.Master
	if result.MatchedCount == 1 {
		err := masterCollection.FindOne(ctx, bson.M{"profilecode": i}).Decode(&updatedMaster)

		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}
	}

	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": updatedMaster}})
}
