package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	r.GET("/api/quiz/:quiz_id", FetchQuiz)
	r.POST("/api/quiz/", AddQuiz)

	r.GET("/api/questions/:question_id", FetchQuestion)
	r.POST("/api/questions/", NewQuestion)

	r.GET("/api/quiz-questions/:quiz_id", AllQuestions)

	r.Run(":8080")
}
