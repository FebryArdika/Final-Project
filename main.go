package main

import (
	commentDelivery "Final-Project/comment/delivery"
	commentRepository "Final-Project/comment/repository"
	commentUseCase "Final-Project/comment/usecase"
	"Final-Project/database"
	photoDelivery "Final-Project/photo/delivery"
	photoRepository "Final-Project/photo/repository"
	photoUseCase "Final-Project/photo/usecase"
	socialMediaDelivery "Final-Project/socialmedia/delivery"
	socialMediaRepository "Final-Project/socialmedia/repository"
	socialMediaUseCase "Final-Project/socialmedia/usecase"
	userDelivery "Final-Project/user/delivery"
	userRepository "Final-Project/user/repository"
	userUseCase "Final-Project/user/usecase"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	db := database.StartDB()
	routers := gin.Default()
	routers.Use(func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Content-Type", "application/json")
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Max-Age", "86400")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, UPDATE")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(200)
		} else {
			ctx.Next()
		}
	})
	userRepository := userRepository.NewUserRepository(db)
	userUseCase := userUseCase.NewUserUseCase(userRepository)
	userDelivery.NewUserHandler(routers, userUseCase)

	photoRepository := photoRepository.NewPhotoRepository(db)
	photoUseCase := photoUseCase.NewPhotoUsecase(photoRepository)
	photoDelivery.NewPhotoHandler(routers, photoUseCase)

	commentRepository := commentRepository.NewCommentRepository(db)
	commentUseCase := commentUseCase.NewCommentUseCase(commentRepository)
	commentDelivery.NewCommentHandler(routers, commentUseCase, photoUseCase)

	socialMediaRepository := socialMediaRepository.NewSocialMediaRepository(db)
	socialMediaUseCase := socialMediaUseCase.NewSocialMediaUseCase(socialMediaRepository)
	socialMediaDelivery.NewSocialMediaHandler(routers, socialMediaUseCase)

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file: ", err)
	}
	port := os.Getenv("PORT")
	if len(os.Args) > 1 {
		requestPort := os.Args[1]
		if requestPort != "" {
			port = requestPort
		}
	}
	if port == "" {
		port = "6969"
	}
	routers.Run(":" + port)
}
