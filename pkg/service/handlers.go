package service

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strings"
)

// NewShortener initializes a new Shortener instance.
func NewShortener() *Shortener {
	return &Shortener{
		urls:      make(map[string]string),
		domains:   make(map[string]int),
		baseURL:   "http://infracloud.demo/",
		shortURLs: make(map[string]bool),
	}
}

// ShortenURL handles the shortening of url wh.
func (s *Shortener) ShortenURL(w http.ResponseWriter, r *http.Request) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	originalURL := r.URL.Query().Get("url")
	if originalURL == "" {
		http.Error(w, "URL parameter 'url' is required", http.StatusBadRequest)
		return
	}
	shortenedURL, exists := s.urls[originalURL]
	if !exists {
		// we are generating a unique shortened URL using SHA-1 hashing and base64 encoding.
		hash := sha1.Sum([]byte(originalURL))
		shortenedURL = base64.URLEncoding.EncodeToString(hash[:])[:8] // Use the first 8 characters.
		shortenedURL = fmt.Sprintf("%s%s", s.baseURL, shortenedURL)

		// Store the mapping of original URL to shortened URL.
		s.urls[originalURL] = shortenedURL

		// Extract domain for metrics.
		parts := strings.Split(originalURL, "/")
		if len(parts) > 2 {
			domain := parts[2]
			s.domains[domain]++
		}
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"shortened_url": shortenedURL}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}

// RedirectURL redirects the user to the original URL corresponding to the shortened URL.
func (u *Shortener) RedirectURL(w http.ResponseWriter, r *http.Request) {
	u.mutex.Lock()
	defer u.mutex.Unlock()

	shortenedURL := r.URL.Path[1:] // Remove the leading '/' from the path.
	originalURL, exists := u.urls[shortenedURL]
	if !exists {
		http.Error(w, "Shortened URL not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusSeeOther)
}

type domainCount struct {
	Domain string
	Count  int
}

// Metrics returns the top 3 domains with the most shortened URLs.
func (s *Shortener) Metrics(w http.ResponseWriter, r *http.Request) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Sort the domains by count.

	var domainCounts []domainCount
	for domain, count := range s.domains {
		domainCounts = append(domainCounts, domainCount{domain, count})
	}

	// Sort domainCounts by count in descending order.
	sortByCountDescending(domainCounts)

	// Prepare the top 3 domains.
	topDomains := make(map[string]int)
	for i, dc := range domainCounts {
		if i >= 3 {
			break
		}
		topDomains[dc.Domain] = dc.Count
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(topDomains)
	if err != nil {
		return
	}
}

// Helper function to sort domain counts by count in descending order.
func sortByCountDescending(domainCounts []domainCount) {
	sort.Slice(domainCounts, func(i, j int) bool {
		return domainCounts[i].Count > domainCounts[j].Count
	})
}
