package webserver

import (
	"github.com/cjoshmartin/virtual-experience-api/database"
	"github.com/gin-gonic/gin"
	"net/mail"
)
func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func SetRoutes( router *gin.Engine) {
	mongoDatabase := database.StartDatebase()
	orderCollection := database.OrderInit(mongoDatabase)
	chefCollection := database.ChefInit(mongoDatabase)
	userCollection := database.UserInit(mongoDatabase)
	experienceCollection := database.ExperienceInit(mongoDatabase)

	orders := router.Group("/order")
	{
		orders.POST("/create", CreateOrder(orderCollection, experienceCollection, chefCollection))
		idRoutes := orders.Group("/:id")
		{
			idRoutes.GET("", GetOrderByID(orderCollection))
			idRoutes.POST("/update", UpdateOrder(orderCollection))
		}
	}
	chefs := router.Group("/chef")
	{
		chefs.POST("/create", CreateChef(chefCollection))
		idRoutes := chefs.Group("/:id")
		{
			idRoutes.GET("", GetChefByID(chefCollection))
			idRoutes.POST("/add-experience",  AddExperienceToAChef(chefCollection))
			idRoutes.POST("/update", UpdateChef(chefCollection))
		}
		chefs.GET("/all", GetAllChefs(chefCollection))
	}
	users := router.Group("/user")
	{
		users.POST("/create", CreateUser(userCollection))
		idRoutes := users.Group("/:id")
		{
			idRoutes.GET("", GetUserByID(userCollection))
			idRoutes.POST("/update", UpdateUser(userCollection))
		}
	}
	experiences := router.Group("/experience")
	{
		experiences.POST("/create", CreateExperience(chefCollection, experienceCollection))
		experiences.POST("/add-attendee", AddAttendeeToExperience(userCollection, experienceCollection))
		idRoutes := experiences.Group("/:id")
		{
			idRoutes.GET("", GetExperienceByID(experienceCollection))
			idRoutes.POST("/update", UpdateExperience(experienceCollection))
		}
	}
}

func Start()   {
	router := gin.Default()
	SetRoutes(router)
	router.Run()
}