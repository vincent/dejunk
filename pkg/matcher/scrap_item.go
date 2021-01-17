package matcher

import "path/filepath"

// ScrapItem is an item to be scrapped
type ScrapItem struct {
	SourcePath string
	StorePath  string
	StoreDir   string
	Rule       *Rule
	Tags       *Tags
}

// NewScrapItem producce a new ScrapItem
func NewScrapItem(filename string) *ScrapItem {
	return &ScrapItem{
		SourcePath: filename,
		Rule:       nil,
	}
}

// EvaluateStorePath set the StorePath and return true if all tags has been used
func (item *ScrapItem) EvaluateStorePath() bool {
	// Abort if we could not use all mandatory tags
	path, ok := TagsToPath(item.Rule.Store, *item.Tags)
	if ok {
		path = filepath.Join(item.Rule.Name, path+(*item.Tags)["extension"])
		item.StoreDir = filepath.Dir(path)
		item.StorePath = path
		return true
	}
	return false
}
