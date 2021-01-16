package rollback

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/vincent/godejunk/pkg/matcher"
)

type RollbackFile struct {
	outputPath string
	startedAt  time.Time
	file       *os.File
}

var logTimeFormat = "2006-01-02 15:04:05"

// NewRollbackFile returns a new RollbackFile struct
func NewRollbackFile(filename string, output string) *RollbackFile {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}

	line := fmt.Sprintf("START: %s\n", time.Now().Format(logTimeFormat))
	if _, err := file.WriteString(line); err != nil {
		log.Println("cannot write rollback start tag")
		file = nil
	}

	return &RollbackFile{
		file:       file,
		startedAt:  time.Now(),
		outputPath: output,
	}
}

// Close the underlying file
func (rf *RollbackFile) Close() {
	if rf.file != nil {
		line := fmt.Sprintf("END: %s: %f seconds\n", time.Now().Format(logTimeFormat), time.Since(rf.startedAt).Seconds())
		if _, err := rf.file.WriteString(line); err != nil {
			log.Println("cannot write rollback start tag")
		}
		rf.file.Close()
	}
}

func (rf *RollbackFile) Write(item *matcher.ScrapItem) {
	if rf.file == nil {
		return
	}

	line := fmt.Sprint(
		strconv.Quote(item.SourcePath),
		">>>>",
		strconv.Quote(filepath.Join(rf.outputPath, item.StorePath)))

	if _, err := rf.file.WriteString(line + "\n"); err != nil {
		log.Println("cannot write rollback line", line)
	} else {
		log.Println("wrote in rollback file", line)
	}
}
