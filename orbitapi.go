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
	Result chan map[string]interface{}
}

func NewClient(apiKey string) (orbitapi *OrbitApi) {
	orbitapi = new(OrbitApi)
	orbitapi.apiKey = apiKey
	orbitapi.Result = make(chan map[string]interface{})
	return
}

func (o *OrbitApi) Get(uri string) error {

	getUrl := orbitApiUrl + uri
	req, err := http.NewRequest("GET", getUrl, nil)
	if err != nil {
		return err
	}

	req.Header.Add("X-Orbit-API-Key", o.apiKey)
	return o.doRequest(req)
}

func (o *OrbitApi) Post(uri string, args url.Values) error {

	postUrl := orbitApiUrl + uri
	args.Add("api_key", o.apiKey)
	req, err := http.NewRequest("POST", postUrl, strings.NewReader(args.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return o.doRequest(req)
}

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
