package main

import (
	"flag"
	"fmt"
	img "image"
	"image/color"
	"log"

	cv "gocv.io/x/gocv"
)

func main() {
	fmt.Println("Face Detection using gocv")
	file := flag.String("image", "", "use -image and the path of the image to give the input")
	flag.Parse()
	classifier := cv.NewCascadeClassifier()
	if !classifier.Load("models/cascade_classifier.xml") {
		log.Fatal("Failed to load the classifier file")
	}
	defer classifier.Close()

	image := cv.IMRead(*file, cv.IMReadAnyColor)
	gray := cv.NewMat()
	defer gray.Close()
	// Changing the color mode into gray for ease of classification.
	cv.CvtColor(image, &gray, cv.ColorBGRToGray)

	faces := classifier.DetectMultiScale(gray)
	for _, rect := range faces {
		// Dimensions
		x := rect.Min.X
		y := rect.Min.Y
		w := rect.Dx()
		h := rect.Dy()

		rectangle := img.Rect(x, y, x+w, y+h)
		red := color.RGBA{
			255, 0, 0, 1.0,
		}
		cv.Rectangle(&image, rectangle, red, 1)
		// Writing image into a file.
		if !cv.IMWrite("output.jpg", image) {
			log.Fatal("Failed to write the output image")
		}
		window := cv.NewWindow("Face Detection")
		window.IMShow(image)
		window.WaitKey(10000)
	}
}
