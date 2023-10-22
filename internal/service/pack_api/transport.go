package pack_api

import "net/http"

type Transport interface {
	GetPacksNumber(w http.ResponseWriter, r *http.Request)
}
