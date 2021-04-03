package handlers

import (
	"InkaTry/warehouse-storage-be/internal/pkg/stores/mysql"
)

type Handler struct {
	db mysql.Clienter
}

type Params struct {
	DB mysql.Clienter
}

func NewAdminHandler(params *Params) *Handler {

	return &Handler{
		db: params.DB,
	}
}
