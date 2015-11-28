package genetic

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

/**
Color defines one of Red Green Blue or Alpha
*/
type Color int

/**
These variables define public access to the target array generated from
the image file.
*/
var (
	Target       [][]uint8
	TargetWidth  int
	TargetHeight int
)

const (
	RED   Color = 1
	GREEN Color = 2
	BLUE  Color = 3
	ALPHA Color = 4
)

/**
SetTarget takes a file name as input and configures the target
data that population Members will be compared against. It returns
the dimensions of the target array.
*/
func SetTarget(fileName string) error {
	//open the file
	file, err := os.Open(fileName)
	if err != nil {
		//abort if error
		return err
	}
	//decode the file as png
	image, err := png.Decode(file)
	if err != nil {
		//abort if error
		return err
	}
	//determine information about file
	fmt.Print("Image Color model: ")
	switch image.ColorModel() {
	case color.AlphaModel:
		fmt.Println("AlphaModel")
	case color.RGBAModel:
		fmt.Println("RGBAModel")
	default:
		fmt.Println("Other Model")
	}
	//determine bounds
	bounds := image.Bounds()
	width := bounds.Max.X - bounds.Min.X
	height := bounds.Max.Y - bounds.Min.Y

	//configure target data
	setUpTarget(height, width)
	fmt.Printf("Image Bounds: %s\n", bounds)
	fmt.Printf("Image At(0,0): %s\n", image.At(bounds.Min.X, bounds.Min.Y))

	fmt.Printf("Target dimensions: width=%d height=%d\n", width, height)

	populateTargetFromImage(image)
	return nil
}

/**
setUpTarget initializes the target array to the dimensions of the image
*/
func setUpTarget(targetHeight, targetWidth int) {
	Target = make([][]uint8, targetHeight)
	TargetHeight = targetHeight
	TargetWidth = targetWidth
	for i := range Target {
		Target[i] = make([]uint8, targetWidth)
	}
}

/**
Populate the contents of the Target with data from the Image
*/
func populateTargetFromImage(image image.Image) {
	var color color.Color
	var red, green, blue, alpha uint32
	for y := range Target {
		for x := range Target[y] {
			color = image.At(x, y)
			red, green, blue, alpha = color.RGBA()
			Target[y][x] = uint8(red)
			fmt.Printf("Extracted rgba(%d,%d,%d,%d)\n", red, green, blue, alpha)
		}
	}
}

/**
WriteMonochromePNG creates a png file with the given name out of the given data array.
It interprets the data array as the provided Color in the RGBA model
*/
func WriteMonochromePNG(fileName string, data [][]uint8, color Color) error {
	switch color {
	case RED:
		return writeRedPNG(fileName, data)
	case GREEN:
		//return writeGreenPNG(fileName, data)
	case BLUE:
		//return writeBluePNG(fileName, data)
	case ALPHA:
		//return writeAlphaPNG(fileName, data)
	default:
		fmt.Println("Unknown color type.")
	}
	return nil
}

/**
writeRedPNG creates a png file with the contents of the data array as its RED values
and with GREEN and BLUE at zero. It's alpha will be maximum for all values.
*/
func writeRedPNG(fileName string, data [][]uint8) error {
	rect := image.Rect(0, 0, TargetWidth, TargetHeight)
	img := image.NewRGBA(rect)
	for y := range data {
		for x := range data[y] {
			img.SetRGBA(x, y, color.RGBA{data[y][x], 0, 0, 255})
		}
	}
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	err = png.Encode(file, img)
	if err != nil {
		return err
	}
	err = file.Close()
	if err != nil {
		return err
	}
	return nil
}
