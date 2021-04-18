package admin

import (
	"InkaTry/warehouse-storage-be/internal/http/admin/handlers"
	"InkaTry/warehouse-storage-be/internal/http/admin/http/api"
	"github.com/gorilla/mux"
	"net/http"
)

func Routes(router *mux.Router, handler *handlers.Handler) {
	adminApis := router.PathPrefix("/admin").Subrouter()

	// public api connection test
	adminApis.HandleFunc("/ping", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("pong"))
	})

	v1Apis := adminApis.PathPrefix("/v1").Subrouter()

	v1Apis.Path("/auto").Handler(api.Autocomplete(handler.Autocomplete)).
		Methods(http.MethodGet)

	v1Apis.Path("/list/warehouses").Handler(api.ListWarehouses(handler.ListWarehouses)).
		Methods(http.MethodGet)
	v1Apis.Path("/list/products").Handler(api.ListProducts(handler.ListProducts)).
		Methods(http.MethodGet)
	v1Apis.Path("/list/product_types").Handler(api.ListProductTypes(handler.ListProductTypes)).
		Methods(http.MethodGet)
	v1Apis.Path("/list/brands").Handler(api.ListBrands(handler.ListBrands)).
		Methods(http.MethodGet)
	v1Apis.Path("/download/products").Handler(api.DownloadProducts(handler.DownloadProducts)).
		Methods(http.MethodGet)

}
