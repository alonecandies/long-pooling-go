package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/alonecandies/long-pooling-go/services"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Message struct {
	ID        uint64    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Message   string    `json:"msg"`
}

type MessageInput struct {
	Body string `json:"body"`
}

type LongPoolingApi struct {
	sugar              *zap.SugaredLogger
	longPoolingService services.LongPoolingService
}

func NewLongPoolingApi(sugar *zap.SugaredLogger, longPoolingService services.LongPoolingService) *LongPoolingApi {
	return &LongPoolingApi{
		sugar:              sugar,
		longPoolingService: longPoolingService,
	}
}

func (a *LongPoolingApi) GetMessages(c *gin.Context) {
	var after uint64

	id := c.Query("id")

	if id == "" {
		after = 0
	} else {
		tmpAfter, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			c.Error(err)
			return
		}
		after = tmpAfter
	}

	for i := 0; i < 10; i++ {
		messages, err := a.longPoolingService.GetMessages(after)
		if err != nil {
			a.sugar.Error("error get long pooling: ", err)
			return
		}
		if len(messages) > 0 {
			c.JSON(http.StatusOK, messages)
			return
		}
		time.Sleep(1 * time.Second)
	}

	messages := []Message{}

	c.JSON(http.StatusOK, messages)
}

func (a *LongPoolingApi) CreateMessage(c *gin.Context) {
	var input MessageInput

	bodyData := c.Request.Body

	body, err := io.ReadAll(io.LimitReader(bodyData, 1048576))
	if err != nil {
		c.Error(err)
		return

	}
	if err := bodyData.Close(); err != nil {
		c.Error(err)
		return
	}

	if err := json.Unmarshal(body, &input); err != nil {
		c.Error(err)
		return
	}

	message, err := a.longPoolingService.CreateMessage(input.Body)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, message)
}
