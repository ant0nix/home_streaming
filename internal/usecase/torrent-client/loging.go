package torrentClient

import (
	"log"
	"time"

	"github.com/anacrolix/torrent"
)

func logTime(timE time.Time) {
	if timE.Second() < 1 {
		return
	}
	if time.Since(timE).Seconds() < 60 {
		log.Printf("Time:%ds", int(time.Since(timE).Seconds()))
	} else {
		log.Printf("Time:%dm %ds", int(time.Since(timE).Minutes()), int(time.Since(timE).Seconds())-60)
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
}
