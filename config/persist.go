package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"

	"j322.ica/gumroad-sammi/ipify"
)

var Config *Configuration

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
		log.Printf("I could get the config file path: %s\n", err)
		return false
	}
	_, err = os.Stat(p)
	ok := err == nil
	if !ok {
		log.Printf("I could not check the config file: %s\n", err)
	}
	return ok
}

func (c *Configuration) Save() error {
	p, err := PathToFile()
	if err != nil {
		return fmt.Errorf("I could not get the config file for saving: %w", err)
	}
	file, err := os.Create(p)
	if err != nil {
		return fmt.Errorf("I could not open the config file for saving: %w", err)
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(&c)
	if err != nil {
		return fmt.Errorf("I could not save the config file: %w\n", err)
	}
	return nil
}

func load() error {
	Config = &Configuration{Advanced{}, SammiConfig{}, GumroadConfig{}, FourthWallConfig{}, StreamElementsConfig{}}
	p, err := PathToFile()
	if err != nil {
		return fmt.Errorf("I could not get the config file for reading: %w", err)
	}
	file, err := os.Open(p)
	if err != nil {
		return fmt.Errorf("I could not open the config file for reading: %w", err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(Config)
	if err != nil {
		return fmt.Errorf("I could not parse the config file: %w", err)
	}
	log.Printf("Loaded config from %s", p)
	return nil
}

func NewConfig() {
	err := load()
	if err != nil {
		log.Printf("Could not create config based on file: %s. Reverting to default configuration.\n", err)
	}

	adv := &Config.Advanced

	if Config.SammiConfig.Host == "" {
		Config.SammiConfig.Host = defaultSammiHost
	}

	if Config.SammiConfig.Port == "" {
		Config.SammiConfig.Port = defaultSammiPort
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
}
