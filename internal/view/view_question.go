package view

import (
	"net/http"
	"strconv"

	"my-guora/conf"
	"my-guora/internal/h"
	"my-guora/internal/model"
	"my-guora/internal/service/rdbservice"

	"github.com/gin-gonic/gin"
)

// Question func
func Question(c *gin.Context) {
	var a model.Answer
	var q model.Question
	var _q model.Question
	var question model.Question
	var answers []model.Answer
	var answersCounts int
	var hotQuestions []model.Question
	var err error

	questionID, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(200, h.Response{
			Status:  500,
			Message: err.Error(),
		})
		return
	}

	q.ID = questionID
	a.QuestionID = questionID

	if question, err = q.Get(); err != nil {

		c.JSON(200, h.Response{
			Status:  404,
			Message: err.Error(),
		})

		return
	}

	if answers, err = a.GetOrderList(10, 0, "supporters_counts desc"); err != nil {

		c.JSON(200, h.Response{
			Status:  404,
			Message: err.Error(),
		})

		return
	}

	if answersCounts, err = a.GetCounts(); err != nil {

		c.JSON(200, h.Response{
			Status:  404,
			Message: err.Error(),
		})
		return
	}

	if hotQuestions, err = _q.GetOrderList(10, 0, "id desc"); err != nil {

		c.JSON(200, h.Response{
			Status:  404,
			Message: err.Error(),
		})
		return

	}

	PID, exist := c.Get("pid")
	if !exist {
		c.JSON(200, h.Response{Status: 404, Message: "Not exist"})
		c.Abort()
		return
	}

	ProfileID, ok := PID.(int)
	if !ok {
		c.JSON(200, h.Response{Status: 404, Message: "Not int"})
		c.Abort()
		return
	}

	if err = rdbservice.RedisWrapListSupported(answers, ProfileID); err != nil {

		c.JSON(200, h.Response{
			Status:  500,
			Message: err.Error(),
		})
		return

	}

	csrdata := map[string]interface{}{
		"question":      question,
		"answers":       answers,
		"answersCounts": answersCounts,
		"hotQuestions":  hotQuestions,
	}

	template := "question.html"
	data := map[string]interface{}{
		"lang":    conf.Config().Lang,
		"csrdata": csrdata,
	}

	c.HTML(http.StatusOK, template, data)

	return
}
