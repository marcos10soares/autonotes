package notes

import (
	"fmt"
	"marcos10soares/autonotes/pkg/utils"
	"os"

	"github.com/asticode/go-astisub"
)

// Generate generate md file from screenshot folder and srt file from otter.ai
// assumes both file and folder have the same name (without extension)
func Generate(fileAndFolderName string) error {
	notes, _ := astisub.OpenFile("input/" + fileAndFolderName + ".srt")
	images := utils.GetListOfImagesInFolder("output/" + fileAndFolderName)
	mdText := ""

	for _, note := range notes.Items {
		// mdText += note.String() + " "
		for _, line := range note.Lines {
			// fmt.Println(line.String())
			text := line.String() + " "
			newText := ""

			for i, char := range text {
				if char == '.' {
					newText = text[:i+1] + "\n\n" + text[i+1:]
				}
			}

			if newText == "" {
				mdText += text
			} else {
				mdText += newText
			}
		}

		for len(images) > 0 {
			if images[0].StartAt < note.StartAt {
				mdText += fmt.Sprintf("\n![%s](%s)\n", images[0].Name, images[0].Name)
				images = images[1:]
			} else {
				break
			}
		}

	}

	// add remaining images
	for len(images) > 0 {
		mdText += fmt.Sprintf("\n![%s](%s)\n", images[0].Name, images[0].Name)
		images = images[1:]
	}

	f, err := os.Create("output/" + fileAndFolderName + "/" + fileAndFolderName + ".md")
	if err != nil {
		return fmt.Errorf("Could not save to file %s, error: %v", fileAndFolderName, err)
	}
	_, err = f.WriteString(mdText)
	f.Close()

	return err
}
