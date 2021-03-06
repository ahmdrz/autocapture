package main

import (
	"fmt"
	"image/png"
	"os"
	"path"
	"time"

	"github.com/kbinani/screenshot"
)

func takeScreenshot(outputDir string) error {
	for i := 0; i < screenshot.NumActiveDisplays(); i++ {
		bounds := screenshot.GetDisplayBounds(i)

		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			return err
		}
		file, err := os.Create(path.Join(outputDir, fmt.Sprintf("%s (%d).png", time.Now().Format("2006_01_02 15_04_05"), i)))
		if err != nil {
			return err
		}
		defer file.Close()

		err = png.Encode(file, img)
		if err != nil {
			return err
		}
	}
	return nil
}
