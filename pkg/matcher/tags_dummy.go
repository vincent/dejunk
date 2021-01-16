package matcher

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

// DummyTagger is a dummy tagger
type DummyTagger struct{}

// For return dummy tags for the given item
func (tagger *DummyTagger) For(item *ScrapItem) *Tags {
	tags := *item.Tags
	dirname := filepath.Dir(item.SourcePath)
	basename := filepath.Base(item.SourcePath)
	extension := filepath.Ext(basename)
	title := basename[0 : len(basename)-len(extension)]
	var setag string
	var ok bool

	// Find season & episode from filename
	if item.Rule.Type == Music {
		files, _ := ioutil.ReadDir(dirname)
		siblings := len(files)
		if siblings > 1 && siblings < 100 {
			tags["artist"] = strings.Title(strings.ToLower(filepath.Base(dirname)))
			tags["album"] = "Unknown Album"
		}
	}

	// Find season & episode from filename
	if item.Rule.Type == TVShow {
		setag, ok = fillSeasonEpisodeTagsFromName(basename, tags)
		if !ok {
			// Find season & episode from parent name
			setag, _ = fillSeasonEpisodeTagsFromName(dirname, tags)
		}
	}

	// TODO: Find year from filename
	ok = fillYearTagFromName(basename, tags)
	if !ok {
		// Find season & episode from parent name
		ok = fillYearTagFromName(dirname, tags)
		if !ok {
			tags["year"] = "unknown"
		}
	}

	// TODO: Find title by removing junk & year from filename
	_, ok = tags["episode"]
	if ok {
		title = strings.ReplaceAll(title, setag, "")
	}
	title = strings.ReplaceAll(title, tags["year"], " ")
	title = RECleanJunk.ReplaceAllString(title, " ")
	title = RECleanSpaces.ReplaceAllString(title, " ")
	title = strings.TrimSpace(strings.Title(strings.ToLower(title)))

	tags["title"] = title
	tags["extension"] = extension

	return &tags
}

func fillSeasonEpisodeTagsFromName(name string, tags Tags) (string, bool) {
	res := RESeasonEpisode.FindStringSubmatch(name)
	if len(res) == 5 {
		tags["season"] = fmt.Sprintf("%s", res[2])
		tags["episode"] = fmt.Sprintf("%02s", res[4])
		return res[0], true
	}
	return "", false
}

func fillYearTagFromName(name string, tags Tags) bool {
	res := REYear.FindStringSubmatch(name)
	if len(res) == 3 {
		tags["year"] = fmt.Sprintf("%s", res[1])
		return true
	}
	return false
}
