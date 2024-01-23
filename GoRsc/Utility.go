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

func getMangaList(endpoint string) ([]Manga, error) {
	url := mangadexAPIURL + endpoint

	// Make a GET request to the MangaDex API
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("failed to make GET request")
		return nil, fmt.Errorf("failed to make GET request: %v", err)
	}
	defer resp.Body.Close()

	// Check HTTP status code
	if resp.StatusCode != http.StatusOK {
		fmt.Println("unexpected status code")
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("failed to read response body")
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}
	fmt.Println(string(body))
	// Decode the JSON response into a struct
	var result ListMangaResult
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("failed to decode JSON")
		return nil, fmt.Errorf("failed to decode JSON: %v", err)
	}

	return result, nil
}
