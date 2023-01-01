package cache

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type FileStorage struct {
	directory string
}

func (f FileStorage) fullKey(id string) string {
	h := sha256.New()
	_, _ = h.Write([]byte(id))
	hash := hex.EncodeToString(h.Sum(nil))
	return filepath.Join(f.directory, fmt.Sprintf("%s.cache", hash))
}
func (f *FileStorage) Store(id string, content []byte) error {
	var err error = nil
	err = ioutil.WriteFile(f.fullKey(id), content, 0666)
	return err
}

func (f *FileStorage) FetchOne(id string) ([]byte, error) {
	var err error = nil
	b, err := ioutil.ReadFile(f.fullKey(id))
	return b, err
}

func (f *FileStorage) KeyExists(id string) (bool, error) {
	_, err := os.Stat(f.fullKey(id))

	if err != nil {

		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil

}

func NewFileStorage(directory string) *FileStorage {
	return &FileStorage{directory: directory}
}

var _ Cache = (*FileStorage)(nil)
