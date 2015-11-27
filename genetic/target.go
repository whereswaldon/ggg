package genetic

import (
	"fmt"
	//"image"
	"image/color"
	"image/png"
	"os"
)

/**
These variables define public access to the target array generated from
the image file.
*/
var (
	Target       [][]uint8
	TargetWidth  int
	TargetHeight int
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
	return nil
}

/**
setUpTarget initializes the target array to the dimensions of the image
*/
func setUpTarget(targetHeight, targetWidth int) {
	Target = make([][]uint8, targetHeight)
	TargetHeight = targetHeight
	TargetWidth = targetWidth
	for i := range target {
		target[i] = make([]uint8, targetWidth)
	}
}
