package main

import (
	"flag"
	"fmt"
	"path"
	"strconv"
	"time"
)

type Config struct {
	Interval time.Duration
	OuputDir string
}

func parseFlags() (*Config, error) {
	var (
		intervalString string
		outputDir      string
	)
	flag.StringVar(&intervalString, "interval", "15s", "Interval to capture screenshot. Ex: 15s means capture every 15 seconds.")
	flag.StringVar(&outputDir, "output", path.Join(".", "captures"), "Directory to save screenshots.")
	flag.Parse()

	if len(intervalString) == 0 {
		return nil, fmt.Errorf("Bad interval format as string, %s", intervalString)
	}

	digits := intervalString[:len(intervalString)-1]
	letter := intervalString[len(intervalString)-1]

	number, err := strconv.ParseInt(digits, 10, 64)
	if err != nil {
		return nil, err
	}

	config := &Config{
		OuputDir: outputDir,
	}

	if letter == 's' {
		config.Interval = time.Second * time.Duration(number)
	} else if letter == 'm' {
		config.Interval = time.Minute * time.Duration(number)
	} else if letter == 'h' {
		config.Interval = time.Hour * time.Duration(number)
	} else {
		return nil, fmt.Errorf("%c not supported in %s", letter, intervalString)
	}

	return config, nil
}
