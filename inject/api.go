package inject

import "github.com/alonecandies/long-pooling-go/api"

func (a *Injector) ProvideLongPoolingApi() *api.LongPoolingApi {
	return api.NewLongPoolingApi(a.ProvideLogger(), *a.ProvideLongPoolingService())
}
