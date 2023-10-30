package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"

	"j322.ica/gumroad-sammi/ipify"
)

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

func (c *Config) Save() error {
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

func Load() (*Config, error) {
	res := &Config{}
	p, err := PathToFile()
	if err != nil {
		return res, fmt.Errorf("I could not get the config file for reading: %w", err)
	}
	file, err := os.Open(p)
	if err != nil {
		return res, fmt.Errorf("I could not open the config file for reading: %w", err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(res)
	if err != nil {
		return res, fmt.Errorf("I could not parse the config file: %w", err)
	}
	log.Printf("Loaded config from %s", p)
	return res, nil
}

func NewConfig() *Config {
	res := &Config{}
	if res.SammiConfig.Host == "" {
		res.SammiConfig.Host = defaultSammiHost
	}
	if res.SammiConfig.Port == "" {
		res.SammiConfig.Port = defaultSammiPort
	}
	if res.GumroadConfig.ServerPort == "" {
		res.GumroadConfig.ServerPort = defaultGumroadPort
	}
	if res.GumroadConfig.PublicIp == "" {
		ip, err := ipify.Get()
		if err != nil {
			panic(err)
		}
		res.GumroadConfig.PublicIp = ip
	}
	return res
}
