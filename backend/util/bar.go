package util

import (
	"fmt"
	"time"

	"github.com/schollz/progressbar/v3"
)

type TimerBar struct {
	Header        string
	Width         int
	CurrentString string
}

func (bar *TimerBar) InitBar(header string) {
	bar.Header = header
	bar.Width = 100
}

func ShowBar(context string, start, width int) {
	bar := progressbar.NewOptions(width,
		progressbar.OptionUseANSICodes(true),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowBytes(false),
		progressbar.OptionSetWidth(30),
		progressbar.OptionSetPredictTime(false),
		progressbar.OptionSetElapsedTime(true),
		progressbar.OptionSetDescription(context),
		progressbar.OptionClearOnFinish(),
	)
	bar.Add(start)
	fmt.Printf("E: %s", time.Now().Add(time.Duration(width-start)*time.Second).Format(time.Kitchen))
	for i := start; i < width; i++ {
		bar.Add(1)
		time.Sleep(1 * time.Second)
	}
}
