package datacore

import (
	"io/ioutil"
	"os"
)

// FileSystemDriver manage the read and write operations.
type FileSystemDriver struct {
	Path string
}

// WriteTextFile create and save a text file in the persistent storage, if something is wrong return an error.
func (fs *FileSystemDriver) WriteTextFile(filename string, content string) error {
	file, err := os.Create(fs.Path + filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, writeErr := file.WriteString(content)

	file.Sync()

	return writeErr
}

// ReadTextFile reads a text file and copy its content to memory, if something is wrong return an error.
func (fs *FileSystemDriver) ReadTextFile(filename string) (string, error) {
	content, err := ioutil.ReadFile(fs.Path + filename)
	return string(content), err
}
