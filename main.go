package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/kamchai-n/api-student-home-visit/model"
	"github.com/kamchai-n/api-student-home-visit/router"
	"github.com/spf13/viper"
)

func main() {

	// Loading Config
	initConfig()

	// Loading Database
	DB := model.InitDatabase()

	// Init fiber
	app := fiber.New(fiber.Config{
		AppName: "Home Visit ITCMTC v1",
	})
	app.Use(cors.New())
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	api := app.Group("/api")

	router.StudentRouter(api.Group("/student"), DB)
	router.AuthRouter(api.Group("/auth"), DB)
	router.SettingRouter(api.Group("/setting"), DB)
	router.VisitRouter(api.Group("/visit"), DB)
	router.StaticRouter(api.Group("/static"), DB)

	port := os.Getenv("PORT")

	if port == "" {
		port = strconv.Itoa(viper.GetInt("post"))
	}
	app.Listen(fmt.Sprintf(":%v", port))
}

func initConfig() {
	if os.Getenv("RUN_MODE") == "production" {
		fmt.Println("production")
		viper.SetConfigName("production")
	} else if os.Getenv("RUN_MODE") == "production_me" {
		fmt.Println("production_me")
		viper.SetConfigName("production_me")
	} else {
		fmt.Println("development")
		viper.SetConfigName("development")
	}
	viper.SetConfigType("json")
	viper.AddConfigPath("./configs")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
