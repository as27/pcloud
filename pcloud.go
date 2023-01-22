package pcloud

import (
	"net/http"
	"time"
)

const (
	HostURLEU    = "https://eapi.pcloud.com"
	HostURLUS    = "https://api.pcloud.com"
	DigestMethod = "/getdigest"
)

const (
	AuthMethodPassword = iota // Send password as get parameter
	AuthMethodDigest          // Send a hashed password as get parameter
)

var (
	HostURL    = HostURLEU        // HostURL is the base URL for the pCloud API.
	AuthMethod = AuthMethodDigest // AuthMethod is the method used to authenticate with the pCloud API.
	Timeout    = 5 * time.Second  // Timeout for the API requests
)

// HTTPClient is the HTTP client used to make requests to the pCloud API.
// If you want to define your own HTTP client, you can do so by setting
// this variable to your client.
// The default client uses the Timout variable of this package.
var HTTPClient = &http.Client{
	Timeout: Timeout,
}

// authToken is the authentication token used to authenticate with the pCloud API.
// This token is set, when the user is authenticated.
var authToken string

// AuthToken returns the authentication token used to authenticate with the pCloud API.
func AuthToken() string {
	return authToken
}
