package vine

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

type request struct {
	client *http.Client
}

var (
	//ErrUserDoesntExist is the error Vine returns when a user doesn't exist.
	ErrUserDoesntExist = "That record does not exist."
)

const (
	//VineAPIEndpoint is the default API endpoint to use.
	//Vine has multiple endpoints, this is the default endpoint used for
	//all web traffic.
	VineAPIEndpoint = "https://api.vineapp.com"
)

//NewRequest returns a new Request. If an http.Client is supplied
//as an argument, it will use that client for all requests. If no
//http.Client is specified, it will use the http.DefaultClient.
func NewRequest(c ...*http.Client) *request {
	r := new(request)
	if len(c) == 0 {
		r.client = http.DefaultClient
	} else {
		r.client = c[0]
	}
	return r
}

//Get sends an HTTP GET request to the specified endpoint.
//Initially, this was meant to be an unexported function. However,
//developments with github.com/austindizzy/davine required a more
//direct interaction with the Vine API.
//Future iterations will likely privatize this function again once
//I've had the chance to re-study the current Vine endpoints and
//response structure to build the rest of the functions to interact with
//said endpoints.
func (v *request) Get(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", VineAPIEndpoint+url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("x-vine-client", "vinewww/1.0")
	resp, err := v.client.Do(req)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(resp.Body)
}

//GetUser gets a user profile record. The supplied argument may be the
//user's numeric user ID (all numeric characters, no vanity-supported alpha
//characters), or the user's registered vanity (e.g. SpaceX, codyko, vine).
func (v *request) GetUser(userID string) (*User, error) {
	url := "/users/profiles/"

	if IsVanity(userID) {
		url += "vanity/" + userID
	} else {
		url += userID
	}

	resp, err := v.Get(url)
	if err != nil {
		return nil, err
	}
	data := new(userWrapper)
	err = json.Unmarshal(resp, &data)
	if !data.Success {
		return nil, errors.New(data.Error)
	}
	return data.Data, nil
}

func (v *request) GetPopularUsers(num int) ([]*PopularRecord, error) {
	uri := fmt.Sprintf("/timelines/popular?size=%d", num)
	resp, err := v.Get(uri)
	if err != nil {
		return nil, err
	}
	data := new(popularWrapper)
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	if !data.Success {
		return data.Data.Records, errors.New(data.Error)
	}
	return data.Data.Records, nil
}

//IsVanity simply checks if the supplied argument is a fully numeric
//user ID, or an alphanumeric vanity. True = vanity, false = numeric user ID.
func IsVanity(user string) bool {
	match, _ := regexp.MatchString("^[0-9]+$", user)
	return !match
}
