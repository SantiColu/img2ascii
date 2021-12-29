package main

import (
	"flag"
	"fmt"
	"image"
	imgColor "image/color"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"strconv"

	"github.com/gookit/color"
	"github.com/nfnt/resize"
)

func resizeImage(img image.Image, w int) (image.Image, int, int) {
	var maxs image.Point = img.Bounds().Max
	h := int(float64(w) * (float64(maxs.Y*10) / float64(maxs.X*20))) // 10 y 16 para mantener relacion aspecto por caracteres
	newImg := resize.Resize(uint(w), uint(h), img, resize.Lanczos3)
	return newImg, w, h
}

func getPixelColor(img image.Image, x, y int) color.RGBColor {
	r, g, b, _ := img.At(x, y).RGBA()
	dR := uint8(r / 257)
	dG := uint8(g / 257)
	dB := uint8(b / 257)
	return color.RGB(dR, dG, dB)
}

func getCharacter(img image.Image, x, y int) string {
	asciiCharacters := [17]string{"M", "N", "D", "8", "O", "Z", "$", "7", "I", "?", "+", "=", "~", ":", ",", ".", "."}
	g := imgColor.GrayModel.Convert(img.At(x, y))
	gScale := g.(imgColor.Gray).Y
	pos := int(int(gScale) * 16 / 255)
	return asciiCharacters[pos]
}

func main() {
	useColors := flag.Bool("colors", false, "Mo mostrar colores")
	flag.Parse()
	fromIndex := 1
	if *useColors {
		fromIndex = 2
	}
	arguments := os.Args[fromIndex:]
	var imgPath string = "./image.png"
	var width int = 100

	if len(arguments) < 1 {
		fmt.Println("Es necesario un archivo")
		os.Exit(1)
	}

	if arguments[0] != "" {
		imgPath = arguments[0]
	}

	if len(arguments) > 1 && arguments[1] != "" {
		nWidth, err := strconv.Atoi(arguments[1])
		if err != nil {
			fmt.Println("El segundo argumento debe ser un numero")
			os.Exit(1)
		}
		width = nWidth
	}

	file, err := os.Open(imgPath)

	if err != nil {
		fmt.Println("Error buscando el archivo \n -> ", err.Error())
		os.Exit(1)
	}

	img, _, err := image.Decode(file)

	if err != nil {
		fmt.Println("Error decodificando el archivo \n -> ", err.Error())
		os.Exit(1)
	}

	reImg, w, h := resizeImage(img, width)

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if *useColors {
				getPixelColor(reImg, x, y).Printf("#")
			} else {
				fmt.Print(getCharacter(reImg, x, y))
			}
		}
		fmt.Println()
	}
}
