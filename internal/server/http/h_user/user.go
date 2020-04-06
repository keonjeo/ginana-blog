package h_user

import (
	"ginana-blog/internal/model"
	"ginana-blog/internal/service"
	"github.com/gin-gonic/gin"
)

type HUser struct {
	svc *service.Service
}

func New(s *service.Service) *HUser {
	return &HUser{
		svc: s,
	}
}

// GetUsers godoc
// @Description 测试
// @Tags Public
// @Accept  json
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {string} string "ok"
// @Failure 500 {string} string "failed"
// @Router / [get]
func (h *HUser) GetUsers(ctx *gin.Context) {
	k := &model.GiNana{
		Hello: "GiNana Server",
	}
	ctx.HTML(200, "index.html", k)
}
