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
	// Create instances of Manga struct for each list
	recentlyUpdatedManga := Manga{Mangas: recentlyUpdated.Mangas}
	recommendedManga := Manga{Mangas: recommended.Mangas}
	popularManga := Manga{Mangas: popular.Mangas}
	data := map[string]interface{}{
		"RecentlyUpdated": recentlyUpdatedManga,
		"Recommended":     recommendedManga,
		"Popular":         popularManga,
	}
	renderTemplate(w, "index", data)
}

func ByTagHandler(w http.ResponseWriter, r *http.Request) {
	// affcher par derniere update avec le / les tag donner en query
	// image link https://uploads.mangadex.org/covers/:manga-id/:cover-filename
	renderTemplate(w, "bytag", nil)
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	// affcher par pertinance via query
	// image link https://uploads.mangadex.org/covers/:manga-id/:cover-filename
	renderTemplate(w, "search", nil)
}

func SelectHandler(w http.ResponseWriter, r *http.Request) {
	// affcher toute les data et lien vers chapitre via query
	// image link https://uploads.mangadex.org/covers/:manga-id/:cover-filename
	renderTemplate(w, "select", nil)
}
