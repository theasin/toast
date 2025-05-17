package toast

import (
	"os"

	uuid "github.com/nu7hatch/gouuid"
)

// Copies an image to `%userprofile%\AppData\Local\Temp` to be easily accessible

var lastTmpFile = ""

func copyBytesTemp(data []byte) (string, error) {
	dir, err := os.UserCacheDir()
	if err != nil {
		return "", err
	}
	id, _ := uuid.NewV4()
	out_path := dir + "/temp/" + id.String()
	err = os.WriteFile(out_path, data, 0600)
	if err != nil {
		return "", err
	}
	lastTmpFile = out_path
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
func deleteLastTmpFile() {
	os.Remove(lastTmpFile)
}
