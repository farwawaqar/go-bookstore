package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//user data will be in json so we need to unmarshal it to be able to use it in controller

// func parsebody will help me in parsing specially for createBook function
// we receive somebody in json form in Request so i need to parse it. i'll unmarshal it.
func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
