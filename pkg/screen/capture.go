package screen

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"marcos10soares/autonotes/pkg/utils"
	"math"
	"os"
	"time"

	"github.com/kbinani/screenshot"
	"github.com/vitali-fedulov/images"
)

// Capture captures screen index
func Capture(startDelay int, captureInterval time.Duration, screen int, imgFileType ImgfileType, folder string) error {
	if startDelay != 0 {
		fmt.Println("Delay Before Starting")
		fmt.Print("Starting in ")
	}

	for ; startDelay > 0; startDelay-- {
		fmt.Printf("%d...", startDelay)
		time.Sleep(time.Second)
	}
	fmt.Print("starting!\n")
	// start counting
	startTime := time.Now()

	// init tmp image
	bounds := screenshot.GetDisplayBounds(screen)
	tmpImage, err := screenshot.CaptureRect(bounds)
	if err != nil {
		return err
	}
	// save first image
	saveImage(tmpImage, imgFileType, folder, startTime)

	for {
		currentImage, equal := sampleImageAndCompare(tmpImage, bounds)
		if !equal {
			fmt.Println("Images are distinct:")
			// to avoid capturing transitions, or ongoing drawings
			for moving := true; moving; {
				time.Sleep(time.Millisecond * 500) // sample images every half second

				fmt.Println("\tsampling for moving image")

				latestImage, equal := sampleImageAndCompare(currentImage, bounds)
				if equal {
					fmt.Println("\timage stopped moving.")
					currentImage = latestImage
					moving = false
				} else {
					currentImage = latestImage
				}
			}

			tmpImage = currentImage
			saveImage(currentImage, imgFileType, folder, startTime)
			fmt.Println("\tsaved image.")
		} else {
			fmt.Println("Images are similar.")
		}

		time.Sleep(time.Millisecond * captureInterval)
	}

}

// sampleImageAndCompare returns true if images are similar
func sampleImageAndCompare(img *image.RGBA, bounds image.Rectangle) (*image.RGBA, bool) {
	tmpImage, err := screenshot.CaptureRect(bounds)
	if err != nil {
		panic(err)
	}
	// Calculate hashes and image sizes.
	hashA, imgSizeA := images.Hash(img)
	hashB, imgSizeB := images.Hash(tmpImage)

	// Image comparison.
	if images.Similar(hashA, hashB, imgSizeA, imgSizeB) {
		return tmpImage, true
	}

	return tmpImage, false
}

func saveImage(img *image.RGBA, imgFileType ImgfileType, folder string, startTime time.Time) error {
	elapsed := time.Since(startTime)

	outputFolder := "output/" + folder

	// create folder if does not exist
	utils.CreateDirectoryIfNotExists(outputFolder)

	elapsedMin, _ := math.Modf(elapsed.Minutes())
	elapsedSec, _ := math.Modf(elapsed.Seconds())

	// add a timestamp for uniqueness
	fileName := fmt.Sprintf("%s/capture_%d_%d_%s_%s.%s", outputFolder, int(elapsedMin), int(elapsedSec)%60, folder, fmt.Sprint(time.Now().Unix()), imgFileType)
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// save as png or jpeg, defaults to jpeg
	switch imgFileType {
	case "png":
		// save as PNG
		err = png.Encode(file, img)
		if err != nil {
			return err
		}
	case "jpeg":
		// Save as JPEG - specify the quality, between 0-100, higher is better
		err = jpeg.Encode(file, img, &jpeg.Options{
			Quality: 80,
		})
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("error saving - unknown image format: %s", imgFileType)
	}

	return nil
}
