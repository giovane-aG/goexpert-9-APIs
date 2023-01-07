package product_controller

import (
	"encoding/json"
	"net/http"

	"github.com/giovane-aG/goexpert/9-APIs/internal/entity"
	"github.com/giovane-aG/goexpert/9-APIs/internal/infra/database"
	"github.com/giovane-aG/goexpert/9-APIs/internal/infra/http/product/dtos"
)

type ProductController struct {
	ProductDB database.ProductInterface
}

func NewProductController(productDB *database.Product) *ProductController {
	return &ProductController{
		ProductDB: productDB,
	}
}

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
func (p *ProductController) FindAll(w http.ResponseWriter, r *http.Request)  {}
func (p *ProductController) FindById(w http.ResponseWriter, r *http.Request) {}
func (p *ProductController) Update(w http.ResponseWriter, r *http.Request)   {}
func (p *ProductController) Delete(w http.ResponseWriter, r *http.Request)   {}
