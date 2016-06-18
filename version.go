package gorocket

import (
	"encoding/json"
	"net/http"
)

type Version struct {
	Api        string `json:"api"`
	RocketChat string `json:"rocketchat"`
}

type VersionResponse struct {
	Status  string `json:"status"`
	Version `json:"versions"`
}

func GetVersion(url string) (Version, error) {
	req, err := http.NewRequest("GET", url+"/version", nil)
	r, err := httpClient.Do(req)
	if r != nil {
		defer r.Body.Close()
	}
	if err != nil {
		return Version{}, err
	}
	version := new(VersionResponse)
	err = json.NewDecoder(r.Body).Decode(&version)
	if err != nil {
		return Version{}, err
	}
	return version.Version, nil
}
