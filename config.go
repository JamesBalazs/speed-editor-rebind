package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"path/filepath"
)

const (
	configPath       = "speed-editor-rebind-config"
	keybindsFilename = "keybinds.json"
)

type Config map[uint16]*ConfigKeybind

type ConfigKeybind struct {
	LedMode string `json:"ledMode"`
}

var (
	config     Config
	configFile *os.File
)

func writeConfig() {
	var (
		configBytes []byte
		err         error
	)

	configBytes, _ = json.Marshal(config)
	if _, err = configFile.Seek(0, 0); err != nil {
		log.Fatalf("error seeking to start of keybinds config file: %v", err)
	} else if _, err = configFile.Write(configBytes); err != nil {
		log.Fatalf("error writing keybinds config file: %v", err)
	}
}

func createOrOpenConfigFile() {
	var err error

	home, _ := os.UserHomeDir()
	path := filepath.Join(home, configPath)
	if err = os.MkdirAll(path, os.ModePerm); err != nil {
		log.Fatalf("error creating config dir: %v", err)
	}
	path = filepath.Join(path, keybindsFilename)
	if configFile, err = os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644); err != nil {
		log.Fatalf("error creating or opening keybinds config file: %v", err)
	}
}

func readOrInitConfigFile() {
	var (
		configBytes []byte
		err         error
	)

	if configBytes, err = io.ReadAll(configFile); err != nil {
		log.Fatalf("error reading keybinds config file: %v", err)
	} else if len(configBytes) == 0 {
		config = Config{}
		writeConfig()
	} else if err = json.Unmarshal(configBytes, &config); err != nil {
		log.Fatalf("error unmarshalling keybinds: %v", err)
	}
}

func ConfigUpdateLedMode(keyId uint16, mode string) {
	if keybind, ok := config[keyId]; ok {
		keybind.LedMode = mode
	} else {
		config[keyId] = &ConfigKeybind{LedMode: mode}
	}

	writeConfig()
}
