/*
Package config is the internal configuration used for Efs2. This configuration is for the internal application execution. It exists to pave the way for non-cli instances of Efs2 in the future.
*/
package config

import (
	"os"
	"strings"

	"github.com/jessevdk/go-flags"
	log "github.com/sirupsen/logrus"
	"github.com/vincent/godejunk/pkg/matcher"
)

// Config are command-line options that are provided by the user.
type Config struct {
	LogLevel     string   `short:"l" long:"level"      description:"Log level (fatal, error, info, debug)" default:"error"`
	DryRun       bool     `short:"d" long:"dryrun"     description:"Print tasks to be executed without actually executing any tasks"`
	Inputs       []string `short:"i" long:"in"         description:"Directories to scan"`
	Output       string   `short:"o" long:"out"        description:"Directoriy to write files"`
	RulesFile    string   `short:"r" long:"rules"      description:"Rules file" default:"./samples/rules.yml"`
	RollbackFile string   `short:"b" long:"rollback"   description:"Rollback file"`
	DoRollback   bool     `short:"x" long:"dorollback" description:"Do rollback last session"`

	// Matching rules
	Rules []matcher.Rule
}

// NewConfig returns an initialized config object
func NewConfig(args []string) (*Config, error) {
	var cfg Config

	// Parse command line arguments
	_, err := flags.ParseArgs(&cfg, os.Args[1:])
	if err != nil {
		return nil, err
	}

	// Trim file paths
	cfg.Output = strings.TrimSpace(cfg.Output)
	cfg.RollbackFile = strings.TrimSpace(cfg.RollbackFile)

	// Set log level
	switch cfg.LogLevel {
	case "none":
		log.SetLevel(log.FatalLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "debug":
		log.SetLevel(log.DebugLevel)
	default:
		log.SetLevel(log.ErrorLevel)
	}

	return &cfg, nil
}
