package torrentClient

import (
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/anacrolix/torrent"
	"github.com/gin-gonic/gin"
)

type TorrnetClient struct {
	Client *torrent.Client
}

type TorrentConfig struct {
	config *torrent.ClientConfig
}

// TODO: реальная настройка конфига
func NewConfig() *TorrentConfig {
	cfg := torrent.NewDefaultClientConfig()
	cfg.DataDir = "./videos/"
	return &TorrentConfig{
		config: cfg,
	}
}

func New(cfg TorrentConfig) (*TorrnetClient, error) {
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

func (tc *TorrnetClient) StartDownload(fileName string, c *gin.Context) (*torrent.Torrent, error) {
	client := &http.Client{}
	resp, err := client.Get(fileName)
	if err != nil {
		log.Println("No such web-site:", err.Error())
		return &torrent.Torrent{}, err
	}
	defer resp.Body.Close()

	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		log.Println(err.Error())
		return &torrent.Torrent{}, err
	}
	defer tmpfile.Close()

	_, err = io.Copy(tmpfile, resp.Body)
	if err != nil {
		log.Println(err.Error())
		return &torrent.Torrent{}, err
	}

	torrentFile, err := tc.NewTorrent(tmpfile.Name())
	if err != nil {
		log.Printf("Error with start downloading: %s", err.Error())
		return &torrent.Torrent{}, err
	}

	log.Printf("Now is downloadnig: %s", torrentFile.Name())
	torrentFile.DownloadAll()

	starttime := time.Now()
	var flag bool
	var yetDownloaded int64

	for !torrentFile.Complete.Bool() {
		select {
		case <-c.Writer.CloseNotify():
			torrentFile.Drop()
			return nil, errors.New("client closed connection")
		default:
			download := torrentFile.BytesCompleted() - yetDownloaded
			logTime(starttime)
			logSpeed(download)
			if download == 0 {
				time.Sleep(5 * time.Second)
				continue
			}
			logProgress(*torrentFile, &flag)
			yetDownloaded = torrentFile.BytesCompleted()
			time.Sleep(5 * time.Second)
		}
	}
	log.Println("Downloading has completed!")
	return torrentFile, nil
}
