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

Endpoints
The service exposes the following endpoints:

Shorten URL: To shorten a long URL, send a POST request to /shorten with the url parameter in the request body.

Example:

bash
Copy code
curl -X POST http://localhost:8080/shorten?url=http://example.com
Redirect URL: To redirect to the original URL associated with a shortened URL, make a GET request to /r/{shortened-url}.

Example:

bash
Copy code
curl http://localhost:8080/r/{shortened-url}
Metrics: To retrieve metrics on the top 3 domains with the most shortened URLs, make a GET request to /metrics.

Example:

bash
Copy code
curl http://localhost:8080/metrics
Unit Tests
Unit tests are available for the codebase to ensure its correctness. To run the tests, execute the following command:

bash
Copy code
go test ./...

## Usage

To use this URL shortener service, follow these steps:

1. Clone or download the project.

2. Run the Go application:

   ```bash
   go run main.go
