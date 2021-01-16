package pipe

import (
	"github.com/vincent/godejunk/pkg/config"
	"github.com/vincent/godejunk/pkg/matcher"
)

// Pipe holds a pipe
type Pipe struct {
	Config *config.Config
	Items  chan *matcher.ScrapItem
	Done   chan int
}
