package goxios

import (
	"bytes"
	"testing"
)

func TestJSON(t *testing.T) {
	json := JSON{
		"name": "Gabriel Luiz",
		"job":  "Backend Developer",
	}
	buf, err := json.Marshal()
	if err != nil {
		t.Fatal(err)
	}
	var gabriel struct {
		Name string `json:"name"`
		Job  string `json:"job"`
	}
	validateValues := func() {
		if gabriel.Name != json["name"] && gabriel.Job != json["job"] {
			t.Fatal("Not decoded")
		}
	}
	if err := DecodeJSON(bytes.NewBuffer(buf), &gabriel); err != nil {
		t.Fatal(err)
	}
	validateValues()
	name := "gabrielluizsf"
	json["name"] = name
	if err := json.Unmarshal(&gabriel); err != nil {
		t.Fatal(err)
	}
	validateValues()
}

func TestGenericJSON(t *testing.T) {
	type Developer struct {
		Name string `json:"name"`
		Job  string `json:"job"`
	}
	json := GenericJSON[[]Developer]{
		"devs": []Developer{
			{Name: "Gabriel Luiz", Job: "Backend Developer"},
		},
	}
	buf, err := json.Marshal()
	if err != nil {
		t.Fatal(err)
	}
	var data struct {
		Devs []Developer `json:"devs"`
	}
	if err := DecodeJSON(bytes.NewBuffer(buf), &data); err != nil {
		t.Fatal(err)
	}
	validateValues := func() {
		if data.Devs[0].Name != json["devs"][0].Name {
			t.Fatal("Not Decoded")
		}
	}
	validateValues()
	name := "gabrielluizsf"
	json["devs"][0].Name = name
	if err := json.Unmarshal(&data); err != nil {
		t.Fatal(err)
	}
	validateValues()
}
