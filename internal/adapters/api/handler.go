package api

import (
	"net/http"

	"github.com/iamtbay/is-management/internal/domain"
	"github.com/iamtbay/is-management/internal/service"
)

type HTTPHandler struct {
	productService *service.ProductService
	orderService   *service.OrderService
}

// create handler
func NewHTTPHandler(productService *service.ProductService, orderService *service.OrderService) *HTTPHandler {
	return &HTTPHandler{
		productService: productService,
		orderService:   orderService,
	}
}

// CreateOrder godoc
// @Summary Create a new order
// @Description Adds a new order to the orders
// @Tags orders
// @Accept json
// @Produce json
// @Param order body domain.Order true "Order Info"
// @Success 201 {object} domain.Order
// @Failure 400 {object} string
// @Router /orders [post]
func (h *HTTPHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order domain.Order
	if err := h.readJSON(w, r, &order); err != nil {
		h.writeError(w, http.StatusBadRequest, "Invalid JSON format or body too large")
		return
	}
	ctx := r.Context()
	err := h.orderService.CreateOrder(&order, ctx)
	if err != nil {
		h.writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	h.writeJSON(w, http.StatusCreated, &order)
}

// CreateProduct godoc
// @Summary Create a new product
// @Description Adds a new product to the inventory
// @Tags products
// @Accept json
// @Produce json
// @Param product body domain.Product true "Product Info"
// @Success 201 {object} domain.Product
// @Failure 400 {object} string
// @Router /products [post]
func (h *HTTPHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product domain.Product
	if err := h.readJSON(w, r, &product); err != nil {
		h.writeError(w, http.StatusBadRequest, "Invalid JSON format or body too large")
		return
	}
	if product.Stock < 1 {
		h.writeError(w, http.StatusBadRequest, "stock must be greater than 0")
		return
	}
	ctx := r.Context()
	err := h.productService.CreateProduct(&product, ctx)
	if err != nil {
		h.writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	h.writeJSON(w, http.StatusCreated, &product)
}

// FindProductByID godoc
// @Summary Find a product by ID
// @Description Finds a product by ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} domain.Product
// @Failure 400 {object} string
// @Router /products/{id} [get]
func (h *HTTPHandler) FindProductByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.PathValue("id")
	if id == "" {
		h.writeError(w, http.StatusBadRequest, "ID is empty")
		return
	}

	product, err := h.productService.FindProductByID(id, ctx)
	if err != nil {
		h.writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	h.writeJSON(w, http.StatusOK, &product)
}

// UpdateStock godoc
// @Summary Update a product's stock
// @Description Updates a product's stock
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Param stock body UpdateStockRequest true "Stock quantity"
// @Success 200 {object} domain.Product
// @Failure 400 {object} string
// @Router /products/{id} [patch]
type UpdateStockRequest struct {
	Quantity int `json:"quantity"`
}

func (h *HTTPHandler) UpdateStock(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var stock UpdateStockRequest
	if err := h.readJSON(w, r, &stock); err != nil {
		h.writeError(w, http.StatusBadRequest, "Invalid JSON format or body too large")
		return
	}
	id := r.PathValue("id")
	if id == "" {
		h.writeError(w, http.StatusBadRequest, "ID is empty")
		return
	}
	product, err := h.productService.UpdateStock(id, stock.Quantity, ctx)
	if err != nil {
		h.writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	h.writeJSON(w, http.StatusOK, &product)
}

// FindAllProducts godoc
// @Summary Find all products
// @Description Finds all products
// @Tags products
// @Accept json
// @Produce json
// @Success 200 {object} []domain.Product
// @Failure 400 {object} string
// @Router /products [get]
func (h *HTTPHandler) FindAllProducts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	products, err := h.productService.FindAll(ctx)
	if err != nil {
		h.writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	h.writeJSON(w, http.StatusOK, &products)
}

// FindAllOrders godoc
// @Summary Find all orders
// @Description Finds all orders
// @Tags orders
// @Accept json
// @Produce json
// @Success 200 {object} []domain.Order
// @Failure 400 {object} string
// @Router /orders [get]
func (h *HTTPHandler) FindAllOrders(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	orders, err := h.orderService.FindAll(ctx)
	if err != nil {
		h.writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	h.writeJSON(w, http.StatusOK, &orders)
}

// FindOrderByID godoc
// @Summary Find an order by ID
// @Description Finds an order by ID
// @Tags orders
// @Accept json
// @Produce json
// @Param id path string true "Order ID"
// @Success 200 {object} domain.Order
// @Failure 400 {object} string
// @Router /orders/{id} [get]
func (h *HTTPHandler) FindOrderByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.PathValue("id")
	if id == "" {
		h.writeError(w, http.StatusBadRequest, "ID is empty")
		return
	}

	order, err := h.orderService.FindByID(id, ctx)
	if err != nil {
		h.writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	h.writeJSON(w, http.StatusOK, &order)
}
