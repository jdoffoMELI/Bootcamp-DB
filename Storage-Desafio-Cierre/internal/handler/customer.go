package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"app_desafio/internal"
	"app_desafio/platform/web/request"
	"app_desafio/platform/web/response"

	"github.com/go-chi/chi/v5"
)

// NewCustomersDefault returns a new CustomersDefault
func NewCustomersDefault(sv internal.ServiceCustomer) *CustomersDefault {
	return &CustomersDefault{sv: sv}
}

// CustomersDefault is a struct that returns the customer handlers
type CustomersDefault struct {
	// sv is the customer's service
	sv internal.ServiceCustomer
}

// CustomerJSON is a struct that represents a customer in JSON format
type CustomerJSON struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Condition int    `json:"condition"`
}

// CustomerMoneySpentJSON is a struct that represents the money spent by a customer in JSON format
type CustomerMoneySpentJSON struct {
	FirstName  string  `json:"first_name"`
	LastName   string  `json:"last_name"`
	TotalSpent float64 `json:"total_spent"`
}

// GetAll returns all customers
func (h *CustomersDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		c, err := h.sv.FindAll()
		if err != nil {
			log.Println(err)
			response.Error(w, http.StatusInternalServerError, "error getting customers")
			return
		}

		// response
		// - serialize
		csJSON := make([]CustomerJSON, len(c))
		for ix, v := range c {
			csJSON[ix] = CustomerJSON{
				Id:        v.Id,
				FirstName: v.FirstName,
				LastName:  v.LastName,
				Condition: v.Condition,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "customers found",
			"data":    csJSON,
		})
	}
}

// RequestBodyCustomer is a struct that represents the request body for a customer
type RequestBodyCustomer struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Condition int    `json:"condition"`
}

// Create creates a new customer
func (h *CustomersDefault) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - body
		var reqBody RequestBodyCustomer
		err := request.JSON(r, &reqBody)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "error deserializing request body")
			return
		}

		// process
		// - deserialize
		c := internal.Customer{
			CustomerAttributes: internal.CustomerAttributes{
				FirstName: reqBody.FirstName,
				LastName:  reqBody.LastName,
				Condition: reqBody.Condition,
			},
		}
		// - save
		err = h.sv.Save(&c)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "error saving customer")
			return
		}

		// response
		// - serialize
		cs := CustomerJSON{
			Id:        c.Id,
			FirstName: c.FirstName,
			LastName:  c.LastName,
			Condition: c.Condition,
		}
		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "customer created",
			"data":    cs,
		})
	}
}

// GetNCostumer returns the top n costumers ordered by money spent
// - query param: n
// - response: top n products ordered by quantity sold
func (h *CustomersDefault) GetTopNCostumer() http.HandlerFunc {
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
		p, err := h.sv.TopNCostumers(n)
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
		pJSON := make([]CustomerMoneySpentJSON, len(p))
		for ix, v := range p {
			pJSON[ix] = CustomerMoneySpentJSON{
				FirstName:  v.FirstName,
				LastName:   v.LastName,
				TotalSpent: v.TotalSpent,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"data": pJSON,
		})
	}
}

// GetTotalMoneySpentByCondition returns the total money spent by a customer
// - query param: condition
// - response: total money spent by a customer
func (h *CustomersDefault) GetTotalMoneySpentByCondition() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - query
		conditionString := chi.URLParam(r, "condition")
		condition, err := strconv.Atoi(conditionString)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "error parsing query param condition")
			return
		}

		// process
		total, err := h.sv.GetTotalMoneySpentByCondition(condition)
		if err != nil {
			fmt.Println(err)
			switch err {
			case internal.ErrServiceInvalidArgument:
				response.Error(w, http.StatusBadRequest, "the argument condition must be 0 or 1")
			default:
				response.Error(w, http.StatusInternalServerError, "error getting total money spent by condition")
			}
			return
		}

		// response
		// - serialize
		var conditionName string
		if condition == 0 {
			conditionName = "Inactivo"
		} else {
			conditionName = "Activo"
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"condition": conditionName,
			"total":     total,
		})
	}
}
