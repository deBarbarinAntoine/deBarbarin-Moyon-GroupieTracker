package server

import (
	"fmt"
	"log"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "404", nil)
}

const mangadexAPIURL = "https://api.mangadex.org"

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// afficher les recently updated (latestUploadedChapter) up to 10
	// afficher les recommanded (rating) up 15
	// afficher les plus populaire (followedCount) up to 10
	// image link https://uploads.mangadex.org/covers/:manga-id/:cover-filename

	// Fetch recently updated manga (up to 10)
	recentlyUpdated, err := getMangaList("/manga?order[latestUploadedChapter]=desc&limit=10&availableTranslatedLanguage[]=en")
	if err != nil {
		log.Printf("Error fetching recently updated manga: %v", err)
		return
	}
	updateCoverArtLinks(recentlyUpdated)
	for _, manga := range recentlyUpdated.Mangas {
		fmt.Println("rencent")
		fmt.Printf("Manga ID: %s, Cover Art Link: %s\n", manga.Id, manga.CoverArtLink)
		fmt.Println("--------------------------------------------------------------------------------------")
	}

	// Fetch recommended manga (up to 26)
	recommended, err := getMangaList("/manga?order[rating]=desc&limit=15")
	if err != nil {
		log.Println("Error fetching recommended manga:", err)
	}
	updateCoverArtLinks(recommended)
	for _, manga := range recommended.Mangas {
		fmt.Println("recomended")
		fmt.Printf("Manga ID: %s, Cover Art Link: %s\n", manga.Id, manga.CoverArtLink)
		fmt.Println("--------------------------------------------------------------------------------------")
	}
	// Fetch popular manga (up to 10)
	popular, err := getMangaList("/manga?order[followedCount]=desc&limit=10")
	if err != nil {
		log.Println("Error fetching popular manga:", err)
	}
	updateCoverArtLinks(popular)
	for _, manga := range popular.Mangas {
		fmt.Println("popular")
		fmt.Printf("Manga ID: %s, Cover Art Link: %s\n", manga.Id, manga.CoverArtLink)
		fmt.Println("--------------------------------------------------------------------------------------")
	}

	renderTemplate(w, "index", nil)
}

func ByTagHandler(w http.ResponseWriter, r *http.Request) {
	// affcher par derniere update
	// image link https://uploads.mangadex.org/covers/:manga-id/:cover-filename
	renderTemplate(w, "bytag", nil)
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	// affcher par pertinance
	// image link https://uploads.mangadex.org/covers/:manga-id/:cover-filename
	renderTemplate(w, "search", nil)
}

func SelectHandler(w http.ResponseWriter, r *http.Request) {
	// affcher toute les data et lien vers chapitre
	// image link https://uploads.mangadex.org/covers/:manga-id/:cover-filename
	renderTemplate(w, "select", nil)
}
