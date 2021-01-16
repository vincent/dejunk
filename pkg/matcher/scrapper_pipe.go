package matcher

import (
	log "github.com/sirupsen/logrus"
	"github.com/vincent/godejunk/pkg/writer"
)

// Pipe holds a pipe
type Pipe struct {
	Items chan *ScrapItem
	Done  chan int
}

var dummy = DummyTagger{}
var common = CommonTagger{}

// NewScrapperPipe provides a scrapper pipe to fill in
func NewScrapperPipe(store *writer.Store) *Pipe {
	pipe := &Pipe{
		Items: make(chan *ScrapItem, 2),
		Done:  make(chan int),
	}

	go func() {
		defer close(pipe.Done)
		for item := range pipe.Items {
			log.Println("scrapping", item.SourcePath)

			// Initialize empty tags
			item.Tags = &Tags{}

			// Is configured, fill with dummy tags
			if contains(item.Rule.With, "dummy") {
				item.Tags = dummy.For(item)
			}

			// Fill with common tags
			item.Tags = common.For(item)

			// Abort if we could not use all mandatory tags
			ok := item.EvaluateStorePath()
			if ok {
				store.Write(item.StorePath)
			} else {
				log.Println("cannot interpolate all tags for", item.SourcePath)
			}
		}
	}()

	return pipe
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
