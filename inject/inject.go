package inject

import (
	"sync"

	"go.uber.org/zap"
)

type Injector struct {
	sugar *zap.SugaredLogger
}

func NewInjector() *Injector {
	a := &Injector{}
	return a
}

func (a *Injector) ProvideLogger() *zap.SugaredLogger {
	return a.sugar
}

func (a *Injector) ProvideMutex() *sync.RWMutex {
	return &sync.RWMutex{}
}
