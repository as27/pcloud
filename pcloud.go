package pcloud

import (
	"net/http"
	"time"
)

// The different host URLs of the pCloud API. The URL can be setz with the HostURL variable.
// If there will be a new host URL, it can be set with the HostURL variable.
// The HostURL variable is set to the HostURLEU variable by default.
const (
	HostURLEU = "https://eapi.pcloud.com"
	HostURLUS = "https://api.pcloud.com"
)

// The different authentication methods, which can be used to authenticate with the pCloud API.
// The AuthMethodPassword method sends the password as parameters. So it is not recomendet to
// use this method, because the password is sent in plain text.
// The AuthMethodDigest method sends the password as a hashed string. To do this a digest value
// is requested from the pCloud API.
const (
	AuthMethodPassword = iota // Send password as get parameter
	AuthMethodDigest          // Send a hashed password as get parameter
)

// digestMethod ist the path to the digest method of the pCloud API.
const digestMethod = "/getdigest"

// The global configuration variables of the pCloud API.
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
