package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"path/filepath"
	"runtime"
)

type RunMode int

const (
	Debug      RunMode = 0
	Preview    RunMode = 1
	Production RunMode = 2
)

var (
	configs *Configs
)

type Configs struct {
	Mode          RunMode
	ServiceName   string `json:"svcname"`
	DBServiceName string `json:"dbsvcname"`
}

func Instance() *Configs {
	return configs
}

//LoadConfigs can only invoked from main.go
func LoadConfigs(m RunMode, relativePath string) (*Configs, error) {
	var abspath string
	var content []byte

	_, wd, _, ok := runtime.Caller(1)
	if !ok {
		return nil, errors.New("runtime caller not ok")
	}

	switch m {
	case Debug:
		abspath = filepath.Join(filepath.Dir(wd), relativePath, "/debug.json")
	case Preview:
		abspath = filepath.Join(filepath.Dir(wd), relativePath, "/preview.json")
	case Production:
		abspath = filepath.Join(filepath.Dir(wd), relativePath, "/production.json")
	}

	content, err := ioutil.ReadFile(abspath)
	if err != nil {
		return nil, err
	}

	configs = &Configs{Mode: m}
	jerr := json.Unmarshal(content, configs)
	if jerr != nil {
		return nil, jerr
	}

	return configs, nil
}
