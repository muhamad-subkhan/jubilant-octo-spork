package handlers

import (
	dto "BE-foodways/dto/result"
	usersdto "BE-foodways/dto/users"
	"BE-foodways/models"
	"BE-foodways/repositories"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)


type handlerUser struct {
	UserRepository repositories.UserRepositories
}

func HandlerUser(UserRepository repositories.UserRepositories) *handlerUser {
	return &handlerUser{UserRepository}
}

  func (h *handlerUser) FindUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users, err := h.UserRepository.FindUsers() // menjalankan query kedatabase
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: users, Status: "Success"}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerUser) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	user, err := h.UserRepository.GetUser(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(user), Status: "Success"}
	json.NewEncoder(w).Encode(response)
}
	

func (h *handlerUser) CreateUser(w http.ResponseWriter, r *http.Request) {
	 w.Header().Set("Content-type", "application/json")

	 request := new(usersdto.CreateUserRequest)
	 if err := json.NewDecoder(r.Body).Decode(&request)
	 err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	 }

	 validation := validator.New()
	 err := validation.Struct(request)
	 if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	 }

	 user := models.User{
		ID: request.ID,
		Name: request.Name,
		Email: request.Email,
		Phone: request.Phone,
		Location: request.Location,
		Image: request.Image,
		Role: request.Role,
	 }

	 data, err := h.UserRepository.CreateUser(user)
	 if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	 }

	 w.WriteHeader(http.StatusOK)
	 response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(data), Status: "Success"}
	 json.NewEncoder(w).Encode(response)


}


func (h *handlerUser) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//uploadfile
	dataContex := r.Context().Value("dataFile") 
	filename := dataContex.(string)

	request := usersdto.UpdateUserRequest{
		Name: r.FormValue("fullName"),
		Email:    r.FormValue("email"),
		Phone:    r.FormValue("phone"),
		Location: r.FormValue("location"),
		Gender: r.FormValue("gender"),
		Image: r.FormValue("image"),

	}

	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	resp, err := cld.Upload.Upload(ctx, filename, uploader.UploadParams{Folder: "waysfood"})

	if err != nil {
		fmt.Println(err.Error())
	}
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	user, err := h.UserRepository.GetUser(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if request.Name != "" {
		user.Name = request.Name

	}

	if request.Email != "" {
		user.Email = request.Email
	}

	if request.Phone != "" {
		user.Phone = request.Phone
	}

	if request.Location != "" {
		user.Location = request.Location
	}

	if filename != "" {
		user.Image = resp.SecureURL
	}

	if request.Gender != "" {
		user.Gender = request.Gender
	}

	data, err := h.UserRepository.UpdateUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data, Status: "Success"}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerUser) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	user, err := h.UserRepository.GetUser(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.UserRepository.DeleteUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(data), Status: "Success"}
	json.NewEncoder(w).Encode(response)
}


func convertResponse(u models.User) usersdto.UserResponse {
	return usersdto.UserResponse{
		ID:       u.ID,
		Name:     u.Name,
		Email:    u.Email,
		Phone: u.Phone,
		Gender: u.Gender,
		Location: u.Location,
		Image: u.Image,
		Role: u.Role,
	}
}

func UpdateRespone(u models.User) usersdto.UpdateRespone {
	return usersdto.UpdateRespone{
		ID:       u.ID,
		Name:     u.Name,
		Email:    u.Email,
		Phone: u.Phone,
		Gender: u.Gender,
		Location: u.Location,
		Image: u.Image,
		Role: u.Role,
	}
}
  