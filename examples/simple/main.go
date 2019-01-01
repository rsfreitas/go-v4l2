package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rsfreitas/go-v4l2"
)

type CliOptions struct {
	device   string
	frames   uint
	loopback string
}

func getOptions() CliOptions {
	var options CliOptions

	flag.StringVar(&options.loopback, "loopback", "",
		"Sets the loopback device to write frames.")

	flag.StringVar(&options.device, "device", "",
		"Sets the video device to capture frames.")

	flag.UintVar(&options.frames, "frames", 1,
		"The number of captured frames to save.")

	flag.Parse()

	return options
}

func run() {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

loop:
	for {
		select {
		case <-quit:
			break loop
		}
	}

	fmt.Println("Finishing loopback")
}

func main() {
	options := getOptions()

	if options.device == "" {
		fmt.Println("Need to pass the device name")
		os.Exit(-1)
	}

	device, err := v4l2.Open(v4l2.Options{
		Device:   options.device,
		Width:    640,
		Height:   480,
		Format:   v4l2.ImageFormatYUYV,
		LoopBack: options.loopback,
		Model:    v4l2.ModelUSBWebcam,
		Channel:  v4l2.ChannelComposite,
	})

	if err != nil {
		fmt.Println("1:", err)
		os.Exit(-1)
	}

	defer device.Close()

	if options.loopback != "" {
		loopback, err := v4l2.NewLoopback(options.loopback, device)

		if err != nil {
			fmt.Println("1.1:", err)
			os.Exit(-1)
		}

		defer loopback.Close()
		run()
	} else {
		time.Sleep(3 * time.Second)

		for i := uint(0); i < options.frames; i++ {
			img, err := device.Capture()

			if err != nil {
				fmt.Println("2:", err)
				os.Exit(-1)
			}

			if err := img.ToJpeg("image.jpg"); err != nil {
				fmt.Println("3:", err)
				os.Exit(-1)
			}
		}

		time.Sleep(3 * time.Second)
	}

	fmt.Println("Finished")
}
