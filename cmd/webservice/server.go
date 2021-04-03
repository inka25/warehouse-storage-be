package webservice

import (
	"InkaTry/warehouse-storage-be/internal/http/admin"
	"InkaTry/warehouse-storage-be/internal/http/admin/handlers"
	"InkaTry/warehouse-storage-be/internal/pkg/config"
	"InkaTry/warehouse-storage-be/internal/pkg/stores/mysql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"time"
)

func Start(cfg *config.Config) func() {

	db := initMysql(cfg)

	router := mux.NewRouter()
	admin.Routes(router, handlers.NewAdminHandler(&handlers.Params{
		DB: mysql.NewClient(db),
	}))

	srv := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),

		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 60 * time.Second,
		ReadTimeout:  60 * time.Second,
	}

	log.Printf("started at %s:%s", cfg.Host, cfg.Port)
	log.Fatal(srv.ListenAndServe())

	return func() {
		// stop function
	}

}

func initMysql(cfg *config.Config) *sqlx.DB {
	db, err := sqlx.Connect("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?",
		cfg.MySQL.User,
		cfg.MySQL.Pass,
		cfg.MySQL.Host,
		cfg.MySQL.Port,
		cfg.MySQL.DB,
	))
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(cfg.MySQL.MaxIdle)
	db.SetMaxOpenConns(cfg.MySQL.MaxOpen)

	return db
}
