package main

import (
	"github.com/ernestechie/go-blog/models"
	"github.com/ernestechie/go-blog/routes"
	"github.com/gin-gonic/gin"
)

var db = make(map[string]models.UserModel)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	router := gin.Default()

	// Get all users
	articleRoutes := router.Group("/articles")
	articleRoutes.GET("/", routes.GetAllArticles)
	articleRoutes.GET("/:id", routes.GetArticle)
	articleRoutes.POST("/", routes.CreateArticle)
	
	// Get user value
	// router.GET("/users/:name", func(c *gin.Context) {
	// 	user := c.Params.ByName("name")
	
	// 	userExists, ok := db[user]
	// 	if ok {
	// 		// fmt.Println("Birthday", userExists.GetBirthday())
	// 		c.JSON(http.StatusOK, gin.H{"data": gin.H{
	// 			"user": userExists,
	// 		}, "success": true, "message": "User successfully retrieved"})
	// 	} else {
	// 		c.JSON(http.StatusNotFound, gin.H{"data": nil, "success": true, "message": "User not found"})
	// 	}
	// })

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	// authorized := router.Group("/", gin.BasicAuth(gin.Accounts{
	// 	"foo":  "bar", // user:foo password:bar
	// 	"manu": "123", // user:manu password:123
	// }))

	/* example curl for /admin with basicauth header
	   Zm9vOmJhcg== is base64("foo:bar")

		curl -X POST \
	  	http://localhost:8080/admin \
	  	-H 'authorization: Basic Zm9vOmJhcg==' \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"bar"}'
	*/
	// authorized.POST("admin", func(c *gin.Context) {
	// 	user := c.MustGet(gin.AuthUserKey).(string)

	// 	// Parse JSON
	// 	var json struct {
	// 		Value 	string 		`json:"value" binding:"required"`
	// 		DoB			time.Time `json:"dob" binding:"required"`
	// 	}

	// 	fmt.Println(user)

	// 	if c.Bind(&json) == nil {
	// 		newUser := models.UserModel{
	// 			// Name: json.Value,
	// 			Role: models.User,
	// 			// DoB: json.DoB,
	// 		}

	// 		db[json.Value] = newUser

	// 		c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{
	// 			"user": newUser,
	// 		}})
	// 	}
	// })

	return router
}

func main() {
	router := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	router.Run(":8080")
}
