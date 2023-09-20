<<<<<<< HEAD
package pipeline

import (
	"encoding/json"
	"fmt"
	"github.com/Kebalepile/job_board/spiders"
	"log"
	"os"
	"path/filepath"
)

func SaveJsonFile(data spiders.Links) error {
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
=======
package pipeline

import (
	"encoding/json"
	"fmt"
	"github.com/Kebalepile/job_board/spiders"
	"log"
	"os"
	"path/filepath"
)

func SaveJsonFile(data spiders.Links) error {
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
>>>>>>> 803aef7a2dea977e110d749638eece7dbcfda347
