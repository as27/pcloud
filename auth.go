package pcloud

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func LogIn(user, pass string) error {
	token, err := getAuthToken(user, pass)
	if err != nil {
		return fmt.Errorf("LogIn getAuthToken: %w", err)
	}
	authToken = token
	return nil
}

func getAuthToken(user, pass string) (string, error) {
	type authResponse struct {
		Result int    `json:"result"`
		Token  string `json:"auth"`
	}

	u, err := url.Parse(HostURL + "/userinfo")
	if err != nil {
		return "", fmt.Errorf("getAuthToken url.Parse: %w", err)
	}
	q := u.Query()
	q.Add("getauth", "1")
	q.Add("username", user)
	q.Add("logout", "1")

	// just the url changes, when AuthMethod changes
	switch AuthMethod {
	case AuthMethodPassword:
		q.Add("password", pass)
	case AuthMethodDigest:
		digest, err := getDigest()
		if err != nil {
			return "", fmt.Errorf("getAuthToken getDigest: %w", err)
		}
		passwordDigest := makePasswordDigest(user, pass, digest)
		q.Add("passworddigest", passwordDigest)
		q.Add("digest", digest)
	}

	u.RawQuery = q.Encode()
	resp, err := HTTPClient.Get(u.String())
	if err != nil {
		return "", fmt.Errorf("getAuthToken http.Get: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("getAuthToken status code: %d\n%s", resp.StatusCode, u.String())
	}
	ar := &authResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ar); err != nil {
		return "", fmt.Errorf("getAuthToken json.Decode: %w", err)
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
