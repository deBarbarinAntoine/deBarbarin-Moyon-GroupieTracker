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
