package main

import (
	"log"
	"os"
	"strings"

	"github.com/ernestechie/go-blog/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init()  {
	// Load environment variables.
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// Middleware to remove trailing slashes
func stripTrailingSlash() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		if path != "/" && strings.HasSuffix(path, "/") {
			c.Request.URL.Path = strings.TrimSuffix(path, "/")
		}
		c.Next()
	}
}

func setupRouter() *gin.Engine {
	var reactAppUri string
	if reactAppUri = os.Getenv("REACT_APP_URL"); reactAppUri == "" {
		log.Fatal("You must set your 'REACT_APP_URL' environment variable.")
	}

	// Disable Console Color
	// gin.DisableConsoleColor()
	router := gin.Default()
	// router.RedirectTrailingSlash = false
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{reactAppUri},
	}))
	router.Use(stripTrailingSlash())

	// Get all users
	articleRoutes := router.Group("/articles")
	articleRoutes.GET("", routes.GetAllArticles)
	articleRoutes.GET(":id", routes.GetArticle)
	articleRoutes.POST("", routes.CreateArticle)

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
