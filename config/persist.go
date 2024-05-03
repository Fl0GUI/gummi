package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"

	"j322.ica/gumroad-sammi/ipify"
)

var config *Configuration

var configFile = "gummi.config.json"

func PathToFile() (string, error) {
	p, err := os.UserConfigDir()
	if err == nil {
		p = path.Join(p, configFile)
	}
	return p, err
}

func FileExists() bool {
	p, err := PathToFile()
	if err != nil {
		log.Printf("Config file path error: %s\n", err)
		return false
	}
	_, err = os.Stat(p)
	ok := err == nil
	if !ok {
		log.Printf("Config file read error: %s\n", err)
	}
	return ok
}

func (c *Configuration) Save() error {
	p, err := PathToFile()
	if err != nil {
		return fmt.Errorf("Config file save: %w", err)
	}
	file, err := os.Create(p)
	if err != nil {
		return fmt.Errorf("Config file save: create failure: %w", err)
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(&c)
	if err != nil {
		return fmt.Errorf("Config file save: write failure: %w\n", err)
	}
	return nil
}

func load() error {
	config = &Configuration{Advanced{}, SammiConfig{}, GumroadConfig{}, FourthWallConfig{}, ThroneConfig{}}

	config.GumroadConfig.Active = true
	config.FourthWallConfig.Active = true
	config.ThroneConfig.Active = true

	p, err := PathToFile()
	if err != nil {
		return fmt.Errorf("Config file load: %w", err)
	}
	file, err := os.Open(p)
	if err != nil {
		return fmt.Errorf("Config file load: open failure: %w", err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(config)
	if err != nil {
		return fmt.Errorf("Config file load: read failure: %w", err)
	}
	log.Printf("Config file load: succes. path=%s", p)
	return nil
}

func NewConfig() *Configuration {
	err := load()
	if err != nil {
		log.Printf("Config initialization failure: %s. Reverting to default Configuration.\n", err)
	}

	adv := &config.Advanced

	if config.SammiConfig.Host == "" {
		config.SammiConfig.Host = defaultSammiHost
	}

	if config.SammiConfig.Port == "" {
		config.SammiConfig.Port = defaultSammiPort
	}

	if adv.ServerConfig.ServerPort == "" {
		adv.ServerConfig.ServerPort = defaultServerPort
	}
	if adv.ServerConfig.PublicIp == "" {
		ip, err := ipify.Get()
		if err != nil {
			panic(err)
		}
		adv.ServerConfig.PublicIp = ip
	}
	if adv.BufferSize == 0 {
		adv.BufferSize = defaultBufferSize
	}
	if adv.BackoffAttempts == 0 {
		adv.BackoffAttempts = defaultBackoffAttempts
	}
	return config
}
