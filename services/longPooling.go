package services

import (
	"go.uber.org/zap"
)

type LongPoolingService struct {
	sugar *zap.SugaredLogger
}

func NewLongPoolingService(sugar *zap.SugaredLogger) *LongPoolingService {
	return &LongPoolingService{
		sugar: sugar,
	}
}

func (a *LongPoolingService) GetLongPooling() (interface{}, error) {
	data := map[string]interface{}{
		"message": "success",
	}
	return data, nil
}
