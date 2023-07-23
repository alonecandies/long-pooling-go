package api

import (
	"net/http"

	"github.com/alonecandies/long-pooling-go/services"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

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

func (a *LongPoolingApi) GetLongPooling(c *gin.Context) {
	res, err := a.longPoolingService.GetLongPooling()
	if err != nil {
		a.sugar.Error("error get long pooling: ", err)
		return
	}
	c.JSON(http.StatusOK, res)
}
