package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type ConfigReg struct {
	*regexp.Regexp
}

func (r ConfigReg) MarshalJSON() ([]byte, error) {
	return []byte(r.Regexp.String()), nil
}

//UnmarshalJSON converts string to ConfigReg which is type of regexp.Regexp
// all `//` will be converted `/`.
func (r *ConfigReg) UnmarshalJSON(b []byte) error {
	formatStr := strings.Replace(string(b), `\\`, `\`, -1)
	reg, err := regexp.Compile(formatStr[1 : len(formatStr)-1])
	if err != nil {
		return err
	}
	r.Regexp = reg
	return nil
}

type Config struct {
	DeleteAfter bool      `json:"deleteAfter,omitempty" `
	MoveDir     string    `json:"moveDir,omitempty"`
	Path        string    `json:"path,omitempty"`
	Ext         string    `json:"ext,omitempty" `
	MatchExt    string    `json:"matchExt,omitempty"`
	FolderName  ConfigReg `json:"folderName,omitempty" `
	WorkerCount int8      `json:"workerCount,omitempty" `
}

var defReg = (regexp.MustCompile(`(.+)?\.S(\d{2})`))

//DefaultConfig the default config
var DefaultConfig Config = Config{
	DeleteAfter: false,
	MoveDir:     "./",
	Path:        "./",
	Ext:         ".zip",
	MatchExt:    ".go",
	FolderName:  ConfigReg{defReg},
	WorkerCount: 1,
}

//GetDir return the directory name to save the file to
func (config Config) GetDir(name string) string {
	subFolders := config.FolderName.FindStringSubmatch(name)
	if len(subFolders) < 2 {
		return ""
	}
	dirname := filepath.Join(subFolders[1:]...)
	if err := os.MkdirAll(dirname, os.ModeDir|os.ModePerm); err != nil {
		if !strings.Contains(err.Error(), "file exist") {
			log.Fatalln(err)
		}
	}
	return dirname
}

func loadFromFile(config *Config) bool {
	file, err := os.Open("./config.json")
	if err != nil {
		return false
	}

	err = json.NewDecoder(file).Decode(config)
	if err != nil {
		return false
	}
	fmt.Println("config created from config.json in directory")
	return true
}

func loadFromFlag(config *Config) bool {
	return false
}

func loadDefault(config *Config) bool {
	fmt.Println("config created from defaultConfig")
	*config = DefaultConfig
	return true
}

//InitConfig create the Config
func InitConfig() *Config {
	config := &Config{}
	if ok := loadFromFile(config); !ok {
		if ok := loadFromFlag(config); !ok {
			loadDefault(config)
		}
	}
	return config
}
