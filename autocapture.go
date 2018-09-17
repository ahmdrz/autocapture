package main

import (
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {
	cfg, err := parseFlags()
	if err != nil {
		log.Fatal(err)
	}

	err = os.MkdirAll(cfg.OuputDir, 0755)
	if err != nil {
		log.Fatal(err)
	}

	death := make(chan os.Signal, 1)
	signal.Notify(death, os.Interrupt, os.Kill)

	for {
		select {
		case <-time.NewTicker(cfg.Interval).C:
			hour := time.Now().Hour()
			if hour >= cfg.StartHour && hour < cfg.FinishHour {
				err = takeScreenshot(cfg.OuputDir)
				if err != nil {
					log.Println(err)
				}
			}
		case <-death:
			log.Println("Signal received !")
			return
		}
	}
}
