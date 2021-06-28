package webserver

import (
	"github.com/cjoshmartin/virtual-experience-api/database"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/mail"
)
func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func setRoutes( router *gin.Engine) {
	mongoDatabase := database.StartDatebase()
	orderCollection := database.OrderInit(mongoDatabase)
	chefCollection := database.ChefInit(mongoDatabase)
	userCollection := database.UserInit(mongoDatabase)
	experienceCollection := database.ExperienceInit(mongoDatabase)

	chefs := router.Group("/chef")
	{
		chefs.POST("/create", CreateChef(chefCollection))
		idRoutes := chefs.Group("/:id")
		{
			idRoutes.GET("", GetChefByID(chefCollection))
			idRoutes.GET("/experiences", GetAChefExperiences(chefCollection, experienceCollection))
		}
		chefs.GET("/all", GetAllChefs(chefCollection))
	}
	users := router.Group("/user")
	{
		users.POST("/create", CreateUser(userCollection))
		idRoutes := users.Group("/:id")
		{
			idRoutes.GET("", GetUserByID(userCollection))
		}

		users.GET("/all", GetAllUsers(userCollection))
	}
	experiences := router.Group("/experience")
	{
		experiences.POST("/create", CreateExperience(chefCollection, experienceCollection))
		experiences.POST("/add-attendee", AddAttendeeToExperience(userCollection, experienceCollection))
		idRoutes := experiences.Group("/:id")
		{
			idRoutes.GET("", GetExperienceByID(experienceCollection))
		}
		experiences.GET("/all", GetAllExperiences(experienceCollection))
	}
	orders := router.Group("/order")
	{
		orders.POST("/create", CreateOrder(orderCollection, experienceCollection, chefCollection, userCollection))
		idRoutes := orders.Group("/:id")
		{
			idRoutes.GET("", GetOrderByID(orderCollection))
		}
	}
}

func setStaticRoutes(router *gin.Engine){
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.GET("/user", func(c *gin.Context) {
		c.HTML(http.StatusOK, "user.html", nil)
	})
	router.GET("/experience", func(c *gin.Context) {
		c.HTML(http.StatusOK, "experience.html", nil)
	})
	router.GET("/chef", func(c *gin.Context) {
		c.HTML(http.StatusOK, "chef.html", nil)
	})
	router.GET("/order", func(c *gin.Context) {
		c.HTML(http.StatusOK, "order.html", nil)
	})
}

func Start()   {
	router := gin.Default()
	setRoutes(router)
	setStaticRoutes(router)
	router.Run()
}