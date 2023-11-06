package dep_container

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sarulabs/di"
	"go.uber.org/zap"

	"shippingPacks/internal/config"
	"shippingPacks/internal/service/pack_api"
)

const (
	httpServerDefName = "http-server"
)

// RegisterHTTPServer registers HTTP Server dependency.
func RegisterHTTPServer(builder *di.Builder) error {
	return builder.Add(di.Def{
		Name: httpServerDefName,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get(configDefName).(*config.Config)
			packApiTransport := ctn.Get(packApiDefName).(pack_api.Transport)

			r := mux.NewRouter()
			headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
			methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
			origins := handlers.AllowedOrigins([]string{"http://localhost:3000"})

			r.HandleFunc("/api/v1/get-packs-number/{itemsOrdered}", packApiTransport.GetPacksNumber).Methods("GET")

			zap.L().Info("started http server",
				zap.String("address", fmt.Sprintf("http://localhost:%s/", cfg.Port)))

			zap.S().Fatal("", zap.Error(http.ListenAndServe(":"+cfg.Port, handlers.CORS(headers, methods, origins)(r))))

			return r, nil
		},
	})
}

// RunHTTPServer runs HTTP Server dependency.
func (c Container) RunHTTPServer() {
	c.container.Get(httpServerDefName)
}
