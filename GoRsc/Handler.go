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
const mangadexCoverURL = "https://uploads.mangadex.org/covers"

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// afficher les recently updated up to 10
	// afficher les recommanded up 26
	// afficher les plus populaire up to 10
	// image link https://uploads.mangadex.org/covers/:manga-id/:cover-filename

	// Fetch recently updated manga (up to 10)
	recentlyUpdated, err := getMangaList("/manga?order[updatedAt]=desc&limit=10")
	if err != nil {
		log.Printf("Error fetching recently updated manga: %v", err)
		return
	}

	if len(recentlyUpdated.Mangas) == 0 {
		fmt.Println("No recently updated manga.")
	} else {
		fmt.Println(recentlyUpdated)
	}
	// // Fetch recommended manga (up to 26)
	// recommended, err := getMangaList("/manga?limit=26")
	// if err != nil {
	// 	log.Println("Error fetching recommended manga:", err)
	// }
	// fmt.Println(recommended)
	// // Fetch popular manga (up to 10)
	// popular, err := getMangaList("/manga?order[views]=desc&limit=10")
	// if err != nil {
	// 	log.Println("Error fetching popular manga:", err)
	// }
	// fmt.Println(popular)

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
