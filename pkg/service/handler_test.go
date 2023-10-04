package service

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestShortenURL(t *testing.T) {
	shortener := NewShortener()

	//HTTP request.
	req, err := http.NewRequest("GET", "/shorten?url=http://infracloudtechaws.com", nil)
	if err != nil {
		t.Fatal(err)
	}

	//ResponseRecorder to capture the response.
	rr := httptest.NewRecorder()

	// Calling the ShortenURL method.
	shortener.ShortenURL(rr, req)

	// Check the HTTP status code.
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, rr.Code)
	}
}

func TestRedirectURL(t *testing.T) {
	shortener := NewShortener()
	shortURL := "infracloud"
	shortener.urls[shortURL] = "http://infracloudtechaws.com"

	//HTTP request with the shortened URL.
	req, err := http.NewRequest("GET", "/"+shortURL, nil)
	if err != nil {
		t.Fatal(err)
	}

	//ResponseRecorder to capture the response.
	rr := httptest.NewRecorder()

	// Calling the RedirectURL method.
	shortener.RedirectURL(rr, req)

	//HTTP status code.
	if rr.Code != http.StatusSeeOther {
		t.Errorf("Expected status code %d, but got %d", http.StatusSeeOther, rr.Code)
	}
}

func TestMetrics(t *testing.T) {
	shortener := NewShortener()
	shortener.domains["infracloudtechaws.com"] = 5
	shortener.domains["google.com"] = 3

	//HTTP request.
	req, err := http.NewRequest("GET", "/metrics", nil)
	if err != nil {
		t.Fatal(err)
	}

	// ResponseRecorder to capture the response.
	rr := httptest.NewRecorder()

	// Calling the Metrics method.
	shortener.Metrics(rr, req)

	//HTTP status code.
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, rr.Code)
	}
}

func TestSortByCountDescending(t *testing.T) {
	domainCounts := []domainCount{
		{"infracloudtechaws.com", 5},
		{"google.com", 3},
		{"yahoo.com", 7},
	}

	// Calling the sorting function.
	sortByCountDescending(domainCounts)

	expectedOrder := []string{"yahoo.com", "infracloudtechaws.com", "google.com"}
	for i, dc := range domainCounts {
		if dc.Domain != expectedOrder[i] {
			t.Errorf("Expected domain '%s' at position %d, but got '%s'", expectedOrder[i], i, dc.Domain)
		}
	}
}
