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

func SetRoutes() *gin.Engine{
	mongoDatabase := database.StartDatebase()
	orderCollection := database.OrderInit(mongoDatabase)
	chefCollection := database.ChefInit(mongoDatabase)
	userCollection := database.UserInit(mongoDatabase)
	experienceCollection := database.ExperienceInit(mongoDatabase)

	r := gin.Default()

	orders := r.Group("/order")
	{
		orders.POST("/create",CreateOrder(orderCollection, experienceCollection, chefCollection))
		orders.GET("/:id",  GetOrderByID(orderCollection))
		orders.POST("/:id/update", UpdateOrder(orderCollection))
	}
	chefs := r.Group("/chef")
	{
		chefs.POST("/create", CreateChef(chefCollection))
		chefs.GET("/:id", GetChefByID(chefCollection))
		chefs.POST("/:id/add-experience",  AddExperienceToAChef(chefCollection))
		chefs.POST("/:id/update", UpdateChef(chefCollection))
		chefs.GET("/all", GetAllChefs(chefCollection))
	}
	users := r.Group("/user")
	{
		users.POST("/create", CreateUser(userCollection))
		users.GET("/:id", GetUserByID(userCollection))
		users.POST("/:id/update", UpdateUser(userCollection))
	}
	experiences := r.Group("/experience")
	{
		experiences.POST("/create", CreateExperience(chefCollection, experienceCollection))
		experiences.GET("/:id", GetExperienceByID(experienceCollection))
		experiences.POST("/:id/update", UpdateExperience(experienceCollection))
	}

	return r
}