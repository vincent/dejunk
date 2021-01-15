package matcher

import (
	"fmt"
	"os"
	"path/filepath"

	gotags "github.com/dhowden/tag"
	log "github.com/sirupsen/logrus"
)

// CommonTagger is an multi-purpose tagger
type CommonTagger struct{}

// For return id3 tags for the given item
func (tagger *CommonTagger) For(item *ScrapItem) *Tags {
	tags := *(item.Tags)
	extension := filepath.Ext(item.Filename)
	var v string

	tags["extension"] = extension

	handle, err := os.Open(item.Filename)
	if err != nil {
		log.Println("cannot open", item.Filename)
		return &tags
	}
	id3Tags, err := gotags.ReadFrom(handle)
	if err != nil {
		log.Println("cannot parse tags from", item.Filename)
		return &tags
	}

	v = id3Tags.Title()
	if v != "" {
		tags["title"] = v
	}

	v = id3Tags.Artist()
	if v != "" {
		tags["artist"] = v
	}

	v = id3Tags.Album()
	if v != "" {
		tags["album"] = v
	}

	v = id3Tags.AlbumArtist()
	if v != "" {
		tags["album_artist"] = v
	}

	v = id3Tags.Composer()
	if v != "" {
		tags["composer"] = v
	}

	v = id3Tags.Genre()
	if v != "" {
		tags["genre"] = v
	}

	v = fmt.Sprint(id3Tags.Year())
	if v != "" {
		tags["year"] = v
	}

	return &tags
}
