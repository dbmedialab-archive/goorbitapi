// Copyright 2014 DB Medialab.  All rights reserved.
// License: MIT

// Package orbitapi provides client access to the Orbit API (http://orbitapi.com/ - http://orbit.ai/documentation/introduction)
package orbitapi

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type TestData struct {
	n string
	v interface{}
	e interface{}
}

func TestConceptTagRequest(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		o, err := ioutil.ReadFile("test_data/tag.json")
		if err != nil {
			t.Fatal("Unable to open info test data file")
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, string(o))
	}))
	defer ts.Close()

	orbitApiUrl = ts.URL + "/"
	api := NewClient("apiKey")
	go func() {
		args := &url.Values{}
		args.Add("text", "Jeg liker politikk sa Solberg til Dagbladet.")
		if err := api.ConceptTag(args); err != nil {
			t.Fatal(err)
		}
	}()

	d := <-api.Result
	data := d.(*OrbitTag)
	testData := []TestData{
		{"Entities[\"Erna_Solberg\"].Image", data.Entities["Erna_Solberg"].Image, "http://upload.wikimedia.org/wikipedia/commons/2/25/Erna_Solberg,_Wesenberg,_2011_(1).jpg"},
		{"Entities[\"Erna_Solberg\"].Label", data.Entities["Erna_Solberg"].Label, "Erna Solberg"},
		{"Entities[\"Erna_Solberg\"].Link", data.Entities["Erna_Solberg"].Link, "http://no.wikipedia.org/wiki/Erna_Solberg"},
		{"Entities[\"Erna_Solberg\"].Relevance", data.Entities["Erna_Solberg"].Relevance, 0.5},
		{"Entities[\"Erna_Solberg\"].Thumbnail", data.Entities["Erna_Solberg"].Thumbnail, "http://upload.wikimedia.org/wikipedia/commons/thumb/2/25/Erna_Solberg,_Wesenberg,_2011_(1).jpg/50px-Erna_Solberg,_Wesenberg,_2011_(1).jpg"},
		{"Entities[\"Erna_Solberg\"].Type", data.Entities["Erna_Solberg"].Type, "Per"},
		{"Entities[\"Erna_Solberg\"].Image", data.Entities["Dagbladet"].Image, "http://upload.wikimedia.org/wikipedia/commons/0/05/Dagbladet_logo.svg"},
		{"Entities[\"Dagbladet\"].Label", data.Entities["Dagbladet"].Label, "Dagbladet"},
		{"Entities[\"Dagbladet\"].Link", data.Entities["Dagbladet"].Link, "http://no.wikipedia.org/wiki/Dagbladet"},
		{"Entities[\"Dagbladet\"].Relevance", data.Entities["Dagbladet"].Relevance, 0.5},
		{"Entities[\"Dagbladet\"].Thumbnail", data.Entities["Dagbladet"].Thumbnail, "http://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Dagbladet_logo.svg/50px-Dagbladet_logo.svg.png"},
		{"Entities[\"Dagbladet\"].Type", data.Entities["Dagbladet"].Type, "Org"},
		{"RemainingWords", data.RemainingWords, 9944},
		{"Text[0].([]interface{})[0]", data.Text[0].([]interface{})[0], "Jeg liker politikk sa"},
		{"Text[0].([]interface{})[1].(map[string]interface{})[\"entity\"]", data.Text[0].([]interface{})[1].(map[string]interface{})["entity"], "Erna_Solberg"},
		{"Text[0].([]interface{})[1].(map[string]interface{})[\"text\"]", data.Text[0].([]interface{})[1].(map[string]interface{})["text"], "Solberg"},
		{"Text[0].([]interface{})[2]", data.Text[0].([]interface{})[2], "til"},
		{"Text[0].([]interface{})[3].(map[string]interface{})[\"entity\"]", data.Text[0].([]interface{})[3].(map[string]interface{})["entity"], "Dagbladet"},
		{"Text[0].([]interface{})[3].(map[string]interface{})[\"text\"]", data.Text[0].([]interface{})[3].(map[string]interface{})["text"], "Dagbladet"},
	}

	for _, d := range testData {
		if d.v != d.e {
			t.Errorf("%v: expecting %v, got %v", d.n, d.e, d.v)
		}
	}
}

func TestAccountInfoRequest(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		o, err := ioutil.ReadFile("test_data/info.json")
		if err != nil {
			t.Fatal("Unable to open info test data file")
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, string(o))
	}))
	defer ts.Close()

	orbitApiUrl = ts.URL + "/"
	api := NewClient("apiKey")
	go func() {
		if err := api.AccountInfo(); err != nil {
			t.Fatal(err)
		}
	}()

	d := <-api.Result
	data := d.(map[string]interface{})
	testData := []TestData{
		{"iptc_requests_limit", data["iptc_requests_limit"], float64(150)},
		{"name", data["name"], "JustAdam"},
		{"concurrent_requests", data["concurrent_requests"], nil},
		{"iptc_requests_today", data["iptc_requests_today"], float64(0)},
		{"daily_word_limit", data["daily_word_limit"], float64(10000)},
		{"langdetect_requests_limit", data["langdetect_requests_limit"], float64(150)},
		{"langdetect_requests_today", data["langdetect_requests_today"], float64(0)},
		{"words_today", data["words_today"], float64(8)},
		{"api_key", data["api_key"], "XXX"},
		{"id", data["id"], float64(1)},
	}

	for _, d := range testData {
		if d.v != d.e {
			t.Errorf("%v: expecting %v, got %v", d.n, d.e, d.v)
		}
	}
}

func TestGetRequestAPIKeyIsSet(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["X-Orbit-Api-Key"][0] != "apiKey" {
			t.Errorf("Expecting 'apiKey', got %v", r.Header["X-Orbit-Api-Key"][0])
		}
	}))
	defer ts.Close()

	orbitApiUrl = ts.URL
	api := NewClient("apiKey")
	api.Get("")
}

func TestPostRequestContentType(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Content-Type"][0] != "application/x-www-form-urlencoded" {
			t.Errorf("Expecting 'application/x-www-form-urlencoded', got %v", r.Header["Content-Type"][0])
		}
	}))
	defer ts.Close()

	orbitApiUrl = ts.URL
	api := NewClient("apiKey")
	api.Post("", &url.Values{})
}

func TestPostRequestAPIKeyIsSet(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := ioutil.ReadAll(r.Body)
		v, _ := url.ParseQuery(string(body))
		if v["api_key"][0] != "apiKey" {
			t.Errorf("Expecting 'apiKey', got %v", v["api_key"][0])
		}
	}))
	defer ts.Close()

	orbitApiUrl = ts.URL
	api := NewClient("apiKey")
	api.Post("", &url.Values{})
}
