package car

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/pranayyb/DriveThrough/models"
	"github.com/pranayyb/DriveThrough/service"
	"io"
	"log"
	"net/http"
)

type CarHandler struct {
	service service.CarServiceInterface
}

func NewCarHandler(service service.CarServiceInterface) *CarHandler {
	return &CarHandler{
		service: service,
	}
}

func (h *CarHandler) GetCarById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id := vars["id"]
	res, err := h.service.GetCarById(id, ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("error: ", err)
		return
	}
	body, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("error: ", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(body)
	if err != nil {
		log.Println("error writing response")
	}
}

func (h *CarHandler) GetCarByBrand(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	brand := r.URL.Query().Get("brand")
	isEngine := r.URL.Query().Get("isEngine") == "true"

	res, err := h.service.GetCarByBrand(brand, ctx, isEngine)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("error: ", err)
		return
	}
	body, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("error writing response")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(body)
	if err != nil {
		log.Println("error writing response")
	}
}

func (h *CarHandler) CreateCar(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("error: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var carReq models.CarRequest
	err = json.Unmarshal(body, &carReq)
	if err != nil {
		log.Println("error while un-marshalling request: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	createdCar, err := h.service.CreateCar(&carReq, ctx)
	if err != nil {
		log.Println("error while creating car: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseBody, err := json.Marshal(createdCar)
	if err != nil {
		log.Println("error while marshalling: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	_, err = w.Write(responseBody)
	if err != nil {
		log.Println("error writing response")
	}
}

func (h *CarHandler) UpdateCar(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	params := mux.Vars(r)
	id := params["id"]
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Err Reading Request body:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var carReq models.CarRequest
	err = json.Unmarshal(body, &carReq)
	if err != nil {
		log.Println("Error while Un-marshalling Request body", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	updatedCar, err := h.service.UpdateCar(id, &carReq, ctx)
	if err != nil {
		log.Println("Error while Updating the Car :", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	responseBody, err := json.Marshal(updatedCar)
	if err != nil {
		log.Println("error while marshalling: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(responseBody)
	if err != nil {
		log.Println("error writing response")
	}
}

func (h *CarHandler) DeleteCar(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	params := mux.Vars(r)
	id := params["id"]
	deletedCar, err := h.service.DeleteCar(id, ctx)
	if err != nil {
		log.Println("Error while deleting the Car :", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	responseBody, err := json.Marshal(deletedCar)
	if err != nil {
		log.Println("error while marshalling: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(responseBody)
	if err != nil {
		log.Println("error writing response")
	}

}
