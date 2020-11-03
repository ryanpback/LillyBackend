package handlers

import (
	"lillyAppBackend/models"
	"net/http"

	"github.com/gorilla/mux"
)

// Index retrieves images from GCS
func Index(w http.ResponseWriter, r *http.Request) {
	//
}

// Show gets one image by image name
func Show(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fileName := params["imageName"]
	imageURL, err := models.GetFileByName(fileName)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	response := Response{
		"status":   "Success",
		"imageURL": imageURL,
	}
	respondWithJSON(w, http.StatusOK, response)
}

// Create takes an image and stores it in GCS
func Create(w http.ResponseWriter, r *http.Request) {
	_, err := models.UploadFile(r)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(
		w, http.StatusCreated,
		Response{"status": "Created"},
	)
}
