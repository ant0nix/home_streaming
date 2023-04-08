package delivery

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) StartPage(c *gin.Context) {
	log.Println(c.Request.URL.Path)
	if c.Request.URL.Path == "/start/w" {
		c.HTML(http.StatusOK, "index2.html", gin.H{"title": "home pageW"})
	} else {
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "home page"})
	}
}

func (h *Handler) Download(c *gin.Context) {
	link := c.PostForm("link")
	torrentFile, err := h.usecase.StartDownload(link, c)
	if err != nil {
		if err.Error() == "no such web-site" {
			c.Redirect(301, "/start/w")
			return
		}
		newErrorResponce(c, err.Error(), http.StatusInternalServerError)
		return
	}
	c.Redirect(301, "/videos/"+torrentFile.Name())
}

func (h *Handler) Play(c *gin.Context) {
	filename := c.Param("filename")
	if filename == "" {
		log.Println("Empty filename")
		return
	}
	os.Rename("videos/"+filename, "videos/"+deleteSpaces(filename))
	os.Remove("videos/" + filename)
	filename = deleteSpaces(filename)

	outputVideoPath, err := h.usecase.ConvertToMP4(filename)
	if err != nil {
		newErrorResponce(c, err.Error(), 500)
	}
	_, err = os.Stat(outputVideoPath)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	videoFile, err := os.Open(outputVideoPath)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	defer videoFile.Close()
	c.Writer.Header().Set("Content-Type", "video/mp4")
	c.File(outputVideoPath)
}

func deleteSpaces(str string) string {
	if !strings.ContainsAny(str, " ") {
		return str
	}
	return strings.ReplaceAll(str, " ", "")
}
