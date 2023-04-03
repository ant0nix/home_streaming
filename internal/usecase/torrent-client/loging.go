package torrentClient

import (
	"log"
	"time"

	"github.com/anacrolix/torrent"
)

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
	if !*flag && torrent.BytesCompleted()*100/torrent.Length() >= 25 {
		log.Println("You can to start wathcing a film")
		*flag = true
	}
}
