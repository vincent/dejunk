package matcher

import (
	"io/ioutil"
	"os"
	"testing"

	log "github.com/sirupsen/logrus"
)

// Tests main
func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}
