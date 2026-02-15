package baseutils

import (
	"encoding/json"
	"os"
)

// ParseJSONFile
// open the passed file from path and try to unmarshal its contents into the passed interface
func ParseJSONFile(path string, data any) (err error) {

	// try to open the JSON file
	file, err := os.ReadFile(path)
	if err != nil {
		return
	}

	// parse JSON file to interface
	return json.Unmarshal([]byte(file), data)
}
