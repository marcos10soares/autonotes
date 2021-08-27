//+build mage

// Mage is a make-like command runner.  See https://magefile.org for full docs.
package main

import (
	"fmt"
	"marcos10soares/autonotes/pkg/notes"
	"marcos10soares/autonotes/pkg/screen"
	"marcos10soares/autonotes/pkg/utils"
	"time"

	"github.com/fatih/color"
	"github.com/magefile/mage/mg"
)

type Screen mg.Namespace
type Notes mg.Namespace

// Test screen capture all screens and saves them to a "screen_test" folder with the number of the screen
func (Screen) Test() error {
	utils.PrintTitle("Testing Screens")
	screen.Test(screen.JPEG)

	return nil
}

// Capture screen - usage: mage screen:capture <start_delay_seconds> <screen_index> <capture_interval_ms> <jpeg_or_png> <output_folder_name>
func (Screen) Capture(startDelay int, screenIndex int, captureInterval int, filetype string, folder string) error {
	utils.PrintTitle(fmt.Sprintf("Capturing Screen #%d", screenIndex))

	if captureInterval == 0 {
		captureInterval = 5000 // default value 5s in milliseconds
	}

	imgFileType, err := screen.ConvertStringToImgfileType(filetype)
	if err != nil {
		return err
	}

	screen.Capture(startDelay, time.Duration(captureInterval), screenIndex, *imgFileType, folder)

	return nil
}

// Generate .md file from screenshot folder and .srt file from otter.ai - usage: mage generate <file_and_folder_name_must_have_the_same_name> <obsidianFormat_true_false>
func (Notes) Generate(fileAndFolderName string, obsidianFormat bool) error {
	utils.PrintTitle("Generating Notes")

	err := notes.Generate(fileAndFolderName, obsidianFormat)
	if err != nil {
		return err
	}
	cyan := color.New(color.FgCyan).SprintFunc()
	fmt.Println(cyan(fmt.Sprintf("Successfully generated file: output/%s/%s.md", fileAndFolderName, fileAndFolderName)))
	return nil
}
