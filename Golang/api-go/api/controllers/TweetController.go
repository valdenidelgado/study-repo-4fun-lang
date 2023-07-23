package controllers

import (
	"api/api/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type tweetController struct {
	tweets []entities.Tweet
}

func NewTweetController() *tweetController {
	return &tweetController{}
}

func (t *tweetController) FindAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, t.tweets)
}

func (t *tweetController) Create(ctx *gin.Context) {
	tweet := entities.NewTweet()

	if err := ctx.BindJSON(&tweet); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	t.tweets = append(t.tweets, *tweet)

	ctx.JSON(http.StatusCreated, tweet)
}

func (t *tweetController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	for index, tweet := range t.tweets {
		if tweet.ID == id {
			t.tweets = append(t.tweets[0:index], t.tweets[index+1:]...)
			ctx.JSON(http.StatusNoContent, gin.H{"message": "tweet deleted"})
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"error": "tweet not found"})
}
