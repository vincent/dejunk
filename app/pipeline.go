package app

import (
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"github.com/vincent/godejunk/pkg/artwork"
	"github.com/vincent/godejunk/pkg/matcher"
	"github.com/vincent/godejunk/pkg/pipe"
	"github.com/vincent/godejunk/pkg/rollback"
	"github.com/vincent/godejunk/pkg/writer"
)

var dummy = matcher.DummyTagger{}
var common = matcher.CommonTagger{}

// NewScrapperPipe provides a scrapper pipe to fill in
func NewScrapperPipe(store *writer.Store, rollback *rollback.RollbackFile) *pipe.Pipe {
	pipeline := &pipe.Pipe{
		Items: make(chan *matcher.ScrapItem, 2),
		Done:  make(chan int),
	}

	var artworker = artwork.NewArtworkFinder()

	go func() {
		defer close(pipeline.Done)
		defer artworker.CleanUp()
		for item := range pipeline.Items {
			log.Debug("scrapping", item.SourcePath)

			// Initialize empty tags
			item.Tags = &matcher.Tags{}

			// Is configured, fill with dummy tags
			if contains(item.Rule.With, "dummy") {
				item.Tags = dummy.For(item)
			}

			// Fill with common tags
			item.Tags = common.For(item)

			// Abort if we could not use all mandatory tags
			ok := item.EvaluateStorePath()
			if ok {
				if err := store.ImportItem(item); err == nil {
					log.Debug("wrote", item.StorePath)

					// Find cover
					tags := *item.Tags
					if contains(item.Rule.With, "artwork") {
						artwork := artworker.FindCover(item.Rule.Type, []string{tags["title"]})
						if artwork != "" {
							if err := store.Copy(artwork, filepath.Join(item.StoreDir, "cover"+filepath.Ext(artwork))); err != nil {
								log.Warn("cannot use artwork: ", err)
							}
						} else {
							log.Debug("found no cover for ", tags["title"])
						}
					}

					// Find background
					if contains(item.Rule.With, "background") {
						artwork := artworker.FindBackground(item.Rule.Type, []string{tags["title"]})
						if artwork != "" {
							if err := store.Copy(artwork, filepath.Join(item.StoreDir, "background"+filepath.Ext(artwork))); err != nil {
								log.Warn("cannot use background: ", err)
							}
							log.Warn("found ", artwork)
						} else {
							log.Debug("found no background for ", tags["title"])
						}
					}

					if rollback.Enabled {
						rollback.Write(item)
					}
				}
			} else {
				log.Println("cannot interpolate all tags for", item.SourcePath)
			}
		}
	}()

	return pipeline
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
