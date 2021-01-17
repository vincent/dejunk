package writer

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// Tree is a faked fs
type Tree map[string]Tree

// Fprint prints the tree
func (tree Tree) Fprint(w io.Writer, root bool, padding string) {
	if tree == nil {
		return
	}

	index := 0
	for k, v := range tree {
		fmt.Fprintf(w, "%s%s\n", padding+getPadding(root, getBoxType(index, len(tree))), k)
		v.Fprint(w, false, padding+getPadding(root, getBoxTypeExternal(index, len(tree))))
		index++
	}
}

// Store a new item in summary FS
func (self *Tree) Write(path string) bool {
	parts := strings.Split(path, string(os.PathSeparator))
	dir := parts[:len(parts)-1]
	file := parts[len(parts)-1]

	// place on root
	dest := *self

	// walk down the tree
	for _, s := range dir {
		found, ok := dest[s]
		if !ok {
			dest[s] = Tree{}
			found = dest[s]
		}
		dest = found
	}

	_, exists := dest[file]
	dest[file] = nil

	return !exists
}

type boxType int

const (
	regular boxType = iota
	last
	afterLast
	between
)

func (boxType boxType) String() string {
	switch boxType {
	case regular:
		return "\u251c" // ├
	case last:
		return "\u2514" // └
	case afterLast:
		return " "
	case between:
		return "\u2502" // │
	default:
		panic("invalid box type")
	}
}

func getBoxType(index int, len int) boxType {
	if index+1 == len {
		return last
	} else if index+1 > len {
		return afterLast
	}
	return regular
}

func getBoxTypeExternal(index int, len int) boxType {
	if index+1 == len {
		return afterLast
	}
	return between
}

func getPadding(root bool, boxType boxType) string {
	if root {
		return ""
	}

	return boxType.String() + " "
}
