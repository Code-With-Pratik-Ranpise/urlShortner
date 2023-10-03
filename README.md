# URL Shortener Service in Go

This is a simple URL shortener service written in Go that accepts a URL as an argument over a REST API and returns a shortened URL as a result.

## Features

- Shortens URLs using SHA-1 hashing and base64 encoding.
- Stores original URLs and their corresponding shortened URLs in memory.
- Supports thread-safe access to the URL mapping using a mutex.
- Provides a REST API endpoint for URL shortening.
- Reuses the same shortened URL for identical original URLs.
- Bonus: Redirects users to the original URL when accessing the shortened URL.
- Bonus: Metrics API for tracking the top 3 domains with the most shortened URLs.

## Usage

To use this URL shortener service, follow these steps:

1. Clone or download the project.

2. Run the Go application:

   ```bash
   go run main.go
