/*
Package app is the main runtime package for Efs2. This package holds all of the logging and task execution controls.
*/
package app

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/vincent/godejunk/pkg/config"
	"github.com/vincent/godejunk/pkg/matcher"
	"github.com/vincent/godejunk/pkg/walker"
	"github.com/vincent/godejunk/pkg/writer"
)

// Run is the primary execution function for this application.
func Run() error {

	cfg, err := config.NewConfig(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Config Output
	output := writer.NewStore()

	m := matcher.NewMatcher(cfg.RulesFile)

	// Process each input directory
	for _, dir := range cfg.Inputs {
		walk := walker.NewWalker(&m)

		log.Info("processing directory:", dir)
		pipe := matcher.NewScrapperPipe(&output)

		walk.WalkDirectory(dir, pipe)

		<-pipe.Done
	}

	// Show a result summary
	output.Tree.Fprint(os.Stdout, false, "")

	return nil
}
