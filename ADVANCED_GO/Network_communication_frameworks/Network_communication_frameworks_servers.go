// Try 3 frameworks to deploy serves processing API
package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"github.com/labstack/echo/v4"
)

type UserRequest struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	// 1. GIN
	r := gin.Default()
	r.POST(
		"/users/new-user",
		func(c *gin.Context) {
			var userRequest UserRequest
			err := c.ShouldBindJSON(&userRequest)
			if err != nil {
				c.JSON(400, gin.H{"error": fmt.Sprint("[GIN] error:", err.Error())})
				return
			}
			fmt.Printf("GIN: %+v \n", userRequest)
			c.JSON(200, gin.H{"message": "[GIN]Success"})
		},
	)
	r.GET(
		"/error/try-error-1",
		func(c *gin.Context) {
			err := errors.New("[GIN]error type 1")
			c.JSON(400, gin.H{"error": fmt.Sprint("[GIN] error: ", err.Error())})
		},
	)
	r.GET(
		"/error/try-error-2",
		func(c *gin.Context) {
			err := errors.New("[GIN]error type 2")
			c.JSON(400, gin.H{"message": fmt.Sprint("[GIN] error: ", err.Error())})
		},
	)

	go r.Run(":8080")

	// 2.ECHO
	e := echo.New()
	e.POST(
		"/users/new-user",
		func(c echo.Context) error {
			var userRequest UserRequest
			err := c.Bind(&userRequest)
			if err != nil {
				fmt.Println("[ECHO] error in Binding JSON body: ", err.Error())
				return err
			}
			return c.JSON(200, echo.Map{"message": "[ECHO]Success"})
		},
	)
	e.GET(
		"/error/try-error-1",
		func(c echo.Context) error {
			err := errors.New("[ECHO]error type 1")
			return err
		},
	)
	e.GET(
		"/error/try-error-2",
		func(c echo.Context) error {
			err := errors.New("[ECHO]error type 2")
			return err
		},
	)
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		c.JSON(
			400,
			echo.Map{"message": fmt.Sprint("[ECHO] Error: ", err.Error())})
	}
	go e.Start(":8081")

	// 3. FIBER
	app := fiber.New()

	app.Post(
		"/users/new-user",
		func(c *fiber.Ctx) error {
			var userRequest UserRequest
			err := c.BodyParser(&userRequest)
			if err != nil {
				fmt.Println("[FIBER] error: ", err.Error())
				c.Status(400)
				c.JSON(fiber.Map{"error": err.Error()})
				return err
			}
			c.Status(200)
			return c.JSON(fiber.Map{"message": "[FIBER]Success"})
		},
	)

	app.Get(
		"/error/try-error-1",
		func(c *fiber.Ctx) error {
			err := errors.New("[FIBER]error type 1")
			c.Status(400)
			return c.JSON(fiber.Map{"error": fmt.Sprint("[FIBER] error: ", err.Error())})
		},
	)

	app.Get(
		"/error/try-error-2",
		func(c *fiber.Ctx) error {
			err := errors.New("[FIBER]error type 2")
			c.Status(400)
			return c.JSON(fiber.Map{"message": fmt.Sprint("[FIBER] error: ", err.Error())})
		},
	)

	go app.Listen(":8082")

	time.Sleep(5 * time.Minute)
}
