package config

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"
)

type Config struct {
	WindowWidth  int
	WindowHeight int
	WindowScale  float64
	WindowTitle  string

	DefaultImageWidth    int
	DefaultImageHeight   int
	DefaultImageFilename string
	DefaultImageMargin   int

	BoardWidth  int
	BoardHeight int
	BoardScale  float64

	GenerationTimeout time.Duration
}

var (
	config Config
	once   sync.Once
)

func Load() *Config {
	once.Do(func() {
		config = Config{
			WindowWidth:  1024,
			WindowHeight: 768,
			WindowScale:  1,
			WindowTitle:  "Game of life",

			DefaultImageWidth:    64,
			DefaultImageHeight:   64,
			DefaultImageFilename: "resources/alive.png",
			DefaultImageMargin:   0,

			BoardWidth:  16, // Default 16
			BoardHeight: 12, // Default 12
			BoardScale:  1,  // Default 1

			GenerationTimeout: time.Millisecond * 500,
		}

		configBytes, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Configuration:", string(configBytes))
	})

	return &config
}
