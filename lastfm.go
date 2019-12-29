package lastfm

import (
	"github.com/dghubble/sling"
)

type Client struct {
	base    *sling.Sling
	common  service
	Album   *AlbumService
	Artist  *ArtistService
	Library *LibraryService
	Tag     *TagService
	Track   *TrackService
	User    *UserService
}

type service struct {
	client *Client
}

type Params struct {
	ApiKey string `url:"api_key,omitempty"`
	Format string `url:"format,omitempty"`
}

func NewClient(apikey string) *Client {
	params := &Params{apikey, "json"}
	base := sling.New().Base("http://ws.audioscrobbler.com/2.0/").QueryStruct(params)
	c := &Client{}
	c.base = base
	c.common.client = c
	c.Album = (*AlbumService)(&c.common)
	c.Artist = (*ArtistService)(&c.common)
	c.Library = (*LibraryService)(&c.common)
	c.Tag = (*TagService)(&c.common)
	c.Track = (*TrackService)(&c.common)
	c.User = (*UserService)(&c.common)
	return c
}
