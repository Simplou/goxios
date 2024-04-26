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
	validateValues := func (){
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
	if err := json.Unmarshal(&gabriel); err != nil{
		t.Fatal(err)
	}
	validateValues()
}
