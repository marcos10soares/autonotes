package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
)

// ImageCapture is a struct for image captures
type ImageCapture struct {
	Name    string        // filename
	StartAt time.Duration //timestamp
}

// GetListOfImagesInFolder fetches a slice of image captures from folder
func GetListOfImagesInFolder(folder string) ([]ImageCapture, error) {
	var images []ImageCapture

	// ignoreFiles := []string{
	// 	folder,
	// 	".DS_Store",
	// }

	// only finds files that start with "capture"
	err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if strings.HasPrefix(info.Name(), "capture") {
			arr := strings.Split(info.Name(), "_")

			min, err := strconv.Atoi(arr[1])
			if err != nil {
				return err
			}
			sec, err := strconv.Atoi(arr[2])
			if err != nil {
				return err
			}

			duration, err := time.ParseDuration(fmt.Sprintf("%dm%ds", min, sec))
			if err != nil {
				return err
			}

			images = append(images, ImageCapture{
				Name:    info.Name(),
				StartAt: duration,
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	// sort
	sort.Slice(images[:], func(i, j int) bool {
		return images[i].StartAt < images[j].StartAt
	})

	return images, err
}

// StringSliceContains checks if a string slice contains a string
func StringSliceContains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// CreateDirectoryIfNotExists creates a directory if it does not exist
func CreateDirectoryIfNotExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.Mkdir(path, os.ModeDir|0755)
	}
	return nil
}

// PrintTitle prints a title
func PrintTitle(a ...interface{}) {
	c := color.New(color.FgGreen).Add(color.Underline)
	c.Println(a)
}
