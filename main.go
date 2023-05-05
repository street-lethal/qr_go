package main

import (
	"fmt"
	"image/color"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

const (
	StdOutBlack = "\033[40m  \033[0m"
	StdOutWhite = "\033[47m  \033[0m"
)

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	message := string(bytes)

	qrCode, err := qr.Encode(message, qr.M, qr.Auto)
	if err != nil {
		panic(err)
	}

	printQrCode(qrCode)
}

func printQrCode(qrCode barcode.Barcode) {
	rect := qrCode.Bounds()
	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			if qrCode.At(x, y) == color.Black {
				fmt.Print(StdOutBlack)
			} else {
				fmt.Print(StdOutWhite)
			}
		}
		fmt.Println()
	}
}
