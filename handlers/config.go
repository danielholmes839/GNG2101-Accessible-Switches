package handlers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

// Config struct
type Config struct {
	Name       string          `json:"type"`
	SerialPort string          `json:"port"`
	InputDelay int             `json:"delay"`
	Config     json.RawMessage `json:"config"`
}

// IsScrollingConfig func
func (config *Config) IsScrollingConfig() bool {
	return config.Name == "scrolling"
}

// GetScrollingConfig func
func (config *Config) GetScrollingConfig() *ScrollingConfig {
	data := &ScrollingConfig{}
	err := json.Unmarshal(config.Config, &data)
	if err != nil {
		panic(err)
	}
	return data
}

// IsBisectionConfig func
func (config *Config) IsBisectionConfig() bool {
	return config.Name == "bisection"
}

// GetBisectionConfig func
func (config *Config) GetBisectionConfig() *BisectionConfig {
	data := &BisectionConfig{}
	err := json.Unmarshal(config.Config, &data)
	if err != nil {
		panic(err)
	}
	return data
}

// ScrollingConfig struct
type ScrollingConfig struct {
	Shortcut        int  `json:"Shortcut"`
	HorizontalFirst bool `json:"HorizontalFirst"`
	FramesPerSecond int  `json:"FramesPerSecond"`
	PixelsPerFrame  int  `json:"PixelsPerFrame"`
	LeftToRight     bool `json:"LeftToRight"`
	TopToBottom     bool `json:"TopToBottom"`
}

// BisectionConfig struct
type BisectionConfig struct {
	Shortcut        int  `json:"Shortcut"`
	HorizontalFirst bool `json:"HorizontalFirst"`
}

// GetConfig func
func GetConfig() *Config {
	file, _ := ioutil.ReadFile("config.json")
	config := &Config{}
	err := json.Unmarshal(file, &config)
	if err != nil {
		panic(err)
	}
	return config
}

// GetHandler func
func GetHandler(config *Config) (Handler, error) {
	if config.IsScrollingConfig() {
		return NewScrollingHandler(config.GetScrollingConfig()), nil
	} else if config.IsBisectionConfig() {
		return NewBisectionHandler(config.GetBisectionConfig()), nil
	} else {
		return nil, errors.New("Invalid config 'type' field should be 'scrolling' or 'bisection' got: " + config.Name)
	}
}
