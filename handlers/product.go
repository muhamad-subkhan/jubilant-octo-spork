package handlers

import (
	productdto "BE-foodways/dto/product"
	dto "BE-foodways/dto/result"
	"BE-foodways/models"
	"BE-foodways/repositories"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"context"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerProduct struct {
	ProductRepositories repositories.ProductRepositories
}

var path_file = "http://localhost:5000/uploads/"

func HandlerProduct(ProductRepository repositories.ProductRepositories) *handlerProduct {
	return &handlerProduct{ProductRepository}
}

func (h *handlerProduct) FindProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	products, err := h.ProductRepositories.FindProducts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Create Embed Path File on Image property here ...
	for i, p := range products {
		imagePath := os.Getenv("PATH_FILE") + p.Image
		products[i].Image = imagePath
	}
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: products, Status: "Success"}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerProduct) GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var product models.Product
	product, err := h.ProductRepositories.GetProduct(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Create Embed Path File on Image property here ...
	// product.Image = path_file + product.Image
	product.Image = os.Getenv("PATH_FILE") + product.Image

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseProduct(product), Status: "Success"}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerProduct) CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// get data user token
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	// get image filepath
	dataContex := r.Context().Value("dataFile")
	filepath := dataContex.(string)

	price, _ := strconv.Atoi(r.FormValue("price"))
	// qty, _ := strconv.Atoi(r.FormValue("qty"))
	// category_id, _ := strconv.Atoi(r.FormValue("category_id"))
	request := productdto.ProductRequest{
		Title: r.FormValue("title"),
		Image: r.FormValue("image"),
		Price: price,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	resp, err := cld.Upload.Upload(ctx, filepath, uploader.UploadParams{Folder: "waysfood"})

	if err != nil {
		fmt.Println(err.Error())
	}

	product := models.Product{
		Title:  request.Title,
		Price:  request.Price,
		Image:  resp.SecureURL, // add this code
		UserID: userId,
	}

	product, err = h.ProductRepositories.CreateProduct(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	product, _ = h.ProductRepositories.GetProduct(product.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: product, Status: "Success"}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerProduct) ProductDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var product models.Product
	product, err := h.ProductRepositories.GetProduct(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Create Embed Path File on Image property here ...
	product.Image = path_file + product.Image

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: ResponseProduct(product)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerProduct) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(productdto.ProductRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	product, err := h.ProductRepositories.GetProduct(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if request.Title != "" {
		product.Title = request.Title
	}

	data, err := h.ProductRepositories.UpdateProduct(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: ResponseProduct(data), Status: "Success"}
	json.NewEncoder(w).Encode(response)

}

func (h *handlerProduct) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	product, err := h.ProductRepositories.GetProduct(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.ProductRepositories.DeleteProduct(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: DeleteProduct(data), Status: "Success"}
	json.NewEncoder(w).Encode(response)
}

func convertResponseProduct(u models.Product) models.ProductResponse {
	return models.ProductResponse{
		ID:    u.ID,
		Title: u.Title,
		Price: u.Price,
		Image: u.Image,
	}
}
func ResponseProduct(u models.Product) models.ProductDetail {
	return models.ProductDetail{
		ID:    u.ID,
		Title: u.Title,
		Price: u.Price,
		Image: u.Image,
		User:  u.User,
	}
}

func DeleteProduct(u models.Product) models.ProductDelete {
	return models.ProductDelete{
		ID: u.ID,
	}
}
