package usecase

import (
	"github.com/anacrolix/torrent"
)

type ITorrents interface {
	//NewTorrentClient(cfg entities.TorrentConfig) (*entities.TorrnetClient, error)
	StartDownload(fileName string) (*torrent.Torrent, error)
}

type IUseCases interface{
	StartDownload(fileName string) (*torrent.Torrent, error)
	ConvertToMP4(fileName string) (string, error)
}