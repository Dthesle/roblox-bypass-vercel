// reverseproxy.go
package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

func main() {
	// Define the target URL to forward requests to
	targetURL, err := url.Parse("https://websitesball.com/")
	if err != nil {
		panic(err)
	}

	// Create a reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	// Get the port from the environment variable or use the default (3000)
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Start the server on the specified port
	fmt.Printf("Reverse Proxy Server is running on :%s\n", port)
	http.ListenAndServe(":"+port, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Serve the request using the reverse proxy
		proxy.ServeHTTP(w, r)
	}))
}
