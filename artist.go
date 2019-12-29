package lastfm

type ArtistService service

type ArtistInfo struct {
	Artist *Artist `json:"artist,omitempty"`
}

type Artist struct {
	Name       string  `json:"name,omitempty"`
	URL        string  `json:"url,omitempty"`
	Image      []Image `json:"image,omitempty"`
	Streamable string  `json:"streamable,omitempty"`
	OnTour     string  `json:"ontour,omitempty"`
	Stats      *struct {
		Listeners string `json:"listeners,omitempty"`
		PlayCount string `json:"playcount,omitempty"`
	} `json:"stats,omitempty"`
	Similar *struct {
		Artist []Artist `json:"artist,omitempty"`
	} `json:"similar,omitempty"`
	Tags *struct {
		Tag []Tag `json:"tag,omitempty"`
	} `json:"tags,omitempty"`
	Bio *Bio `json:"bio,omitempty"`
}

type Image struct {
	Text string `json:"#text,omitempty"`
	Size string `json:"size,omitempty"`
}

type Tag struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}

type Bio struct {
	Links *struct {
		Link *Link `json:"link,omitempty"`
	} `json:"links,omitempty"`
	Published string `json:"published,omitempty"`
	Summary   string `json:"summary,omitempty"`
}

type Link struct {
	Text string `json:"#text,omitempty"`
	Rel  string `json:"rel,omitempty"`
	Href string `json:"href,omitempty"`
}

type ArtistQuery struct {
	Method string `url:"method,omitempty"`
	Artist string `url:"artist,omitempty"`
}

func (s *ArtistService) Search(artist string) (*SearchResults, error) {
	var err error
	params := &ArtistQuery{
		"artist.search",
		artist,
	}
	result := new(SearchResults)
	s.client.base.Get("").QueryStruct(params).Receive(result, err)
	return result, err
}

func (s *ArtistService) GetInfo(artist string) (*ArtistInfo, error) {
	var err error
	params := &ArtistQuery{
		"artist.getinfo",
		artist,
	}
	info := new(ArtistInfo)
	s.client.base.Get("").QueryStruct(params).Receive(info, err)
	return info, err
}
