package services

import (
	"fmt"
	"sync"
	"time"

	"go.uber.org/zap"
)

type Message struct {
	ID        uint64    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Message   string    `json:"msg"`
}

type LongPoolingService struct {
	sugar    *zap.SugaredLogger
	mutex    *sync.RWMutex
	messages map[uint64]Message
}

func NewLongPoolingService(
	sugar *zap.SugaredLogger,
	mutex *sync.RWMutex,
	messages map[uint64]Message,
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

	id := uint64(len(a.messages) + 1)

	a.messages[id] = Message{
		ID:        id,
		CreatedAt: time.Now(),
		Message:   message,
	}

	return a.messages[id], nil
}

func (a *LongPoolingService) GetMessages(id uint64) ([]Message, error) {
	a.mutex.RLock()
	defer a.mutex.RUnlock()

	fmt.Print(uint64(len(a.messages)), id)

	if uint64(len(a.messages)) > id {
		messages := make([]Message, uint64(len(a.messages))-id+1)
		// get messages from id to last
		for i := id; i <= uint64(len(a.messages)); i++ {
			messages[i-id] = a.messages[i]
		}
		return messages, nil
	} else {
		return []Message{}, nil
	}
}
