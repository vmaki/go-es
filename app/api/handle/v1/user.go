package v1

import (
	"github.com/gin-gonic/gin"
	"go-es/app/api/handle"
	"go-es/common/ctxdata"
	"go-es/common/responsex"
	"go-es/internal/services"
)

type UserHandle struct {
	handle.BaseAPIController
}

func (h *UserHandle) List(ctx *gin.Context) {
	s := services.NewUserService()
	list, pager := s.List(ctx, 10)

	responsex.List(ctx, list, pager)
}

// Info
// @Tags      User
// @Summary   获取用户信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200   {object}  responsex.Response{data=dto.UserInfoResp}
// @Router    /user [get]
func (h *UserHandle) Info(ctx *gin.Context) {
	uid := ctxdata.CurrentUID(ctx)

	s := services.NewUserService()
	data, err := s.Info(uid)
	if err != nil {
		responsex.Error(ctx, err)
		return
	}

	responsex.Success(ctx, data)
}
