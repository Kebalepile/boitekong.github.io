package pipeline

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Kebalepile/job_board/spiders/types"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

// saves scraped data into a json file in database public folder
func GovPageFile(data *types.Links) error {

	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)

	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "  ")

	err := encoder.Encode(*data)

	if err != nil {
		return err
	}
	title := cleanStr(data.Title)
	filePath := filepath.Join("database", "public", fmt.Sprintf("%s.json", title))

	err = os.WriteFile(filePath, buffer.Bytes(), 0644)
	if err != nil {
		return err
	}
	log.Print(data.Title, " Saved at ", filePath)
	return nil
}

// saves scraped data into a json file in database private folder
func HeithaFile(data *types.HeithaJobs) error {

	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)

	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "  ")

	err := encoder.Encode(*data)
	if err != nil {
		return err
	}
	title := cleanStr(data.Title)
	filePath := filepath.Join("database", "private", fmt.Sprintf("%s.json", title))

	err = os.WriteFile(filePath, buffer.Bytes(), 0644)
	if err != nil {
		return err
	}
	log.Print(data.Title, " Saved at ", filePath)
	return nil
}

// saves scraped data into a json file in database private folder
func ProPersonnelFile(data *types.ProPersonnelJobs) error {

	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)

	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "  ")

	err := encoder.Encode(*data)
	if err != nil {
		return err
	}
	title := cleanStr(data.Title)
	filePath := filepath.Join("database", "private", fmt.Sprintf("%s.json", title))

	err = os.WriteFile(filePath, buffer.Bytes(), 0644)
	if err != nil {
		return err
	}
	log.Print(data.Title, " Saved at ", filePath)
	return nil
}

// replaces all `,` and spaces in s with `-`
func cleanStr(s string) string {
	re := regexp.MustCompile("[, ]")
	return re.ReplaceAllString(s, "-")
}
