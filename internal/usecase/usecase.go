package usecase

import (
	"io"
	"io/ioutil"
	"log"
	"time"

	"github.com/anacrolix/torrent"
	"github.com/ant0nix/home_streaming/internal/entities"
)

type UseCase struct {
	torents ITorrents
}

func New(it ITorrents) *UseCase {
	return &UseCase{
		torents: it,
	}
}

// func (uc *UseCase) InitClient() (*entities.TorrnetClient, error) {
// 	client, err := uc.torents.NewTorrentClient(*cfg)
// 	if err != nil {
// 		log.Fatalf("Error with InitClient: %s", err.Error())
// 	}
// 	return client, nil
// }

func (us *UseCase) StartDownload(fileName string) error {

	client := entities.NewHttpClient()
	resp, err := client.Client.Get(fileName)
	if err != nil {
		log.Println("No such web-site:",err.Error())
		return err
	}
	defer resp.Body.Close()

	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		log.Println(err.Error())
		return err
	}
	defer tmpfile.Close()

	_, err = io.Copy(tmpfile, resp.Body)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	torrent, err := us.torents.NewTorrent(tmpfile.Name())
	if err != nil {
		log.Printf("Error with start downloading: %s", err.Error())
		return err
	}

	log.Printf("Now is downloadnig: %s", torrent.Name())
	torrent.DownloadAll()

	starttime := time.Now()
	var flag bool
	var yetDownloaded int64

	for !torrent.Complete.Bool() {
		download := torrent.BytesCompleted() - yetDownloaded
		logTime(starttime)
		logSpeed(download)
		if download == 0 {
			time.Sleep(5 * time.Second)
			continue
		}
		logProgress(*torrent, &flag)
		yetDownloaded = torrent.BytesCompleted()
		time.Sleep(5 * time.Second)
	}
	log.Println("Downloading has completed!")
	return nil
}

func logTime(time time.Time) {
	if time.Second() < 1 {
		return
	}
	if time.Second()/60 < 1 {
		log.Printf("Time has passed: %d s", time.Second())
		return
	} else {
		log.Printf("Time has passed: %dm%ds", time.Minute(), time.Second())
		return
	}
}

func logSpeed(downloadBytes int64) {
	if downloadBytes == 0 {
		log.Println("Wait connection to seed")
		return
	}
	fbytes := float64(downloadBytes) / 1024
	if fbytes > 1024 {
		speed := float64(fbytes) / float64(1024)
		log.Printf("Speed: %.2f MB/s", speed/5)
	} else {
		log.Printf("Speed: %.1f KB/s", fbytes/5)
	}
}

func logProgress(torrent torrent.Torrent, flag *bool) {
	log.Printf("Downloaded: %d%%\n", torrent.BytesCompleted()*100/torrent.Length())
	if !*flag && torrent.BytesCompleted()*100/torrent.Length() >= 5 {
		log.Println("You can to start wathcing a film")
		*flag = true
	}
}
