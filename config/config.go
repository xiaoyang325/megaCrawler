package config

import (
	"encoding/json"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path"
	"time"
)

type Config struct {
	Id       string    `json:"id"`
	LastIter time.Time `json:"lastIter"`
	Disabled bool      `json:"disabled"`
}

type CfgMap map[string]Config

var Configs = CfgMap{}
var Port = ":7171"

func init() {
	userPath, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	folderPath := path.Join(userPath, "/.crawler/")
	filePath := path.Join(folderPath, "config.json")
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		err = os.MkdirAll(folderPath, 0700)
		if err != nil {
			panic(err)
		}
	}

	reader, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}

	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&Configs)
	if err != nil {
		Configs = CfgMap{}
		err := Configs.Save()
		if err != nil {
			panic(err)
		}
	}
}

func (c CfgMap) Save() (err error) {
	userPath, err := os.UserHomeDir()
	if err != nil {
		return
	}

	folderPath := path.Join(userPath, "/.crawler/")
	filePath := path.Join(folderPath, "config.json")
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		err = os.MkdirAll(folderPath, 0700)
		if err != nil {
			return err
		}
	}

	if reader, err := os.Create(filePath); err == nil {
		err = json.NewEncoder(reader).Encode(&Configs)
		if err != nil {
			return err
		}
	} else {
		return err
	}
	return err
}
