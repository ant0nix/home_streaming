package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler)Download(c *gin.Context){
	//ВРЕМЕННО Тестовое решение
	err := h.usecase.StartDownload("file.torrent")
	if err != nil{
		newErrorResponce(c, err.Error(),http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK,map[string]interface{}{
		"answer":"File has downloaded",
	})
}