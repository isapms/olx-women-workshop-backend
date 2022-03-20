package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"olx-women-workshop-2022-backend/models"

	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
)

// Response .
type Response struct {
	Err  string      `json:"error,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

func List(w http.ResponseWriter, r *http.Request) {
	var res Response

	adverts, err := models.List()

	for i := range adverts {
		adverts[i].Image = fmt.Sprintf("%s/%s", os.Getenv("IMAGE_PATH"), adverts[i].Image)
	}

	if err != nil {
		res.Err = err.Error()
	} else {
		res.Data = adverts
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var err error
	var res Response

	newAdvert := models.Advert{
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
	}

	if s, err := strconv.ParseFloat(r.FormValue("price"), 64); err == nil {
		newAdvert.Price = s
	}

	newAdvert.Image, err = getFormFile(r)
	if err != nil {
		res.Err = err.Error()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	}

	_, err = models.Create(newAdvert)
	if err != nil {
		res.Err = err.Error()
	} else {
		res.Data = "created"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	err := models.Delete(id)

	res := Response{}
	if err != nil {
		res.Err = err.Error()
	} else {
		res.Data = "deleted"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func getFormFile(r *http.Request) (string, error) {
	r.ParseMultipartForm(10 << 20) //10mb

	file, _, err := r.FormFile("ad_image")
	if err != nil {
		return "", err
	}
	defer file.Close()

	tempFile, err := ioutil.TempFile("static/images", "upload-*.png")
	if err != nil {
		return "", err
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	_, err = tempFile.Write(fileBytes)
	if err != nil {
		return "", err
	}

	return tempFile.Name(), nil
}
