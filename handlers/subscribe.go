package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	dto "wayshub-server/dto/result"
	"wayshub-server/models"
	"wayshub-server/repositories"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerSubscribe struct {
	SubscribeRepository repositories.SubscribeRepository
}

func HandlerSubscribe(SubscribeRepository repositories.SubscribeRepository) *handlerSubscribe {
	return &handlerSubscribe{SubscribeRepository}
}

func (h *handlerSubscribe) Subscription(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// channelInfo := r.Context().Value("channelInfo").(jwt.MapClaims)
	// channelID := int(channelInfo["id"].(float64))

	subscription, err := h.SubscribeRepository.Subscription()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: subscription}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerSubscribe) Subscribe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	channelInfo := r.Context().Value("channelInfo").(jwt.MapClaims)
	channelID := int(channelInfo["id"].(float64))

	subscribe := models.Subscribe{
		ChannelID: channelID,
	}

	subscribe, err := h.SubscribeRepository.Subscribe(subscribe)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	subscribe, _ = h.SubscribeRepository.GetSubscription(channelID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: subscribe}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerSubscribe) Unsubscribe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	channelInfo := r.Context().Value("channelInfo").(jwt.MapClaims)
	channelID := int(channelInfo["id"].(float64))

	if channelID != id {
		w.WriteHeader(http.StatusUnauthorized)
		response := dto.ErrorResult{Code: http.StatusUnauthorized, Message: "PLease Login First!"}
		json.NewEncoder(w).Encode(response)
		return
	}

	unsubscribe, err := h.SubscribeRepository.GetSubscription(channelID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.SubscribeRepository.Unsubscribe(unsubscribe)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: data}
	json.NewEncoder(w).Encode(response)
}
