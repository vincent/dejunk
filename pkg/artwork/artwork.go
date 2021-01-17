package artwork

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"

	log "github.com/sirupsen/logrus"

	"github.com/vincent/godejunk/pkg/matcher"
)

type ArtworkFinder interface {
	FindCover(itemType matcher.ItemType, terms []string) string
	FindBackground(itemType matcher.ItemType, terms []string) string
	CleanUp()
}

type ArtworkDownloader interface {
	FindCoverUrl(itemType matcher.ItemType, terms []string) string
	FindBackgroundUrl(itemType matcher.ItemType, terms []string) string
}

type ArtworkCache struct {
	mu                sync.Mutex
	urls              map[string]string
	files             map[string]string
	coverBackend      ArtworkDownloader
	backgroundBackend ArtworkDownloader
	Dry               bool
}

func NewArtworkFinder() *ArtworkCache {
	return &ArtworkCache{
		urls:              map[string]string{},
		files:             map[string]string{},
		coverBackend:      &ItunesSearch{},
		backgroundBackend: &ItunesSearch{},
		Dry:               false,
	}
}

func (a *ArtworkCache) FindCover(itemType matcher.ItemType, terms []string) string {
	log.Info("try to find cover for ", itemType, ":", terms)

	a.mu.Lock()
	defer a.mu.Unlock()
	key := strings.Join(append(terms, string(itemType), "cover"), "-")
	if res, ok := a.files[key]; ok {
		return res
	}
	if res, ok := a.urls[key]; ok {
		return res
	}
	a.urls[key] = a.coverBackend.FindCoverUrl(itemType, terms)

	if a.urls[key] != "" {
		tmpfileName, err := a.downloadFile(a.urls[key])
		if err == nil {
			a.files[key] = tmpfileName
		}
	}

	return a.files[key]
}

func (a *ArtworkCache) FindBackground(itemType matcher.ItemType, terms []string) string {
	log.Info("try to find background for ", itemType, ":", terms)

	a.mu.Lock()
	defer a.mu.Unlock()
	key := strings.Join(append(terms, string(itemType), "background"), "-")
	if res, ok := a.files[key]; ok {
		return res
	}
	if res, ok := a.urls[key]; ok {
		return res
	}
	a.urls[key] = a.backgroundBackend.FindBackgroundUrl(itemType, terms)

	if a.urls[key] != "" {
		tmpfileName, err := a.downloadFile(a.urls[key])
		if err == nil {
			a.files[key] = tmpfileName
		}
	}

	return a.files[key]
}

func (a *ArtworkCache) downloadFile(url string) (string, error) {
	log.Info("download ", url)

	if a.Dry {
		return "fake-file", nil
	}

	// Create a tempfile
	tempfile, err := ioutil.TempFile("", "artwork*"+filepath.Ext(url))
	if err != nil {
		log.Fatal(err)
	}

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(tempfile, resp.Body)
	if err != nil {
		return "", err
	}

	return tempfile.Name(), nil
}

func (a *ArtworkCache) CleanUp() {
	for _, file := range a.files {
		log.Info("remove ", file)
		os.Remove(file)
	}
}
