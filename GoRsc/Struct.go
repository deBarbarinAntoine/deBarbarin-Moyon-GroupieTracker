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
		Id         string `json:"id"`
		Type       string `json:"type"`
		Attributes struct {
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
	}
}
