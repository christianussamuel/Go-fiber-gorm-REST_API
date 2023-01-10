package controllers

import (
	"context"
	"go_fiber/main/configs"
	"go_fiber/main/models"
	"go_fiber/main/responses"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"

)

var employmentCollection *mongo.Collection = configs.GetCollection(configs.DB, "employment")

func CreateEmployment(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	var employment models.Employment
	profileCode := c.Params("profileCode")
	defer cancel()

	i, _ := strconv.ParseInt(profileCode, 10, 64)

	//validate req body

	if err := c.BodyParser(&employment); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}
	//validate required field using validator library
	if validationErr := validate.Struct(&employment); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	newEmploy := models.Employment{
		ProfileCode: i,
		Id:          (i - (10 ^ 4)),
		JobTitle:    employment.JobTitle,
		Employer:    employment.Employer,
		StartDate:   employment.StartDate,
		EndDate:     employment.EndDate,
		City:        employment.City,
		Description: employment.Description,
	}

	newEmployment := models.Data{
		Employment: []models.Employment{
			newEmploy,
		},
	}

	result, err := employmentCollection.InsertOne(ctx, newEmployment)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.EmploymentResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}
	return c.Status(http.StatusCreated).JSON(responses.EmploymentResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}, ProfileCode: newEmploy.ProfileCode, Id: int64(newEmploy.Id)})
}

// func GetEmployment(c *fiber.Ctx) error{

// 	return c.Status(http.StatusCreated).JSON(responses.GetUserResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})

// }
