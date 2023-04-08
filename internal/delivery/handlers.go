package delivery

import (
	"github.com/ant0nix/home_streaming/internal/usecase"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	usecase usecase.IUseCases
}

func NewHandler(uc usecase.IUseCases) *Handler {
	return &Handler{
		usecase: uc,
	}
}

func (h *Handler) InitRouters() *gin.Engine {
	router := gin.New()

	torrents := router.Group("/download")
	{
		torrents.POST("/", h.Download)
	}
	stream := router.Group("/videos")
	{
		stream.GET("/:filename", h.Play)
	}
	start := router.Group("/start")
	{
		start.GET("/", h.StartPage)
		start.GET("/w", h.StartPage)
	}
	router.LoadHTMLGlob("./templates/*")
	return router
}
