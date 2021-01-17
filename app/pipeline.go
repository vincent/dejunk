package app

import (
	log "github.com/sirupsen/logrus"
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

	go func() {
		defer close(pipeline.Done)
		for item := range pipeline.Items {
			log.Println("scrapping", item.SourcePath)

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
				if ok = store.Write(item); ok {
					log.Println("wrote", item.StorePath)

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
