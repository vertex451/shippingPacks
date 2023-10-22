package dep_container

import (
	"github.com/sarulabs/di"

	"shippingPacks/internal/config"
)

const configDefName = "config"

// RegisterConfig registers Config dependency.
func RegisterConfig(builder *di.Builder) error {
	return builder.Add(di.Def{
		Name: configDefName,
		Build: func(ctn di.Container) (interface{}, error) {
			return config.LoadConfig(".env")
		},
	})
}
