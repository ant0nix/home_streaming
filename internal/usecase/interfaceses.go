package usecase

import (
	"github.com/anacrolix/torrent"
	"github.com/gin-gonic/gin"
)

type ITorrents interface {
	StartDownload(fileName string, c *gin.Context) (*torrent.Torrent, error)
}

type IUseCases interface {
	StartDownload(fileName string, c *gin.Context) (*torrent.Torrent, error)
	ConvertToMP4(fileName string) (string, error)
}
