package handlers

import (
	"groupi_tracker/api"
	"groupi_tracker/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// SearchHandler filters artists by name, member, creation year, or first album
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	query := strings.TrimSpace(r.URL.Query().Get("q"))
	yearStr := strings.TrimSpace(r.URL.Query().Get("year"))

	// Fetch all artists
	artists, err := api.GetArtists()
	if err != nil {
		log.Printf("Error fetching artists in search: %v", err)
		RenderError(w, http.StatusInternalServerError, "Failed to fetch artists data")
		return
	}

	filtered := filterArtists(artists, query, yearStr)

	// Reuse index.html template — it already ranges over a slice of artists
	tmpl, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		RenderError(w, http.StatusInternalServerError, "Template error")
		return
	}

	// Pass search data to template
	data := struct {
		Artists []models.Artist
		Query   string
		Year    string
	}{
		Artists: filtered,
		Query:   query,
		Year:    yearStr,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Printf("Error executing search template: %v", err)
	}
}

// filterArtists returns artists matching the query string and/or creation year
func filterArtists(artists []models.Artist, query, yearStr string) []models.Artist {
	filtered := []models.Artist{}

	yearFilter, yearErr := strconv.Atoi(yearStr)

	for _, artist := range artists {
		// Year filter — skip if year provided and doesn't match
		if yearErr == nil && yearStr != "" {
			if artist.CreationDate != yearFilter {
				continue
			}
		}

		// If no text query, include (already passed year filter above)
		if query == "" {
			filtered = append(filtered, artist)
			continue
		}

		lq := strings.ToLower(query)

		// Match on name
		if strings.Contains(strings.ToLower(artist.Name), lq) {
			filtered = append(filtered, artist)
			continue
		}

		// Match on any member name
		matched := false
		for _, member := range artist.Members {
			if strings.Contains(strings.ToLower(member), lq) {
				matched = true
				break
			}
		}
		if matched {
			filtered = append(filtered, artist)
			continue
		}

		// Match on first album
		if strings.Contains(strings.ToLower(artist.FirstAlbum), lq) {
			filtered = append(filtered, artist)
		}
	}

	return filtered
}
