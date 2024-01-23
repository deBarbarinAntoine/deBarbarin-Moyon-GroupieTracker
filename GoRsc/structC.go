package server

// type Manga struct {
// 	Id         string `json:"id"`
// 	Type       string `json:"type"`
// 	Attributes struct {
// 		Title struct {
// 			En string `json:"en"`
// 		} `json:"title"`
// 		AltTitles []struct {
// 			Ja string `json:"ja,omitempty"`
// 			Zh string `json:"zh,omitempty"`
// 			Ko string `json:"ko,omitempty"`
// 		} `json:"altTitles"`
// 		Description struct {
// 			En string `json:"en"`
// 		} `json:"description"`
// 		IsLocked bool `json:"isLocked"`
// 		Links    struct {
// 			Raw   string `json:"raw,omitempty"`
// 			Engtl string `json:"engtl,omitempty"`
// 		} `json:"links"`
// 		OriginalLanguage       string  `json:"originalLanguage"`
// 		LastVolume             string  `json:"lastVolume"`
// 		LastChapter            string  `json:"lastChapter"`
// 		PublicationDemographic *string `json:"publicationDemographic"`
// 		Status                 string  `json:"status"`
// 		Year                   int     `json:"year"`
// 		ContentRating          string  `json:"contentRating"`
// 		Tags                   []struct {
// 			Id         string `json:"id"`
// 			Type       string `json:"type"`
// 			Attributes struct {
// 				Name struct {
// 					En string `json:"en"`
// 				} `json:"name"`
// 				Description struct {
// 				} `json:"description"`
// 			} `json:"attributes"`
// 			Relationships []interface{} `json:"relationships"`
// 		} `json:"tags"`
// 		State                          string    `json:"state"`
// 		ChapterNumbersResetOnNewVolume bool      `json:"chapterNumbersResetOnNewVolume"`
// 		CreatedAt                      time.Time `json:"createdAt"`
// 		UpdatedAt                      time.Time `json:"updatedAt"`
// 		Version                        int       `json:"version"`
// 		AvailableTranslatedLanguages   []string  `json:"availableTranslatedLanguages"`
// 		LatestUploadedChapter          string    `json:"latestUploadedChapter"`
// 	} `json:"attributes"`
// 	Relationships []struct {
// 		Id      string `json:"id"`
// 		Type    string `json:"type"`
// 		Related string `json:"related,omitempty"`
// 	} `json:"relationships"`
// }
