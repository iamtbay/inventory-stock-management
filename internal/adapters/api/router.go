package api

import (
	"net/http"

	_ "github.com/iamtbay/is-management/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

func NewRouter(handler *HTTPHandler) http.Handler {
	mux := http.NewServeMux()
	//SWAGGER
	mux.HandleFunc("GET /swagger/", httpSwagger.WrapHandler)
	//PRODUCT ROUTES
	mux.HandleFunc("GET /products", handler.FindAllProducts)
	mux.HandleFunc("POST /products", handler.CreateProduct)
	mux.HandleFunc("GET /products/{id}", handler.FindProductByID)
	mux.HandleFunc("PATCH /products/{id}", handler.UpdateStock)
	//ORDER ROUTES
	mux.HandleFunc("POST /orders", handler.CreateOrder)
	mux.HandleFunc("GET /orders", handler.FindAllOrders)
	mux.HandleFunc("GET /orders/{id}", handler.FindOrderByID)

	//HEALTH CHECK
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	return LoggerMiddleware(mux)
}
