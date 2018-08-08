package dataset

import (
	"io"
	"os"
	"path"
)

// Opener is.
type Opener interface {
	GetFile(root string, relPath string, readOnly bool, binary bool) (io.ReadWriteCloser, error)
	GetDir(root string, relPath string, readOnly bool) ([]string, error)
}

// DefaultOpener is.
type DefaultOpener struct{}

// GetFile is.
func (DefaultOpener) GetFile(root string, relPath string, readOnly bool, binary bool) (io.ReadWriteCloser, error) {
	path := path.Join(root, relPath)
	return os.Open(path)
}

// GetDir is.
func (DefaultOpener) GetDir(root string, relPath string, readOnly bool) ([]string, error) {
	path := path.Join(root, relPath)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return file.Readdirnames(0)
}
