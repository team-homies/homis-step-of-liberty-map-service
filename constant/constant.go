package constant

import (
	"main/constant/path/core"
	"sync"
)

var (
	instance *core.InternalApi
	once     sync.Once
)

func LocationPath() *core.InternalApi {
	once.Do(func() {
		instance = &core.InternalApi{
			Location: core.LocationPath{
				FindEvent: "/histories",
			},
		}
	})

	return instance
}
