package services

import (
	"sync"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type Message struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Message   string    `json:"msg"`
}

type LongPoolingService struct {
	sugar    *zap.SugaredLogger
	mutex    *sync.RWMutex
	messages map[uuid.UUID]Message
}

func NewLongPoolingService(
	sugar *zap.SugaredLogger,
	mutex *sync.RWMutex,
	messages map[uuid.UUID]Message,
) *LongPoolingService {
	return &LongPoolingService{
		sugar:    sugar,
		mutex:    mutex,
		messages: messages,
	}
}

func (a *LongPoolingService) CreateMessage(message string) (Message, error) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	id := uuid.New()
	a.messages[id] = Message{
		ID:        id,
		CreatedAt: time.Now(),
		Message:   message,
	}

	return a.messages[id], nil
}

func (a *LongPoolingService) GetMessages(after *uuid.UUID) ([]Message, error) {
	if after == nil {
		return []Message{}, nil
	}

	a.mutex.RLock()
	defer a.mutex.RUnlock()

	messages := make([]Message, len(a.messages))
	for _, message := range a.messages {
		messages = append(messages, message)
	}

	return messages, nil
}
