package server

import "time"

type BaseData struct {
	Title      string
	StaticPath string
}

type ListMangaResult struct {
	Result   string `json:"result"`
	Response string `json:"response"`
	Mangas   []struct {
		Id         string `json:"id"`
		Type       string `json:"type"`
		Attributes struct {
			Title struct {
				En string `json:"en"`
			} `json:"title"`
			AltTitles []struct {
				Fr   string `json:"fr,omitempty"`
				Ja   string `json:"ja,omitempty"`
				Zh   string `json:"zh,omitempty"`
				Ko   string `json:"ko,omitempty"`
				JaRo string `json:"ja-ro,omitempty"`
				En   string `json:"en,omitempty"`
				Vi   string `json:"vi,omitempty"`
				PtBr string `json:"pt-br,omitempty"`
				ZhHk string `json:"zh-hk,omitempty"`
				ZhRo string `json:"zh-ro,omitempty"`
				Pl   string `json:"pl,omitempty"`
				Tr   string `json:"tr,omitempty"`
				EsLa string `json:"es-la,omitempty"`
				Ru   string `json:"ru,omitempty"`
			} `json:"altTitles"`
			Description struct {
				En   string `json:"en"`
				EsLa string `json:"es-la,omitempty"`
				PtBr string `json:"pt-br,omitempty"`
				Ja   string `json:"ja,omitempty"`
			} `json:"description"`
			IsLocked bool `json:"isLocked"`
			Links    struct {
				Al    string `json:"al,omitempty"`
				Ap    string `json:"ap,omitempty"`
				Bw    string `json:"bw,omitempty"`
				Kt    string `json:"kt,omitempty"`
				Mu    string `json:"mu,omitempty"`
				Amz   string `json:"amz,omitempty"`
				Cdj   string `json:"cdj,omitempty"`
				Ebj   string `json:"ebj,omitempty"`
				Mal   string `json:"mal,omitempty"`
				Raw   string `json:"raw,omitempty"`
				Engtl string `json:"engtl,omitempty"`
				Nu    string `json:"nu,omitempty"`
			} `json:"links"`
			OriginalLanguage       string  `json:"originalLanguage"`
			LastVolume             string  `json:"lastVolume"`
			LastChapter            string  `json:"lastChapter"`
			PublicationDemographic *string `json:"publicationDemographic"`
			Status                 string  `json:"status"`
			Year                   int     `json:"year"`
			ContentRating          string  `json:"contentRating"`
			Tags                   []struct {
				Id         string `json:"id"`
				Type       string `json:"type"`
				Attributes struct {
					Name struct {
						En string `json:"en"`
					} `json:"name"`
					Description struct {
					} `json:"description"`
					Group   string `json:"group"`
					Version int    `json:"version"`
				} `json:"attributes"`
				Relationships []interface{} `json:"relationships"`
			} `json:"tags"`
			State                          string    `json:"state"`
			ChapterNumbersResetOnNewVolume bool      `json:"chapterNumbersResetOnNewVolume"`
			CreatedAt                      time.Time `json:"createdAt"`
			UpdatedAt                      time.Time `json:"updatedAt"`
			Version                        int       `json:"version"`
			AvailableTranslatedLanguages   []string  `json:"availableTranslatedLanguages"`
			LatestUploadedChapter          string    `json:"latestUploadedChapter"`
		} `json:"attributes"`
		Relationships []struct {
			Id      string `json:"id"`
			Type    string `json:"type"`
			Related string `json:"related,omitempty"`
		} `json:"relationships"`
	} `json:"data"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}

type Manga struct {
	Mangas []struct {
		Id           string `json:"id"`
		Type         string `json:"type"`
		CoverArtLink string `json:"coverArtLink,omitempty"`
		Attributes   struct {
			Title struct {
				En string `json:"en"`
			} `json:"title"`
			AltTitles []struct {
				Ja string `json:"ja,omitempty"`
				Zh string `json:"zh,omitempty"`
				Ko string `json:"ko,omitempty"`
			} `json:"altTitles"`
			Description struct {
				En string `json:"en"`
			} `json:"description"`
			IsLocked bool `json:"isLocked"`
			Links    struct {
				Raw   string `json:"raw,omitempty"`
				Engtl string `json:"engtl,omitempty"`
			} `json:"links"`
			OriginalLanguage       string  `json:"originalLanguage"`
			LastVolume             string  `json:"lastVolume"`
			LastChapter            string  `json:"lastChapter"`
			PublicationDemographic *string `json:"publicationDemographic"`
			Status                 string  `json:"status"`
			Year                   int     `json:"year"`
			ContentRating          string  `json:"contentRating"`
			Tags                   []struct {
				Id         string `json:"id"`
				Type       string `json:"type"`
				Attributes struct {
					Name struct {
						En string `json:"en"`
					} `json:"name"`
					Description struct {
					} `json:"description"`
				} `json:"attributes"`
			} `json:"tags"`
			State                          string    `json:"state"`
			ChapterNumbersResetOnNewVolume bool      `json:"chapterNumbersResetOnNewVolume"`
			CreatedAt                      time.Time `json:"createdAt"`
			UpdatedAt                      time.Time `json:"updatedAt"`
			AvailableTranslatedLanguages   []string  `json:"availableTranslatedLanguages"`
			LatestUploadedChapter          string    `json:"latestUploadedChapter"`
		} `json:"attributes"`
		Relationships []struct {
			Id      string `json:"id"`
			Type    string `json:"type"`
			Related string `json:"related,omitempty"`
		} `json:"relationships"`
	} `json:"data"`
}

// CoverArt represents the structure of the JSON data
type CoverArt struct {
	Result   string       `json:"result"`
	Response string       `json:"response"`
	Data     CoverArtData `json:"data"`
}

// CoverArtData represents the data field in the JSON
type CoverArtData struct {
	ID            string                 `json:"id"`
	Type          string                 `json:"type"`
	Attributes    CoverArtAttributes     `json:"attributes"`
	Relationships []CoverArtRelationship `json:"relationships"`
}

// CoverArtAttributes represents the attributes field in the JSON
type CoverArtAttributes struct {
	Description string    `json:"description"`
	Volume      string    `json:"volume"`
	FileName    string    `json:"fileName"`
	Locale      string    `json:"locale"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Version     int       `json:"version"`
}

