package handler

import (
	"net/http"
	"strconv"

	"app_desafio/internal"
	"app_desafio/platform/web/request"
	"app_desafio/platform/web/response"

	"github.com/go-chi/chi/v5"
)

// NewProductsDefault returns a new ProductsDefault
func NewProductsDefault(sv internal.ServiceProduct) *ProductsDefault {
	return &ProductsDefault{sv: sv}
}

// ProductsDefault is a struct that returns the product handlers
type ProductsDefault struct {
	// sv is the product's service
	sv internal.ServiceProduct
}

// ProductJSON is a struct that represents a product in JSON format
type ProductJSON struct {
	Id          int     `json:"id"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

// ProductSaleCountJSON is a struct that represents the sale count of a product in JSON format
type ProductSaleCountJSON struct {
	Description string `json:"description"`
	SaleCount   int    `json:"sale_count"`
}

// GetAll returns all products
func (h *ProductsDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		p, err := h.sv.FindAll()
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "error getting products")
			return
		}

		// response
		// - serialize
		pJSON := make([]ProductJSON, len(p))
		for ix, v := range p {
			pJSON[ix] = ProductJSON{
				Id:          v.Id,
				Description: v.Description,
				Price:       v.Price,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "products found",
			"data":    pJSON,
		})
	}
}

// RequestBodyProduct is a struct that represents the request body for a product
type RequestBodyProduct struct {
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

// Create creates a new product
func (h *ProductsDefault) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - body
		var reqBody RequestBodyProduct
		err := request.JSON(r, &reqBody)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "error parsing request body")
			return
		}

		// process
		// - deserialize
		p := internal.Product{
			ProductAttributes: internal.ProductAttributes{
				Description: reqBody.Description,
				Price:       reqBody.Price,
			},
		}
		// - save
		err = h.sv.Save(&p)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "error creating product")
			return
		}

		// response
		// - serialize
		pr := ProductJSON{
			Id:          p.Id,
			Description: p.Description,
			Price:       p.Price,
		}
		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "product created",
			"data":    pr,
		})
	}
}

// GetTopNProducts returns the top n products ordered by quantity sold
// - query param: n
// - response: top n products ordered by quantity sold
func (h *ProductsDefault) GetTopNProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - query
		nString := chi.URLParam(r, "n")
		n, err := strconv.Atoi(nString)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "error parsing query param n")
			return
		}

		// process
		p, err := h.sv.TopNProducts(n)
		if err != nil {
			switch err {
			case internal.ErrServiceInvalidArgument:
				response.Error(w, http.StatusBadRequest, "the argument N must be greater than 0")
			default:
				response.Error(w, http.StatusInternalServerError, "error getting top N products")
			}
			return
		}

		// response
		// - serialize
		pJSON := make([]ProductSaleCountJSON, len(p))
		for ix, v := range p {
			pJSON[ix] = ProductSaleCountJSON{
				Description: v.Description,
				SaleCount:   v.SaleCount,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"data": pJSON,
		})
	}
}
