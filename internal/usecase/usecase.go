package usecase

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/anacrolix/torrent"
)

type UseCase struct {
	torents ITorrents
}

func New(it ITorrents) *UseCase {
	return &UseCase{
		torents: it,
	}
}

func (us *UseCase) StartDownload(fileName string) (*torrent.Torrent, error) {
	return us.torents.StartDownload(fileName)
}

func (us *UseCase) ConvertToMP4(fileName string) (string, error) {
	videoPath := filepath.Join("videos", fileName)
	outputVideoPath := "./videos/" + "output.mp4"
	_, err := os.Stat(outputVideoPath)
	if os.IsNotExist(err) {
		cmd := exec.Command("ffmpeg", "-i", videoPath, "-preset", "ultrafast", "-threads", "4", outputVideoPath)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			log.Fatalf("Video convertaning error: %s", err)

			log.Printf("command failed with output: %s\n", outputVideoPath)

			log.Println(err.Error())
			return "", err
		}
	}
	return outputVideoPath, nil
}
