package routes

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ernestechie/go-blog/db"
	"github.com/ernestechie/go-blog/models"
	"github.com/ernestechie/go-blog/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var articlesColl *mongo.Collection

// Initialize mongodb client
func init()  {
	c := db.ConnectDB()
	articlesColl = c.Database("db").Collection("articles")
}

func GetAllArticles(ctx *gin.Context) {
	sort := bson.D{{Key: "createdAt", Value: 1}}

	cursor, err := articlesColl.Find(context.TODO(), bson.M{}, options.Find().SetSort(sort))
	if err != nil {
		log.Println("usersColl.Find \n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
			"success": false,
	 		"message": "Error getting articles",
		})
		return
	}

	articles := []models.ArticleModel{}

	if err = cursor.All(context.TODO(), &articles); err != nil {
		log.Println("cursor.All \n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
			"success": false,
	 		"message": "Error getting articles",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": gin.H{
		"articles": articles,
	}, "success": true, "message": "Articles retrieved successfully"})
}


func CreateArticle(ctx *gin.Context) {
	article := models.ArticleModel{}
	article.CreatedAt = time.Now()
	article.UpdatedAt = time.Now()

	// Parse and validate request body using utils
	if errs := utils.ParseAndValidate(ctx, &article); len(errs) > 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": errs,
			"success": false, 
			"message": "Error processing request",
		})
		return
	}

	// article.Author = strings.ToLower(user.Email)

	// Create new user if the user does not exist.
	result, err := articlesColl.InsertOne(context.TODO(), article)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
			"success": false, 
			"message": "Error creating article",
		})
		return
	}

	articleId := result.InsertedID.(bson.ObjectID)
	article.ID = articleId

	ctx.JSON(http.StatusOK, gin.H{"data": gin.H{
		"article": article,
	}, "success": true, "message": "Article retrieved successfully"})
}


func GetArticle(ctx *gin.Context)  {
	articleId := ctx.Params.ByName("id")
	article := models.ArticleModel{}

	articleObjectId, userIdErr := bson.ObjectIDFromHex(articleId);
	if userIdErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"data": nil, "success": false, "message": fmt.Sprintf("Invalid article ID, %v", articleId),})
	return
	}

	fmt.Println(articleObjectId)
	filter := bson.M{"_id": articleObjectId}
	err := articlesColl.FindOne(context.TODO(), filter).Decode(&article)
		if err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusNotFound, gin.H{"data": nil, "success": false, "message": "Article not found"})
		return
		}

	ctx.JSON(http.StatusOK, gin.H{"data": gin.H{
		"article": article,
	}, "success": true, "message": "Article retrieved successfully"})
}
