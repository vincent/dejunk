package writer

import (
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"

	"github.com/vincent/godejunk/pkg/matcher"
)

// Store is a writable store
type Store struct {
	Count  int
	Dry    bool
	Output string
	Tree   Tree
}

// NewStore returns a new writable store
func NewStore(output string, dry bool) *Store {
	return &Store{
		Count:  0,
		Dry:    dry,
		Tree:   Tree{},
		Output: output,
	}
}

// ImportItem into this store
func (store *Store) ImportItem(item *matcher.ScrapItem) error {
	store.Count++
	log.Debug("import item ", item.SourcePath)

	if !store.Tree.Write(item.StorePath) {
		return nil
	}

	if !store.Dry {
		fullPath := filepath.Join(store.Output, item.StorePath)
		err := os.MkdirAll(filepath.Dir(fullPath), os.ModePerm)
		if err != nil {
			return err
		}

		err = os.Rename(item.SourcePath, fullPath)
		if err != nil {
			return err
		}
	}

	return nil
}

// Copy a file into this store
func (store *Store) Copy(from string, to string) error {
	store.Count++
	log.Debug("copy from ", from, " to ", to)

	if !store.Tree.Write(to) {
		return nil
	}

	if !store.Dry {
		fullPath := filepath.Join(store.Output, to)
		err := os.MkdirAll(filepath.Dir(fullPath), os.ModePerm)
		if err != nil {
			return err
		}

		err = os.Rename(from, fullPath)
		if err != nil {
			return err
		}
	}

	return nil
}

// Restore a file to another location
func (store *Store) Restore(from string, to string) error {
	store.Count++
	log.Debug("restore ", from, " to ", to)

	if !store.Dry {
		err := os.Rename(from, to)
		if err != nil {
			return err
		}
	}
	return nil
}
