package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	NumWorkers int `json:"workers"`
	NumTasks int `json:"tasks"`
	RetryCount int `json:"retry_limit"`
}

func LoadConfig(path string) (Config,error) {
	var cfg Config

	file,err:= os.Open(path)

	if err != nil {
		return cfg, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&cfg)

	return cfg, err
}