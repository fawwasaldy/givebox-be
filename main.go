package main

import (
	"github.com/gin-gonic/gin"
	"givebox/application/service"
	"givebox/command"
	"givebox/infrastructure/database/config"
	infrastructure_category "givebox/infrastructure/database/donation/category"
	infrastructure_donated_item "givebox/infrastructure/database/donation/donated_item"
	infrastructure_donated_item_recipient "givebox/infrastructure/database/donation/donated_item_recipient"
	infrastructure_image "givebox/infrastructure/database/donation/image"
	infrastructure_user "givebox/infrastructure/database/profile/user"
	infrastructure_refresh_token "givebox/infrastructure/database/refresh_token"
	"givebox/infrastructure/database/transaction"
	"givebox/presentation/controller"
	"givebox/presentation/middleware"
	"givebox/presentation/route"
	"gorm.io/gorm"
	"log"
	"os"
)

func args(db *gorm.DB) bool {
	if len(os.Args) > 1 {
		flag := command.Commands(db)
		return flag
	}

	return true
}

func run(server *gin.Engine) {
	server.Static("/assets", "./assets")

	if os.Getenv("IS_LOGGER") == "true" {
		route.LoggerRoute(server)
	}

	port := os.Getenv("GOLANG_PORT")
	if port == "" {
		port = "8888"
	}

	var serve string
	if os.Getenv("APP_ENV") == "localhost" {
		serve = "0.0.0.0:" + port
	} else {
		serve = ":" + port
	}

	if err := server.Run(serve); err != nil {
		log.Fatalf("error running server: %v", err)
	}
}

func main() {
	db := config.SetUpDatabaseConnection()

	jwtService := service.NewJWTService()

	// repositories
	transactionRepository := transaction.NewRepository(db)
	refreshTokenRepository := infrastructure_refresh_token.NewRepository(transactionRepository)
	userRepository := infrastructure_user.NewRepository(transactionRepository)
	donatedItemRepository := infrastructure_donated_item.NewRepository(transactionRepository)
	donatedItemRecipientRepository := infrastructure_donated_item_recipient.NewRepository(transactionRepository)
	imageRepository := infrastructure_image.NewRepository(transactionRepository)
	categoryRepository := infrastructure_category.NewRepository(transactionRepository)

	// services
	userService := service.NewUserService(
		userRepository,
		refreshTokenRepository,
		jwtService,
		transactionRepository)
	donationService := service.NewDonationService(
		donatedItemRepository,
		donatedItemRecipientRepository,
		imageRepository,
		categoryRepository,
		userRepository,
		transactionRepository)

	// controllers
	userController := controller.NewUserController(userService)
	donationController := controller.NewDonationController(donationService)

	defer config.CloseDatabaseConnection(db)

	if !args(db) {
		return
	}

	server := gin.Default()
	server.Use(middleware.CORSMiddleware())

	// routes
	route.UserRoute(server, userController, jwtService)
	route.DonationRoute(server, donationController, jwtService)

	run(server)
}
