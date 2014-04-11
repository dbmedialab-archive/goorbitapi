package orbitapi

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strings"
)

var (
	infoURL = "http://api.orbitapi.com/info"
	tagURL  = "http://api.orbitapi.com/tag"
)

type OrbitApi struct {
	apiKey string
}

func NewClient(apiKey string) (orbitapi *OrbitApi) {
	orbitapi = new(OrbitApi)
	orbitapi.apiKey = apiKey
	return
}

func (o *OrbitApi) Info() (map[string]interface{}, error) {

	req, err := http.NewRequest("GET", infoURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("X-Orbit-API-Key", o.apiKey)
	return doRequest(req)
}

func (o *OrbitApi) Tag(args url.Values) (map[string]interface{}, error) {

	// APIs required arguments
	_, text := args["text"]
	_, url := args["url"]
	if text == true && url == true || text == false && url == false {
		return nil, errors.New("You must specify either text or url")
	}

	args.Add("api_key", o.apiKey)
	req, err := http.NewRequest("POST", tagURL, strings.NewReader(args.Encode()))
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
