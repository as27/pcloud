package pcloud

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// Param is a struct to hold a parameter for an API request
type Param struct {
	Name string
	Val  string
}

func makeURL(method string, params ...Param) string {
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
		q.Add(p.Name, p.Val)
	}
	u.RawQuery = q.Encode()
	return u.String()
}

// ApiRequest is a generic function to request data from the pCloud API
// It takes a pointer to a struct as response and a method name and
// a list of parameters. The parameters are added to the url as query
func ApiRequest(respStruct any, method string, params ...Param) error {
	u := makeURL(method, params...)
	resp, err := HTTPClient.Get(u)
	if err != nil {
		return fmt.Errorf("apiRequest http.Get: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("apiRequest status code: %d", resp.StatusCode)
	}
	if err := json.NewDecoder(resp.Body).Decode(respStruct); err != nil {
		return fmt.Errorf("apiRequest json.Decode: %w", err)
	}
	return nil
}
