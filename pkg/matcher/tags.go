package matcher

import (
	"regexp"
)

// Tags holds tags data
type Tags map[string]string

// Tagger is interfce which can discover file tags
type Tagger interface {
	For(item *ScrapItem) *Tags
}

var tagReplace = regexp.MustCompile(`(:\w*)`)

// TagsToPath interpolates tags from pattern into a new string
func TagsToPath(pattern string, tags Tags) (string, bool) {
	allfound := true
	res := tagReplace.ReplaceAllFunc([]byte(pattern), func(s []byte) []byte {
		tag, ok := tags[string(s[1:])]
		allfound = allfound && ok
		return []byte(tag)
	})
	return string(res), allfound
}
