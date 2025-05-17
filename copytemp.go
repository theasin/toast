package toast

import (
	"os"
	"path/filepath"

	uuid "github.com/nu7hatch/gouuid"
)

// Copies an image to `%userprofile%\AppData\Local\Temp` to be easily accessible

var lastTmpFiles = []string{}

func copyBytesTemp(data []byte) (string, error) {
	dir, err := os.UserCacheDir()
	if err != nil {
		return "", err
	}
	id, _ := uuid.NewV4()
	out_path := dir + "\\temp\\" + id.String()
	err = os.WriteFile(out_path, data, 0600)
	if err != nil {
		return "", err
	}
	lastTmpFiles = append(lastTmpFiles, out_path)
	return out_path, nil
}

func copyFileTemp(in_path string) (string, error) {
	file, err := os.ReadFile(in_path)
	if err != nil {
		return "", err
	}
	return copyBytesTemp(file)
}

// yay dysfunctional programming!
func deleteLastTmpFiles() {
	dir, err := os.UserCacheDir()
	if err != nil {
		return
	}
	for _, v := range lastTmpFiles {
		if filepath.Dir(v) != (dir + "\\temp") {
			// not our files!
			continue
		}
		os.Remove(v)
	}
	lastTmpFiles = []string{}
}
