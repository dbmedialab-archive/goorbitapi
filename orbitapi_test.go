// Copyright 2014 DB Medialab.  All rights reserved.
// License: MIT

// Package orbitapi provides client access to the Orbit API (http://orbitapi.com/ - http://orbit.ai/documentation/introduction)
package orbitapi

import (
	"encoding/json"
	"os"
	"testing"
)

func TestOrbitTagDecode(t *testing.T) {
	td, err := os.Open("test_data/tag.json")
	if err != nil {
		t.Fatal("Unable to open tag test data file")
	}

	data := new(OrbitTag)
	err = json.NewDecoder(td).Decode(&data)
	if err != nil {
		t.Error("Failed decoding json data:", err)
	}

	testData := []struct {
		n string
		v interface{}
		e interface{}
	}{
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
