package helpers

import (
	"fmt"
	"os"
	"strings"

	"github.com/kayslay/todaytvseries_organizer/config"
)

//findExt get all the compressed files with the config.Ext name
func findExt(c config.Config) ([]os.FileInfo, error) {
	dir, err := os.Open(c.Path)
	if err != nil {
		return nil, err
	}
	stat, err := dir.Stat()

	if err != nil {
		fmt.Println("the error message", err)
		return nil, err
	}
	if !stat.IsDir() {
		return nil, fmt.Errorf("path specified must be a directory")
	}
	matchFile := []os.FileInfo{}
	files, err := dir.Readdir(0)
	if err != nil {
		fmt.Println("the error message", err)
		return nil, err
	}
	for _, v := range files {
		if !v.IsDir() && strings.HasSuffix(v.Name(), fmt.Sprintf(".%s", c.Ext)) {
			matchFile = append(matchFile, v)
		}
	}
	return matchFile, nil
}
