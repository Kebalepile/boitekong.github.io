package pipeline

import (
	"encoding/json"
	"fmt"
	"github.com/Kebalepile/job_board/spiders/types"
	"log"
	"os"
	"path/filepath"
)

func SaveJsonFile(data types.Links) error {
	contentBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	filePath := filepath.Join("database", "public", fmt.Sprintf("%s.json", data.Title))

	err = os.WriteFile(filePath, contentBytes, 0644)
	if err != nil {
		return err
	}
	log.Print(data.Title, " Saved at ", filePath)
	return nil
}
