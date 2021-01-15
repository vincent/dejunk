package walker

import (
	"github.com/karrick/godirwalk"
	log "github.com/sirupsen/logrus"
	"github.com/vincent/godejunk/pkg/matcher"
)

// Walker blah blah
type Walker interface {
	WalkDirectory(osDirname string, pipe *matcher.Pipe) error
}

// FilesWalker blah blah
type FilesWalker struct {
	Matcher *matcher.Matcher
}

// NewWalker returns a walker
func NewWalker(matcher *matcher.Matcher) Walker {
	return &FilesWalker{
		Matcher: matcher,
	}
}

// WalkDirectory scans files to dejunk
func (walker *FilesWalker) WalkDirectory(osDirname string, pipe *matcher.Pipe) error {
	defer close(pipe.Items)
	return godirwalk.Walk(osDirname, &godirwalk.Options{
		Unsorted: true,
		Callback: func(osPathname string, de *godirwalk.Dirent) error {
			if osPathname == osDirname {
				return nil
			}

			if de.IsDir() && osPathname != osDirname {
				log.Debug("entering directory:", osPathname)
				return nil
			}

			scrapItem := matcher.NewScrapItem(osPathname)

			rule, err := walker.Matcher.FindRuleFor(scrapItem)
			if err != nil {
				log.Warn(err)

			} else if rule != nil {
				scrapItem.Rule = rule
				pipe.Items <- scrapItem
				log.Debug(osPathname, "retained for scrapping by rule", rule.Name)

			} else {
				log.Info("ignore", osPathname)
			}

			return nil
		},
		ErrorCallback: func(osPathname string, err error) godirwalk.ErrorAction {
			log.Println(err)
			return godirwalk.SkipNode
		},
		PostChildrenCallback: func(osPathname string, de *godirwalk.Dirent) error {
			log.Debug("leaving directory:", osPathname)
			return nil
		},
	})
}
