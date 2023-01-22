package pcloud

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// LogIn authenticates the user with the pCloud API.
// With the given AuthMethod the auth-token is requested.
func LogIn(user, pass string) error {
	token, err := getAuthToken(user, pass)
	if err != nil {
		return fmt.Errorf("LogIn getAuthToken: %w", err)
	}
	authToken = token
	return nil
}

// LogOut logs the user out of the pCloud API. When the
// logout request is successful the authToken is set to "".
func LogOut() error {
	type logoutResponse struct {
		Result      int  `json:"result"`
		AuthDeleted bool `json:"auth_deleted"`
	}
	lr := &logoutResponse{}
	err := ApiRequest(lr, "logout")
	if err != nil {
		return fmt.Errorf("LogOut apiRequest: %w", err)
	}

	if !lr.AuthDeleted {
		return fmt.Errorf("LogOut at server not succsessfull")
	}
	authToken = ""
	return nil
}

func getAuthToken(user, pass string) (string, error) {
	type authResponse struct {
		Result int    `json:"result"`
		Token  string `json:"auth"`
	}
	params := []Param{
		{Name: "getauth", Val: "1"},
		{Name: "logout", Val: "1"},
		{Name: "username", Val: user},
	}

	// just the url changes, when AuthMethod changes
	switch AuthMethod {
	case AuthMethodPassword:
		params = append(params, Param{"password", pass})
	case AuthMethodDigest:
		digest, err := getDigest()
		if err != nil {
			return "", fmt.Errorf("getAuthToken getDigest: %w", err)
		}
		passwordDigest := makePasswordDigest(user, pass, digest)
		params = append(params, Param{"passworddigest", passwordDigest})
		params = append(params, Param{"digest", digest})
	}
	ar := &authResponse{}
	err := ApiRequest(ar, "userinfo", params...)

	if err != nil {
		return "", fmt.Errorf("getAuthToken apiRequest: %w", err)
	}
	return ar.Token, nil
}

func getDigest() (string, error) {
	digestURL := HostURL + DigestMethod
	type digestResponse struct {
		Result int    `json:"result"`
		Digest string `json:"digest"`
	}
	resp, err := HTTPClient.Get(digestURL)
	if err != nil {
		return "", fmt.Errorf("getDigest http.Get: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("getDigest status code: %d", resp.StatusCode)
	}
	dr := &digestResponse{}
	if err := json.NewDecoder(resp.Body).Decode(dr); err != nil {
		return "", fmt.Errorf("getDigest json.Decode: %w", err)
	}
	return dr.Digest, nil
}

func makePasswordDigest(user, pass, digest string) string {
	// Definition of passworddigest from pCloud API docs:
	// passworddigest = sha1( password + sha1( lowercase of username ) + digest)
	return sha1Hash(pass + sha1Hash(strings.ToLower(user)) + digest)
}

// sha1Hash is a simple helper function for calculating the
// the password digest.
func sha1Hash(s string) string {
	data := []byte(s)
	return fmt.Sprintf("%x", sha1.Sum(data))
}
