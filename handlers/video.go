package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"time"
	dto "wayshub-server/dto/result"
	videodto "wayshub-server/dto/video"
	"wayshub-server/models"
	"wayshub-server/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerVideo struct {
	VideoRepository repositories.VideoRepository
}

func HandlerVideo(VideoRepository repositories.VideoRepository) *handlerVideo {
	return &handlerVideo{VideoRepository}
}

func (h *handlerVideo) FindVideos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	videos, err := h.VideoRepository.FindVideos()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	for i, p := range videos {
		videos[i].Thumbnail = os.Getenv("PATH_FILE") + p.Thumbnail
	}

	for i, p := range videos {
		videos[i].Video = os.Getenv("PATH_FILE") + p.Video
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: videos}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerVideo) GetVideo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var video models.Video
	video, err := h.VideoRepository.GetVideo(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	video.Thumbnail = os.Getenv("PATH_FILE") + video.Thumbnail
	video.Video = os.Getenv("PATH_FILE") + video.Video

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: video}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerVideo) CreateVideo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	channelInfo := r.Context().Value("channelInfo").(jwt.MapClaims)
	channelID := int(channelInfo["id"].(float64))

	dataThumbnail := r.Context().Value("dataThumbnail")
	fileThumbnail := dataThumbnail.(string)

	dataVideo := r.Context().Value("dataVideo")
	fileVideo := dataVideo.(string)

	request := videodto.CreateVideoRequest{
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	video := models.Video{
		Title:       request.Title,
		Thumbnail:   fileThumbnail,
		Description: request.Description,
		Video:       fileVideo,
		CreatedAt:   time.Now(),
		ChannelID:   channelID,
	}

	video, err = h.VideoRepository.CreateVideo(video)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	video, _ = h.VideoRepository.GetVideo(video.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: video}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerVideo) UpdateVideo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	channelInfo := r.Context().Value("channelInfo").(jwt.MapClaims)
	channelID := int(channelInfo["id"].(float64))

	if channelID != id {
		w.WriteHeader(http.StatusUnauthorized)
		response := dto.ErrorResult{Code: http.StatusUnauthorized, Message: "Can't update channel!"}
		json.NewEncoder(w).Encode(response)
		return
	}

	dataThumbnail := r.Context().Value("dataThumbnail")
	fileThumbnail := dataThumbnail.(string)

	dataVideo := r.Context().Value("dataVideo")
	fileVideo := dataVideo.(string)

	request := videodto.UpdateVideoRequest{
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
	}

	video, err := h.VideoRepository.GetVideo(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if request.Title != "" {
		video.Title = request.Title
	}

	if request.Description != "" {
		video.Description = request.Description
	}

	if fileThumbnail != "false" {
		video.Thumbnail = fileThumbnail
	}

	if fileVideo != "false" {
		video.Video = fileVideo
	}

	data, err := h.VideoRepository.UpdateVideo(video)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	video.Thumbnail = os.Getenv("PATH_FILE") + video.Thumbnail
	video.Video = os.Getenv("PATH_FILE") + video.Video

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: data}
	json.NewEncoder(w).Encode(response)

}

func (h *handlerVideo) DeleteVideo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	channelInfo := r.Context().Value("channelInfo").(jwt.MapClaims)
	channelID := int(channelInfo["id"].(float64))

	video, err := h.VideoRepository.FindVideosByChannelId(channelID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.VideoRepository.DeleteVideo(video)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: DeleteVideoResponse(data)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerVideo) FindVideosByChannelId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	channelInfo := r.Context().Value("channelInfo").(jwt.MapClaims)
	channelID := int(channelInfo["id"].(float64))

	videos, err := h.VideoRepository.FindVideosByChannelId(channelID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	videos.Thumbnail = os.Getenv("PATH_FILE") + videos.Thumbnail
	videos.Video = os.Getenv("PATH_FILE") + videos.Video

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: videos}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerVideo) FindMyVideos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	channelInfo := r.Context().Value("channelInfo").(jwt.MapClaims)
	channelID := int(channelInfo["id"].(float64))

	videos, err := h.VideoRepository.FindMyVideos(channelID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	for i, p := range videos {
		videos[i].Thumbnail = os.Getenv("PATH_FILE") + p.Thumbnail
	}

	for i, p := range videos {
		videos[i].Video = os.Getenv("PATH_FILE") + p.Video
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: videos}
	json.NewEncoder(w).Encode(response)
}

func DeleteVideoResponse(u models.Video) videodto.DeleteResponse {
	return videodto.DeleteResponse{
		ID: u.ID,
	}
}
