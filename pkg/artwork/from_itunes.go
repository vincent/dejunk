package artwork

import (
	itunes "github.com/Vonng/go-itunes-search"
	"github.com/vincent/godejunk/pkg/matcher"
)

var entities = map[matcher.ItemType]string{
	matcher.Movie:  itunes.Movie,
	matcher.TVShow: itunes.TvShow,
	matcher.Music:  itunes.Music,
	// matcher.Music: itunes.Podcast,
	// matcher.Music: itunes.MusicVideo,
	// matcher.Music: itunes.AudioBook,
	// matcher.Music: itunes.ShortFilm,
	// matcher.Music: itunes.Software,
	// matcher.Music: itunes.Ebook,
	// matcher.Music: itunes.All,
}

type ItunesSearch struct{}

func (it *ItunesSearch) FindBackgroundUrl(itemType matcher.ItemType, terms []string) string {
	res, _ := itunes.Search(terms).Media(entities[itemType]).Country(itunes.US).Limit(1).Results()

	for _, size := range res[0].ScreenshotURLs {
		if size != "" {
			return size
		}
	}

	return ""
}

func (it *ItunesSearch) FindCoverUrl(itemType matcher.ItemType, terms []string) string {
	res, _ := itunes.Search(terms).Media(entities[itemType]).Country(itunes.US).Limit(1).Results()

	for _, size := range []string{res[0].ArtworkURL512, res[0].ArtworkURL100, res[0].ArtworkURL60} {
		if size != "" {
			return size
		}
	}
	return ""
}
