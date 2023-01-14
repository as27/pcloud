package pcloud

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type param struct {
	name string
	val  string
}

func makeURL(method string, params ...param) string {
	u, err := url.Parse(fmt.Sprintf("%s/%s", HostURL, method))
	if err != nil {
		// if the url is not valid the program should exit
		// with a panic. Because then no data can be requested
		// from the pCloud API
		panic("makeURL url.Parse:" + err.Error())
	}
	q := u.Query()
	if authToken != "" {
		q.Add("auth", authToken)
	}
	for _, p := range params {
		q.Add(p.name, p.val)
	}
	u.RawQuery = q.Encode()
	return u.String()
}

func apiRequest(response any, method string, params ...param) error {
	u := makeURL(method, params...)
	fmt.Println(u)
	resp, err := HTTPClient.Get(u)
	if err != nil {
		return fmt.Errorf("apiRequest http.Get: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("apiRequest status code: %d", resp.StatusCode)
	}
	if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
		return fmt.Errorf("apiRequest json.Decode: %w", err)
	}
	return nil
}
