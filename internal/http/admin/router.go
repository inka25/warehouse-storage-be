package admin

import (
	"InkaTry/warehouse-storage-be/internal/http/admin/handlers"
	"InkaTry/warehouse-storage-be/internal/http/admin/http/api"
	"InkaTry/warehouse-storage-be/internal/pkg/config"
	"InkaTry/warehouse-storage-be/middleware"
	"github.com/gorilla/mux"
	"net/http"
)

func Routes(router *mux.Router, handler *handlers.Handler, cfg *config.Config) {
	// public api connection test
	router.HandleFunc("/ping", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("pong"))
	})

	adminApis := router.PathPrefix("/admin").Subrouter()
	adminApis.Use(middleware.Middleware(cfg))
	v1Apis := adminApis.PathPrefix("/v1").Subrouter()

	v1Apis.Path("/auto").Handler(api.Autocomplete(handler.Autocomplete)).
		Methods(http.MethodGet)

	v1Apis.Path("/list/warehouses").Handler(api.ListWarehouses(handler.ListWarehouses)).
		Methods(http.MethodGet)
	v1Apis.Path("/list/products").Handler(api.ListProducts(handler.ListProducts)).
		Methods(http.MethodGet)
	v1Apis.Path("/list/inventories").Handler(api.ListInventories(handler.ListInventories)).
		Methods(http.MethodGet)
	v1Apis.Path("/list/product_types").Handler(api.ListProductTypes(handler.ListProductTypes)).
		Methods(http.MethodGet)
	v1Apis.Path("/list/brands").Handler(api.ListBrands(handler.ListBrands)).
		Methods(http.MethodGet)
	v1Apis.Path("/list/countries").Handler(api.ListCountries(handler.ListCountries)).
		Methods(http.MethodGet)

	v1Apis.Path("/download/products").Handler(api.DownloadProducts(handler.DownloadProducts)).
		Methods(http.MethodGet)
	v1Apis.Path("/download/inventories").Handler(api.DownloadInventories(handler.DownloadInventories)).
		Methods(http.MethodGet)

	v1Apis.Path("/upload/inventories/template").Handler(api.UploadInventoriesTemplate(handler.UploadInventoriesTemplate)).
		Methods(http.MethodGet)
	v1Apis.Path("/upload/inventories").Handler(api.UploadInventories(handler.UploadInventories)).
		Methods(http.MethodPost)

	v1Apis.Path("/product").Handler(api.GetProductDetail(handler.GetProductDetail)).
		Methods(http.MethodGet)

}
