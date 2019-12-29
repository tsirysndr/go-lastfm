package lastfm

type SearchResults struct {
	Results *struct {
		OpenSearchQuery *struct {
			Text        string `json:"#text,omitempty"`
			Role        string `json:"role,omitempty"`
			SearchTerms string `json:"searchTerms,omitempty"`
			StartPage   string `json:"startPage,omitempty"`
		} `json:"opensearch:Query,omitempty"`
		OpenSearchTotalResults string `json:"opensearch:totalResults,omitempty"`
		OpenSearchStartIndex   string `json:"opensearch:startIndex,omitempty"`
		OpenSearchItemsPerPage string `json:"opensearch:itemsPerPage,omitempty"`
		AlbumMatches           *struct {
			Album []Album `json:"album,omitempty"`
		} `json:"albummatches,omitempty"`
		ArtistMatches *struct {
			Artist []Artist `json:"artist,omitempty"`
		} `json:"artistmatches,omitempty"`
		Attr *struct {
			For string `json:"for,omitempty"`
		} `json:"@attr,omitempty"`
	} `json:"results,omitempty"`
}
