package lastfm

type SearchResults struct {
	Results *struct {
		OpenSearchQuery *struct {
			Text        string `json:"#text"`
			Role        string `json:"role"`
			SearchTerms string `json:"searchTerms"`
			StartPage   string `json:"startPage"`
		} `json:"opensearch:Query"`
		OpenSearchTotalResults string `json:"opensearch:totalResults"`
		OpenSearchStartIndex   string `json:"opensearch:startIndex"`
		OpenSearchItemsPerPage string `json:"opensearch:itemsPerPage"`
		AlbumMatches           *struct {
			Album []Album `json:"album"`
		} `json:"albummatches"`
		ArtistMatches *struct {
			Artist []Artist `json:"artist"`
		} `json:"artistmatches"`
		Attr *struct {
			For string `json:"for"`
		} `json:"@attr"`
	} `json:"results"`
}
