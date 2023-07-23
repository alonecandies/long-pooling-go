package inject

import (
	"github.com/alonecandies/long-pooling-go/services"
)

func (a *Injector) ProvideLongPoolingService() *services.LongPoolingService {
	return services.NewLongPoolingService(
		a.ProvideLogger(),
	)
}
