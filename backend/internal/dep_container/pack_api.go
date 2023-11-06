package dep_container

import (
	"github.com/sarulabs/di"
	"shippingPacks/internal/config"
	"shippingPacks/internal/service/pack_api/transport/gorilla_mux"
	"shippingPacks/internal/service/pack_api/usecase"
)

const packApiDefName = "pack-api"

// RegisterPackApiService registers PackApiService dependency.
func RegisterPackApiService(builder *di.Builder) error {
	return builder.Add(di.Def{
		Name: packApiDefName,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get(configDefName).(*config.Config)
			return gorilla_mux.New(usecase.New(cfg.PackSize)), nil
		},
	})
}
