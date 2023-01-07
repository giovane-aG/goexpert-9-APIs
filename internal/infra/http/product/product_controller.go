package product_controller

import (
	"net/http"

	"github.com/giovane-aG/goexpert/9-APIs/internal/infra/database"
)

type ProductController struct {
	ProductDB database.ProductInterface
}

func NewProductController(productDB *database.Product) *ProductController {
	return &ProductController{
		ProductDB: productDB,
	}
}

func (p *ProductController) Create(w http.ResponseWriter, r *http.Request)   {}
func (p *ProductController) FindAll(w http.ResponseWriter, r *http.Request)  {}
func (p *ProductController) FindById(w http.ResponseWriter, r *http.Request) {}
func (p *ProductController) Update(w http.ResponseWriter, r *http.Request)   {}
func (p *ProductController) Delete(w http.ResponseWriter, r *http.Request)   {}
