package screen

import (
	"fmt"
	"image/png"
	"marcos10soares/autonotes/pkg/utils"
	"os"

	"github.com/fatih/color"
	"github.com/kbinani/screenshot"
)

// ImgfileType for saving screenshots
type ImgfileType string

const (
	// JPEG value for screen functions
	JPEG ImgfileType = "jpeg"

	// PNG value for screen functions
	PNG ImgfileType = "png"
)

// ConvertStringToImgfileType converts user input to type ImgfileType
func ConvertStringToImgfileType(imgtype string) (*ImgfileType, error) {
	var filetype ImgfileType
	switch imgtype {
	case "jpeg":
		filetype = "jpeg"
		return &filetype, nil
	case "png":
		filetype = "png"
		return &filetype, nil
	default:
		return nil, fmt.Errorf("filetype not supported - supported types: jpeg | png")
	}
}

// Test captures all screens and saves them to an output folder
func Test(filetype ImgfileType) error {
	folder := "output/screen_test"

	err := utils.CreateDirectoryIfNotExists(folder)
	if err != nil {
		return err
	}

	n := screenshot.NumActiveDisplays()

	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)

		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			return err
		}

		fileName := fmt.Sprintf("%s/%d_%dx%d.%s", folder, i, bounds.Dx(), bounds.Dy(), filetype)
		file, err := os.Create(fileName)
		if err != nil {
			return err
		}

		defer file.Close()

		err = png.Encode(file, img)
		if err != nil {
			return err
		}

		cyan := color.New(color.FgCyan).SprintFunc()
		fileName = fmt.Sprintf("%s/%s_%dx%d.%s", folder, cyan(i), bounds.Dx(), bounds.Dy(), filetype)
		fmt.Printf("%s : %v \"%s\"\n", cyan(fmt.Sprintf("screen #%d", i)), bounds, fileName)
	}

	return nil
}
