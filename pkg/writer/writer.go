package writer

import (
	"os"
	"path/filepath"
	"strings"

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

// Write a new item in store
func (store *Store) Write(item *matcher.ScrapItem) bool {
	store.Count++

	parts := strings.Split(item.StorePath, string(os.PathSeparator))
	path := parts[:len(parts)-1]
	file := parts[len(parts)-1]

	// place on root
	dest := store.Tree

	// walk down the tree
	for _, s := range path {
		found, ok := dest[s]
		if !ok {
			dest[s] = Tree{}
			found = dest[s]
		}
		dest = found
	}

	if _, ok := dest[file]; ok {
		return false
	}

	if !store.Dry {
		fullPath := filepath.Join(store.Output, item.StorePath)
		err := os.MkdirAll(filepath.Dir(fullPath), os.ModePerm)
		if err != nil {
			panic(err)
		}

		err = os.Rename(item.SourcePath, fullPath)
		if err != nil {
			panic(err)
		}
	}

	dest[file] = nil
	return true
}

func (store *Store) Move(from string, to string) error {
	store.Count++
	if !store.Dry {
		err := os.Rename(from, to)
		if err != nil {
			return err
		}
	}
	return nil
}
