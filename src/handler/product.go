package handler

import (
	"../pkg"
	"../product"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type productHandler struct {
	service product.Service
}

func NewProductHandler(repo product.PRepository) *productHandler {
	return &productHandler{service: product.NewProductService(repo)}
}

// ListProducts godoc
// @Summary List all products
// @Description get products
// @Tags Product
// @Accept  json
// @Produce  json
// @Success 200 {array} product.Product
// @Security BearerAuth
// @Router /v1/products [get]
func (ph productHandler) HandleGetAll(w http.ResponseWriter, r *http.Request) {
	u, err := ph.service.FindAll(r.Context())
	if err != nil {
		pkg.Fail(err).ToJSON(w)
		return
	}

	pkg.OK("", u).ToJSON(w)
	return
}

// ListProduct godoc
// @Summary Find product by ID
// @Description get product by ID
// @Tags Product
// @Accept  json
// @Produce  json
// @Param id path int true "Product ID"
// @Success 200 {object} product.Product
// @Router  /v1/products/{id} [get]
func (ph productHandler) HandleGetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		pkg.Fail(err).ToJSON(w)
		return
	}

	u, err := ph.service.FindByID(r.Context(), uint(id))
	if err != nil {
		pkg.Fail(err).ToJSON(w)
		return
	}

	pkg.OK("", u).ToJSON(w)
	return
}

// CreateProduct godoc
// @Summary Add a new product
// @Description create a new product
// @Tags Product
// @Accept  json
// @Produce  json
// @Param  product body product.Product true "Create product"
// @Success 200 {object} product.Product
// @Router /v1/products [post]
func (ph productHandler) HandleStore(w http.ResponseWriter, r *http.Request) {
	productModel := product.Product{}
	if err := json.NewDecoder(r.Body).Decode(&productModel); err != nil {
		pkg.Fail(err).ToJSON(w)
		return
	}

	u, err := ph.service.Store(r.Context(), productModel)
	if err != nil {
		pkg.Fail(err).ToJSON(w)
		return
	}

	pkg.OK("Product created successfully", u).ToJSON(w)
	return
}

// UpdateProduct godoc
// @Summary Update an existing product
// @Description update an existing product by ID
// @ID int
// @Tags Product
// @Accept  json
// @Produce  json
// @Param id path int true "Product ID"
// @Param user body product.Product true "Update product"
// @Success 200 {object} product.Product
// @Router /v1/products/{id} [put]
func (ph productHandler) HandleUpdate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		pkg.Fail(err).ToJSON(w)
		return
	}

	productModel := product.Product{}
	if err := json.NewDecoder(r.Body).Decode(&productModel); err != nil {
		pkg.Fail(err).ToJSON(w)
		return
	}

	u, err := ph.service.Update(r.Context(), uint(id), productModel)
	if err != nil {
		pkg.Fail(err).ToJSON(w)
		return
	}

	pkg.OK("Product updated successfully", u).ToJSON(w)
	return
}
