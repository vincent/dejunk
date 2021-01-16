package writer

import (
	"os"
	"strings"
)

// Store is a writable store
type Store struct {
	Count int
	Tree  Tree
}

// NewStore returns a new writable store
func NewStore() Store {
	return Store{
		Count: 0,
		Tree:  Tree{},
	}
}

// Write a new item in store
func (store *Store) Write(item string) bool {
	store.Count++

	parts := strings.Split(item, string(os.PathSeparator))
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

	dest[file] = nil
	return true
}
