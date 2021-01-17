package rollback

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/icza/backscanner"
	log "github.com/sirupsen/logrus"

	"github.com/vincent/godejunk/pkg/matcher"
)

type RollbackFile struct {
	Enabled    bool
	outputPath string
	startedAt  time.Time
	file       *os.File
}

var logTimeFormat = "2006-01-02 15:04:05"
var splitMarker = ">>>>"
var startMarker = []byte("START:")
var endMarker = []byte("END:")

// NewRollbackFile returns a new RollbackFile struct
func NewRollbackFile(filename string, output string) *RollbackFile {
	var file *os.File = nil
	var err error = nil
	if filename != "" {
		file, err = os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		line := fmt.Sprintf("%s %s\n", string(startMarker), time.Now().Format(logTimeFormat))
		if _, err = file.WriteString(line); err != nil {
			log.Println("cannot write rollback start tag")
			file = nil
		}
	}

	return &RollbackFile{
		file:       file,
		startedAt:  time.Now(),
		outputPath: output,
		Enabled:    file != nil,
	}
}

// UndoLast play the last session backwards
func UndoLast(filename string, dry bool) {
	fmt.Println("Start rollback the last session from", filename)

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	index := readUntilLastStartMarker(file)

	scanner := bufio.NewScanner(file)
	_, err = file.Seek(index, io.SeekStart)
	if err != nil {
		panic(err)
	}

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 || (strings.Contains(line, string(endMarker)) || strings.Contains(line, string(startMarker))) {
			continue
		}

		parts := strings.Split(line, splitMarker)

		log.Println("move", parts[1], "back to", parts[0])
		fmt.Println("move", parts[1], "back to", parts[0])

		from, _ := strconv.Unquote(parts[1])
		to, _ := strconv.Unquote(parts[0])

		if !dry {
			err := os.MkdirAll(filepath.Dir(to), os.ModePerm)
			if err != nil {
				panic(err)
			}

			err = os.Rename(from, to)
			if err != nil {
				panic(err)
			}
		}
	}

	if !dry {
		os.Truncate(filename, index)
	}
}

// Close the underlying file
func (rf *RollbackFile) Close() {
	if rf.file != nil {
		line := fmt.Sprintf("%s %s: %f seconds\n", string(endMarker), time.Now().Format(logTimeFormat), time.Since(rf.startedAt).Seconds())
		if _, err := rf.file.WriteString(line); err != nil {
			log.Println("cannot write rollback end tag")
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
		splitMarker,
		strconv.Quote(filepath.Join(rf.outputPath, item.StorePath)))

	if _, err := rf.file.WriteString(line + "\n"); err != nil {
		log.Println("cannot write rollback line", line)
	} else {
		log.Println("wrote in rollback file", line)
	}
}

func readUntilLastStartMarker(file *os.File) int64 {
	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}
	index := stat.Size()

	scanner := backscanner.New(file, int(stat.Size()))

	for {
		line, pos, err := scanner.LineBytes()
		if err != nil && err != io.EOF {
			panic(err)
		}
		if err != nil && err == io.EOF {
			break
		}
		if bytes.Contains(line, startMarker) {
			index = int64(pos)
			break
		}
		// if len(line) > 0 && !bytes.Contains(line, endMarker) {
		// 	chline <- string(line)
		// }
	}
	return index
}
