package usecase

import (
	"github.com/anacrolix/torrent"
	"github.com/ant0nix/home_streaming/internal/entities"
)

type ITorrents interface {
	// NewTorrentConfig() *entities.TorrentConfig
	NewTorrentClient(cfg entities.TorrentConfig) (*entities.TorrnetClient, error)
	NewTorrent(fileName string) (*torrent.Torrent, error)
	// StartStreaming() error
}
