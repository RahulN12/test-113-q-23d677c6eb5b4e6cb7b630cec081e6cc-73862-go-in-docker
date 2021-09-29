package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Failure struct {
	Status string `json:"status"`
	Reason string `json:"reason"`
}

func FetchQuiz(c *gin.Context) {

	id := c.Param("quiz_id")
	if len(id) == 0 {
		msg := Failure{"failure", "invalid input"}
		c.IndentedJSON(400, msg)
		return
	}

	responseQ, fail := GetQuiz(id)
	if responseQ.Id == 0 {
		if strings.Contains(fail, "no rows") {
			c.PureJSON(404, gin.H{})
			return
		}
		msg := Failure{"failure", fail}
		c.IndentedJSON(400, msg)
		return
	}

	c.IndentedJSON(200, responseQ)
}

func AddQuiz(c *gin.Context) {

	var quiz Quiz
	if err := c.BindJSON(&quiz); err != nil {
		fmt.Println(err)
		msg := Failure{"failure", err.Error()}
		c.IndentedJSON(400, msg)
		return
	}

	if len(quiz.Name) < 1 {
		msg := Failure{"failure", "invalid input name"}
		c.IndentedJSON(400, msg)
		return
	}

	if len(quiz.Desc) < 1 {
		msg := Failure{"failure", "invalid input description"}
		c.IndentedJSON(400, msg)
		return
	}

	responseQ := CreateQuiz(quiz)
	c.IndentedJSON(201, responseQ)
}

func FetchQuestion(c *gin.Context) {

	id := c.Param("question_id")
	if len(id) == 0 {
		msg := Failure{"failure", "invalid input"}
		c.IndentedJSON(400, msg)
		return
	}

	responseQ, fail := GetQuestion(id)
	if responseQ.Id == 0 {
		if strings.Contains(fail, "no rows") {
			c.PureJSON(404, gin.H{})
			return
		}
		msg := Failure{"failure", fail}
		c.IndentedJSON(400, msg)
		return
	}

	c.IndentedJSON(200, responseQ)
}

func NewQuestion(c *gin.Context) {

	var quiz Question
	if err := c.BindJSON(&quiz); err != nil {
		msg := Failure{"failure", err.Error()}
		c.IndentedJSON(400, msg)
		return
	}

	if len(quiz.Name) < 1 {
		msg := Failure{"failure", "invalid input name"}
		c.IndentedJSON(400, msg)
		return
	}

	if len(quiz.Options) < 1 {
		msg := Failure{"failure", "invalid input options"}
		c.IndentedJSON(400, msg)
		return
	}

	if quiz.CorrectOption <= 0 {
		msg := Failure{"failure", "invalid input correct"}
		c.IndentedJSON(400, msg)
		return
	}

	if quiz.Quiz <= 0 {
		msg := Failure{"failure", "invalid input quiz"}
		c.IndentedJSON(400, msg)
		return
	}

	if quiz.Points <= 0 {
		msg := Failure{"failure", "invalid input points"}
		c.IndentedJSON(400, msg)
		return
	}

	q, str := GetQuiz(strconv.Itoa(quiz.Quiz))
	if q.Id == 0 || len(str) > 0 {
		msg := Failure{"failure", "invalid input quiz id"}
		c.IndentedJSON(400, msg)
		return
	}

	res, id := EnterQuestion(quiz)
	if !res {
		msg := Failure{"failure", "error in insert que"}
		c.IndentedJSON(400, msg)
		return
	}

	quiz.Id = id
	c.IndentedJSON(201, quiz)
}

func AllQuestions(c *gin.Context) {

	id := c.Param("quiz_id")
	if len(id) == 0 {
		msg := Failure{"failure", "invalid input"}
		c.IndentedJSON(400, msg)
		return
	}

	responseQ, fail := GetAllQuestion(id)
	if len(fail) > 0 {
		if strings.Contains(fail, "no rows") {
			c.PureJSON(404, gin.H{})
			return
		}
		msg := Failure{"failure", fail}
		c.IndentedJSON(400, msg)
		return
	}

	if len(responseQ.Name) == 0 {
		c.PureJSON(404, gin.H{})
		return
	}

	c.IndentedJSON(200, responseQ)
}
