package helpers

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"path/filepath"
	"strings"
	"sync"

	unarr "github.com/gen2brain/go-unarr"
	"github.com/kayslay/todaytvseries_organizer/config"
)

var (
	w                         sync.WaitGroup
	ch                        = make(chan os.FileInfo)
	progressCount, numOfFiles int
)

// moveCompressedExt this moves the content in the compressed file to the destination folder
func moveCompressedExt(c config.Config, f os.FileInfo) {
	compressName := f.Name()
	r, err := unarr.NewArchive(c.Path + compressName)
	if err != nil {
		fmt.Println("file", c.Path+compressName, "failed to unCompress due to", err)
		return
	}
	defer func() { log.Println("transfer finished") }()
	defer r.Close()

	files, err := r.List()

	if err != nil {
		fmt.Println("failed listing files in", c.Path+compressName, "due to", err)
		return
	}

	matchExt := strings.Split(fmt.Sprintf(".%s", c.MatchExt), ",")
	// loop through the content in the compressed file
	// pick the files that match the extension suffix
	for _, v := range files {
		if !matchesExt(v, matchExt...) {
			continue
		}
		err := r.EntryFor(v)
		if err != nil {
			fmt.Println("error opening file", v, "in", c.Path+compressName, "due to", err)
			return
		}
		filename := filepath.Join(c.GetDir(v), v)
		nFile, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0777)
		log.Println("moving to", filename)
		_, err = io.Copy(nFile, r)
		if err != nil {
			fmt.Println("error copying file to destination due to", err)
			return
		}

	}
	// delete compressed file if config.Config.DeleteAfter is set to true
	if c.DeleteAfter {
		err := os.Remove(c.Path + compressName)
		if err != nil {
			fmt.Println("error deleting ", filepath.Join(c.Path, compressName), "due to", err)
			return
		}
	}
	return
}

func worker(c config.Config, i int) {
	for val := range ch {
		moveCompressedExt(c, val)
		w.Done()
		fmt.Println("worker", i, "completed")
	}
}

func startWorker(c config.Config, n int) {
	for i := 0; i < int(math.Max(float64(c.WorkerCount), 1)); i++ {
		go worker(c, i)
	}
}

//Start starts the organization of series from the rar files
func Start(c config.Config) error {
	files, err := findExt(c)
	if err != nil {
		return err
	}
	w.Add(len(files))
	go startWorker(c, 2)
	for _, v := range files {
		ch <- v
	}
	w.Wait()
	close(ch)
	return nil
}
