// Copyright 2014 DB Medialab.  All rights reserved.
// License: MIT

// Package orbitapi provides client access to the Orbit API (http://orbitapi.com/)
package orbitapi

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

var (
	// URL on which the Orbit API can be reached
	orbitApiUrl = "http://api.orbitapi.com/"
)

type OrbitApi struct {
	// Key to access the API with
	apiKey string

	// Result will be sent on this channel
	Result chan map[string]interface{}
}

// Create a new Orbit API client
func NewClient(apiKey string) (orbitapi *OrbitApi) {
	orbitapi = new(OrbitApi)
	orbitapi.apiKey = apiKey
	orbitapi.Result = make(chan map[string]interface{})
	return
}

// Send a new GET request to the API
func (o *OrbitApi) Get(uri string) error {
	getUrl := orbitApiUrl + uri
	req, err := http.NewRequest("GET", getUrl, nil)
	if err != nil {
		return err
	}

	// Get requests require the API key to be sent as a header
	req.Header.Add("X-Orbit-API-Key", o.apiKey)
	return o.doRequest(req)
}

// Send a new POST request to the API
func (o *OrbitApi) Post(uri string, args url.Values) error {
	postUrl := orbitApiUrl + uri
	// Post requests require the API key to be sent as a key=value pair
	args.Add("api_key", o.apiKey)
	req, err := http.NewRequest("POST", postUrl, strings.NewReader(args.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return o.doRequest(req)
}

// Do the actual request and return the response on o.Result
func (o *OrbitApi) doRequest(req *http.Request) error {
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return err
	}
	o.Result <- data
	return nil
}
