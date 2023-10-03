package service

import (
	"net/http"
	"sync"
)

type ShortenerMethods interface {
	ShortenURL(w http.ResponseWriter, r *http.Request)
	RedirectURL(w http.ResponseWriter, r *http.Request)
	Metrics(w http.ResponseWriter, r *http.Request)
}

// Shortener represents the URL shortener service.
type Shortener struct {
	urls      map[string]string // Map to store original URLs and their corresponding shortened URLs.
	domains   map[string]int    // Map to store counts of shortened URLs per domain.
	mutex     sync.Mutex        // Mutex for thread-safe access to maps.
	baseURL   string            // Base URL for shortened links.
	shortURLs map[string]bool   // Set to keep track of used shortened URLs.
}
