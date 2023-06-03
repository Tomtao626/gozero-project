package handler

import (
	"userapi/internal/svc"
)

type UserHandler struct {
	serCtx *svc.ServiceContext
}

func NewUserHandler(serverCtx *svc.ServiceContext) *UserHandler {
	return &UserHandler{
		serCtx: serverCtx,
	}
}
