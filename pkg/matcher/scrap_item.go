package matcher

// ScrapItem is an item to be scrapped
type ScrapItem struct {
	Filename string
	Rule     *Rule
	Tags     *Tags
}

// NewScrapItem producce a new ScrapItem
func NewScrapItem(filename string) *ScrapItem {
	return &ScrapItem{
		Filename: filename,
		Rule:     nil,
	}
}
