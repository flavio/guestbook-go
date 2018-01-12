package handlers

import (
	"net/http"
	"strconv"

	"github.com/flavio/guestbook-go/models"

	"github.com/go-redis/redis"
	"github.com/labstack/echo"
)

type GetMessagesResponse struct {
	Messages []models.Message `json:"messages"`
	Error    error            `json:"error"`
}

func GetMessages(db *redis.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		var response GetMessagesResponse
		messages, err := models.GetMessages(db)

		if err != nil {
			response.Error = err
			c.Logger().Error(err)
			return c.JSON(http.StatusInternalServerError, response)
		}

		response.Messages = messages

		return c.JSON(http.StatusOK, response)
	}
}

type PutMessageResponse struct {
	Index int64 `json:"index"`
	Error error `json:"error"`
}

func PutMessage(db *redis.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		var message models.Message

		// Map imcoming JSON body to the new Entry
		c.Bind(&message)

		// Add a task using our new model
		index, err := models.PutMessage(db, message.Data)
		// Return a JSON response if successful

		response := PutMessageResponse{
			Index: index,
			Error: err,
		}

		status := http.StatusCreated
		if err != nil {
			c.Logger().Error(err)
			status = http.StatusInternalServerError
		}

		return c.JSON(status, response)
	}
}

type DeleteMessageResponse struct {
	Error error `json:"error"`
}

func DeleteMessage(db *redis.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		response := DeleteMessageResponse{}
		status := http.StatusOK

		index, err := strconv.ParseInt(c.Param("index"), 10, 64)

		if err != nil {
			response.Error = err
		} else {
			err = models.DeleteMessage(db, index)
			if err != nil {
				status = http.StatusInternalServerError
				c.Logger().Error(err)
				response.Error = err
			}
		}

		return c.JSON(status, response)
	}
}
