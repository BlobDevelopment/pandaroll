package fs

import (
	"os"
)

func MkdirIfNotExists(name string) error {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		err := os.Mkdir(name, 0755)
		if err != nil {
			return err
		}
	}

	return nil
}

func CreateFile(name string, content string) error {
	return os.WriteFile(name, []byte(content), 0644)
}

func Exists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

func ReadFile(name string) (*string, error) {
	bytes, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}

	str := string(bytes)
	return &str, nil
}