// CoverArtRelationship represents the relationship field in the JSON
type CoverArtRelationship struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

// Attributes struct to represent the attributes field in the JSON
type Attributes struct {
	Name        map[string]string `json:"name"`
	Description map[string]string `json:"description"`
	Group       string            `json:"group"`
	Version     int               `json:"version"`
}

// DataItem struct to represent each item in the "data" array
type DataItem struct {
	ID            string     `json:"id"`
	Type          string     `json:"type"`
	Attributes    Attributes `json:"attributes"`
	Relationships []struct{} `json:"relationships"`
}

// YourResponse struct to represent the entire JSON response
type YourResponse struct {
	Result   string     `json:"result"`
	Response string     `json:"response"`
	Data     []DataItem `json:"data"`
}

// structs for chapters handling
type Attributes1 struct {
	Volume             string    `json:"volume"`
	Chapter            string    `json:"chapter"`
	Title              string    `json:"title"`
	TranslatedLanguage string    `json:"translatedLanguage"`
	ExternalURL        *string   `json:"externalUrl"`
	PublishAt          time.Time `json:"publishAt"`
	ReadableAt         time.Time `json:"readableAt"`
	CreatedAt          time.Time `json:"createdAt"`
	UpdatedAt          time.Time `json:"updatedAt"`
	Pages              int       `json:"pages"`
	Version            int       `json:"version"`
}

// structs for chapters handling
type Relationship struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

// structs for chapters handling
type Data struct {
	ID            string         `json:"id"`
	Type          string         `json:"type"`
	Attributes    Attributes1    `json:"attributes"`
	Relationships []Relationship `json:"relationships"`
}

// structs for chapters handling
type Response struct {
	Result   string `json:"result"`
	Response string `json:"response"`
	Data     []Data `json:"data"`
}

// Define a structure for the chapter data
type MangaChapter struct {
	Hash      string   `json:"hash"`
	Data      []string `json:"data"`
	DataSaver []string `json:"dataSaver"`
}

// Define a structure for the overall response
type MangaChapterResponse struct {
	Result  string       `json:"result"`
	BaseURL string       `json:"baseUrl"`
	Chapter MangaChapter `json:"chapter"`
}
