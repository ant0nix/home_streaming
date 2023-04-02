package delivery

import (
	"github.com/ant0nix/home_streaming/internal/usecase"
	"github.com/gin-gonic/gin"
)

type Handler struct{
	usecase usecase.UseCase
}

func NewHandler(uc *usecase.UseCase) *Handler{
	return &Handler{
		usecase: *uc,
	}
}

func (h *Handler)InitRouters() *gin.Engine{
	router := gin.New()

	torrents:= router.Group("/download")
	{
		torrents.GET("/",h.Download)
	}
	return router
}