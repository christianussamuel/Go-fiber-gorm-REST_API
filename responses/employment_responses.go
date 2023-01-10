package responses

import "github.com/gofiber/fiber/v2"

type EmploymentResponse struct {
	Status      int        `json:"status"`
	Message     string     `json:"message"`
	Data        *fiber.Map `json:"data"`
	ProfileCode int64      `json:"profileCode"`
	Id          int64      `json:"id"`
}
