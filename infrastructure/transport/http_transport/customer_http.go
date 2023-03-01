package http_transport

import (
	"domain-driven-design/domain/model/entity"
	"domain-driven-design/domain/service"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type CustomerHandler struct {
	customerService service.CustomerService
}

func NewCustomerHandler(customerService service.CustomerService) *CustomerHandler {
	return &CustomerHandler{
		customerService: customerService,
	}
}

func (ch *CustomerHandler) GetCustomers(c echo.Context) error {
	customers, err := ch.customerService.GetCustomers()
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, customers)
}

func (ch *CustomerHandler) GetCustomer(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	customer, err := ch.customerService.GetCustomer(int64(id))
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, customer)
}

func (ch *CustomerHandler) CreateCustomer(c echo.Context) error {
	var newCustomer service.NewCustomer
	err := json.NewDecoder(c.Request().Body).Decode(&newCustomer)
	if err != nil {
		return err
	} else {
		customer := &entity.Customer{
			Name:    newCustomer.Name,
			Email:   newCustomer.Email,
			Address: newCustomer.Address,
		}
		data, err := ch.customerService.CreateCustomer(customer)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, data)
	}
}
