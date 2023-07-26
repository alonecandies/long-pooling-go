package inject

import (
	"github.com/alonecandies/long-pooling-go/services"
	"github.com/google/uuid"
)

func (a *Injector) ProvideLongPoolingService() *services.LongPoolingService {
	return services.NewLongPoolingService(
		a.ProvideLogger(),
		a.ProvideMutex(),
		map[uuid.UUID]services.Message{},
	)
}
