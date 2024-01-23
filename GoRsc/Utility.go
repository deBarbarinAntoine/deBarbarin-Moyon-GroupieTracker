package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"text/template"
)

func renderTemplate(w http.ResponseWriter, tmplName string, data interface{}) {

	tmplPath := path + "templates/"

	tmpl := template.Must(template.ParseFiles(tmplPath+tmplName+".html", tmplPath+"base.html"))

	err := tmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		fmt.Println("Error executing template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func getMangaList(endpoint string) (Manga, error) {
	url := mangadexAPIURL + endpoint

	// Make a GET request to the MangaDex API
	resp, err := http.Get(url)
	if err != nil {
		return Manga{}, fmt.Errorf("failed to make GET request: %v", err)
	}
	defer resp.Body.Close()

	// Check HTTP status code
	if resp.StatusCode != http.StatusOK {
		return Manga{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Manga{}, fmt.Errorf("failed to read response body: %v", err)
	}

	// Decode the JSON response into a struct
	var data Manga
	if err := json.Unmarshal(body, &data); err != nil {
		return data, fmt.Errorf("failed to decode JSON: %v", err)
	}

	return data, nil
}

func getDataFromSeach(url string) Manga {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("failed to make GET request: %v", err)
		return Manga{}
	}
	defer resp.Body.Close()

	// Check HTTP status code
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("unexpected status code: %d", resp.StatusCode)
		return Manga{}
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("failed to read response body %v", err)
		return Manga{}
	}

	// Decode the JSON response into a struct
	var data Manga
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Printf("failed to decode JSON: %v", err)

		return Manga{}
	}

	return data
}

func getCoverFileNameByCover_ArtId(id string) string {
	url := mangadexAPIURL + "/cover/" + id
	// Make a GET request to the MangaDex API
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("failed to make GET request: %v", err)
		return ""
	}
	defer resp.Body.Close()

	// Check HTTP status code
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("unexpected status code: %d", resp.StatusCode)
		return ""
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("failed to read response body %v", err)
		return ""
	}

	// Decode the JSON response into a struct
	var data CoverArt
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Printf("failed to decode JSON: %v", err)

		return ""
	}
	data1 := data.Data.Attributes.FileName
	return data1
}

func updateCoverArtLinks(mangaData Manga) {
	for i, manga := range mangaData.Mangas {
		for _, relationship := range manga.Relationships {
			if relationship.Type == "cover_art" {
				// Use the getCoverFileNameByCover_ArtId function to get the cover art link
				coverArtLink := getCoverFileNameByCover_ArtId(relationship.Id)
				//https://uploads.mangadex.org/covers/:manga-id/:cover-filename
				mangaData.Mangas[i].CoverArtLink = "https://uploads.mangadex.org/covers/" + manga.Id + "/" + coverArtLink
			}
		}
	}
}

func getTagId(name string) string {
	url := "https://api.mangadex.org/manga/tag"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("failed to make GET request: %v", err)
		return ""
	}
	defer resp.Body.Close()

	// Check HTTP status code
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("unexpected status code: %d", resp.StatusCode)
		return ""
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("failed to read response body %v", err)
		return ""
	}

	// Decode the JSON response into a struct
	var data YourResponse
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Printf("failed to decode JSON: %v", err)

		return ""
	}
	for _, item := range data.Data {
		data1 := item.Attributes.Name["en"]
		if name == data1 {
			return item.ID
		}
	}
	fmt.Println("failed to get item ID")
	return ""
}

func buildMangaDEXQuery(title, authorOrArtist, year, order, orderEnum string, includedTags, excludedTags, status, availableTranslatedLanguage, publicationDemographic, ids []string) string {
	// Initialize the base URL for MangaDEX API
	var fullURL string
	baseURL := "https://api.mangadex.org/manga?limit=20&contentRating[]=safe"
	fullURL = baseURL
	if title != "" {
		titleUrl := "&title=" + title
		fullURL = fullURL + titleUrl
	} else {
		fmt.Println("no title")
	}

	if authorOrArtist != "" {
		authorOrArtistUrl := "&authorOrArtist=" + authorOrArtist
		fullURL = fullURL + authorOrArtistUrl
	} else {
		fmt.Println("no autorOrArtist")
	}

	if year != "" {
		yearUrl := "&year=" + year
		fullURL = fullURL + yearUrl
	} else {
		fmt.Println("no year")
	}

	if order != "" {
		if orderEnum != "" {
			orderUrl := "&order[" + order + "]=" + orderEnum
			fullURL = fullURL + orderUrl
		} else {
			fmt.Println("no orderEnum")
		}
	} else {
		fmt.Println("no order")
	}

	var includedTagsUrl string
	if len(includedTags) > 0 {
		for _, item := range includedTags {
			tag := getTagId(item)
			includedTagsUrl += "&includedTags[]=" + tag
		}
		fullURL = fullURL + includedTagsUrl
	} else {
		fmt.Println("no inclued tags")
	}

	var excludedTagsUrl string
	if len(excludedTags) > 0 {
		for _, item := range includedTags {
			tag := getTagId(item)
			excludedTagsUrl += "&excludedTags[]=" + tag
		}
		fullURL = fullURL + excludedTagsUrl
	} else {
		fmt.Println("no excluded tags")
	}
	var statusUrl string
	if len(status) > 0 {
		for _, item := range status {
			statusUrl += "&status[]=" + item
		}
		fullURL = fullURL + statusUrl
	} else {
		fmt.Println("no status")
	}

	var availableTranslatedLanguageUrl string
	if len(availableTranslatedLanguage) > 0 {
		for _, item := range availableTranslatedLanguage {
			availableTranslatedLanguageUrl += "&availableTranslatedLanguage[]=" + item
		}
		fullURL = fullURL + availableTranslatedLanguageUrl
	} else {
		fmt.Println("no available translated language")
	}
	var publicationDemographicUrl string
	if len(publicationDemographic) > 0 {
		for _, item := range publicationDemographic {
			publicationDemographicUrl += "&publicationDemographic[]=" + item
		}
		fullURL = fullURL + publicationDemographicUrl
	} else {
		fmt.Println("no publication demographic")
	}

	var idsUrl string
	if len(ids) > 0 && len(ids) < 100 {
		for _, item := range ids {
			idsUrl += "&ids[]=" + item
		}
		fullURL = fullURL + idsUrl
	} else {
		fmt.Println("no ids")
	}

	return fullURL
}
