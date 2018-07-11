package helpers

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strings"
	"sync"

	unarr "github.com/gen2brain/go-unarr"
	"github.com/kayslay/todaytvseries_organizer/config"
)

var w sync.WaitGroup
var ch = make(chan os.FileInfo)

func moveZipContent(c config.Config, f os.FileInfo) {
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

	for _, v := range files {
		if !strings.Contains(v, c.MatchExt) {
			continue
		}
		err := r.EntryFor(v)
		if err != nil {
			fmt.Println("error opening file", v, "in", c.Path+compressName, "due to", err)
			return
		}
		filename := c.GetDir(v) + v
		nFile, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0777)
		log.Println("moving to", filename)
		_, err = io.Copy(nFile, r)
		if err != nil {
			fmt.Println("error coping file to destination due to", err)
			return
		}
		if c.DeleteAfter {
			err := os.Remove(c.Path + filename)
			if err != nil {
				fmt.Println("error deleting ", c.Path+compressName, "due to", err)
				return
			}
		}

	}

	return
}

func worker(c config.Config, i int) {
	for val := range ch {
		moveZipContent(c, val)
		w.Done()
		fmt.Println("worker", i, "completed")
	}
}

func startWorker(c config.Config, n int) {
	for i := 0; i < int(math.Max(float64(c.WorkerCount), 1)); i++ {
		go worker(c, i)
	}
}

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
	return nil
}
