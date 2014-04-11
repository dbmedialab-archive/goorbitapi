package orbitapi

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

var (
	orbitApiUrl = "http://api.orbitapi.com/"
)

type OrbitApi struct {
	apiKey string
}

func NewClient(apiKey string) (orbitapi *OrbitApi) {
	orbitapi = new(OrbitApi)
	orbitapi.apiKey = apiKey
	return
}

func (o *OrbitApi) Get(uri string) (map[string]interface{}, error) {

	getUrl := orbitApiUrl + uri
	req, err := http.NewRequest("GET", getUrl, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("X-Orbit-API-Key", o.apiKey)
	return doRequest(req)
}

func (o *OrbitApi) Post(uri string, args url.Values) (map[string]interface{}, error) {

	postUrl := orbitApiUrl + uri
	args.Add("api_key", o.apiKey)
	req, err := http.NewRequest("POST", postUrl, strings.NewReader(args.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return doRequest(req)
}

func doRequest(req *http.Request) (map[string]interface{}, error) {
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
