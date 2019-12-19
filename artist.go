package lastfm

type ArtistService service

type ArtistInfo struct {
	Artist *Artist `json:"artist"`
}

type Artist struct {
	Name       string  `json:"name"`
	URL        string  `json:"url"`
	Image      []Image `json:"image"`
	Streamable string  `json:"streamable"`
	OnTour     string  `json:"ontour"`
	Stats      *struct {
		Listeners string `json:"listeners"`
		PlayCount string `json:"playcount"`
	} `json:"stats"`
	Similar *struct {
		Artist []Artist `json:"artist"`
	} `json:"similar"`
	Tags *struct {
		Tag []Tag `json:"tag"`
	} `json:"tags"`
	Bio *Bio `json:"bio"`
}

type Image struct {
	Text string `json:"#text"`
	Size string `json:"size"`
}

type Tag struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Bio struct {
	Links *struct {
		Link *Link `json:"link"`
	} `json:"links"`
	Published string `json:"published"`
	Summary   string `json:"summary"`
}

type Link struct {
	Text string `json:"#text"`
	Rel  string `json:"rel"`
	Href string `json:"href"`
}

type ArtistQuery struct {
	Method string `url:"method"`
	Artist string `url:"artist"`
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
