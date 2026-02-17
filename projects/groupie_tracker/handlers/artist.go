package handlers

import (
	"groupi_tracker/api"
	"groupi_tracker/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type ArtistData struct {
	Artist    models.Artist
	Locations models.Locations
	Dates     models.Dates
	Relations models.Relation
}

// ArtistHandler displays individual artist details
func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Get artist ID from URL (query parameter or path)
	idStr := r.URL.Query().Get("id")
	targetId, err := strconv.Atoi(idStr)
	if err != nil {
		RenderError(w, http.StatusBadRequest, "Invalid Artist ID")
		return
	}

	// 2. Fetch all artists
	artists, err := api.GetArtists()
	if err != nil {
		log.Printf("Error fetching artists: %v", err)
		RenderError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	// 3. Find the artist with matching ID
	var selectedArtist models.Artist
	found := false
	for _, artist := range artists {
		if targetId == artist.ID {
			selectedArtist = artist
			found = true
			break
		}
	}
	if !found {
		http.NotFound(w, r)
		return
	}
	// 4. Fetch additional data (locations, dates, relations)
	locations, err := api.GetLocations(selectedArtist.Locations)
	dates, err := api.GetDates(selectedArtist.ConcertDates)
	relations, err := api.GetRelations(selectedArtist.Relations)
	// 5. Combine data into a struct for template
	data := ArtistData{
		Artist:    selectedArtist,
		Locations: locations,
		Dates:     dates,
		Relations: relations,
	}
	// 6. Render artist.html template
	// Use a relative path that matches your project structure
	tmpl, err := template.ParseFiles("./templates/artist.html")
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		RenderError(w, http.StatusInternalServerError, "Template error")
		return
	}

	// 4. Execute template with artists data
	// Note: No need for &artists, passing the slice directly is standard
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Printf("Error executing template: %v", err)
	}
}
