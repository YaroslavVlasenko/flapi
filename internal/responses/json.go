package responses

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

func Success(status string, message string, data interface{}) interface{} {
	return fiber.Map{
		"status":  status,
		"errors":  nil,
		"data": data,
		"server": ServerInfo(),
		"message": message,
	}
}

func Error(status string, message string, err interface{}) interface{} {
	return fiber.Map{
		"status":  status,
		"errors":  err,
		"payload": nil,
		"server": ServerInfo(),
		"message": message,
	}
}

func ServerInfo() interface{} {
	return fiber.Map{
		"time": time.Now(),
	}
}
