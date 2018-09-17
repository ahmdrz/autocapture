package main

import (
	"flag"
	"fmt"
	"path"
	"strconv"
	"time"
)

type Config struct {
	Interval   time.Duration
	OuputDir   string
	StartHour  int
	FinishHour int
}

func parseFlags() (*Config, error) {
	var (
		intervalString string
	)

	config := &Config{}

	flag.StringVar(&intervalString, "interval", "15s", "Interval to capture screenshot. Ex: 15s means capture every 15 seconds.")
	flag.StringVar(&config.OuputDir, "output", path.Join(".", "captures"), "Directory to save screenshots.")
	flag.IntVar(&config.StartHour, "start-hour", 0, "Start capture from which hour ? 0-24")
	flag.IntVar(&config.FinishHour, "finish-hour", 24, "Finish capture from which hour ? 0-24")
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
