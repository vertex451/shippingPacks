package gorilla_mux

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go.uber.org/zap"

	"shippingPacks/internal/service/pack_api"
)

type Transport struct {
	useCase pack_api.UseCase
}

func New(useCase pack_api.UseCase) *Transport {
	return &Transport{
		useCase: useCase,
	}
}

func (t Transport) GetPacksNumber(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemsOrdered, err := strconv.Atoi(vars["itemsOrdered"])
	if err != nil {
		http.Error(w, "failed to convert itemsOrdered to number", http.StatusBadRequest)

		return
	}

	if itemsOrdered < 0 {
		http.Error(w, "itemsOrdered should not be negative", http.StatusBadRequest)

		return
	}

	jsonData, err := json.Marshal(t.useCase.CalculatePacksNumber(itemsOrdered))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonData)
	if err != nil {
		zap.L().Error("failed to write to json", zap.Error(err))
	}
}
