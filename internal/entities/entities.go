package entities

import "github.com/anacrolix/torrent"

type TorrnetClient struct {
	Client *torrent.Client
}

type TorrentConfig struct {
	config *torrent.ClientConfig
}

// TODO: реальная настройка конфига
func NewTorrentConfig() *TorrentConfig {
	return &TorrentConfig{
		config: torrent.NewDefaultClientConfig(),
	}
}

func (tc *TorrnetClient) NewTorrentClient(cfg TorrentConfig) (*TorrnetClient, error) {
	cln, err := torrent.NewClient(cfg.config)
	if err != nil {
		return &TorrnetClient{}, err
	} else {
		return &TorrnetClient{
			Client: cln,
		}, nil
	}
}

func (tc *TorrnetClient) NewTorrent(fileName string) (*torrent.Torrent, error) {
	return tc.Client.AddTorrentFromFile(fileName)
}

