package product_controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/giovane-aG/goexpert/9-APIs/internal/entity"
	"github.com/giovane-aG/goexpert/9-APIs/internal/infra/database"
	"github.com/giovane-aG/goexpert/9-APIs/internal/infra/http/product/dtos"
	"github.com/go-chi/chi/v5"
)

type ProductController struct {
	ProductDB database.ProductInterface
}

func NewProductController(productDB *database.Product) *ProductController {
	return &ProductController{
		ProductDB: productDB,
	}
}

// @Summary 	Creates a product
// @Tags		products
// @Accept		json
// @Produce		json
// @Param		product		body		dtos.CreateProductDto	true	"the product info"
// @Success		201			{object}	entity.Product
// @Failure		500			{object}	errors.Error
// @Failure		400			{object}	errors.Error
// @Router		/product	[post]
// @Security	ApiKeyAuth
func (p *ProductController) Create(w http.ResponseWriter, r *http.Request) {
	var parsedBody dtos.CreateProductDto
	err := json.NewDecoder(r.Body).Decode(&parsedBody)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = parsedBody.ValidateFields()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	newProduct, err := entity.NewProduct(parsedBody.Name, parsedBody.Price)
	err = newProduct.Validate()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = p.ProductDB.Create(newProduct)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(201)
	json.NewEncoder(w).Encode(newProduct)
	return
}

func (p *ProductController) FindAll(w http.ResponseWriter, r *http.Request) {
	var page, limit int
	var sort string

	page, _ = strconv.Atoi(chi.URLParam(r, "page"))
	limit, _ = strconv.Atoi(chi.URLParam(r, "limit"))
	sort = chi.URLParam(r, "sort")

	products, err := p.ProductDB.FindAll(page, limit, sort)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(products)
}

func (p *ProductController) FindById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	product, err := p.ProductDB.FindById(id)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(product)
}

func (p *ProductController) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var parsedBody dtos.UpdateProductDto
	err := json.NewDecoder(r.Body).Decode(&parsedBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = parsedBody.ValidateFields()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	foundProduct, err := p.ProductDB.FindById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if foundProduct == nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"message": "no product with that id was found"})
		return
	}

	updatedProduct := entity.Product{}
	if parsedBody.Name != "" {
		updatedProduct.Name = parsedBody.Name
	}

	if parsedBody.Price > 0 {
		updatedProduct.Price = parsedBody.Price
	}

	p.ProductDB.Update(&updatedProduct)
	json.NewEncoder(w).Encode(updatedProduct)

}
func (p *ProductController) Delete(w http.ResponseWriter, r *http.Request) {}
