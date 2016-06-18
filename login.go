package gorocket

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Credentical struct {
	Token  string `json:"authToken"`
	UserId string `json:"userId"`
}

type LoginResponse struct {
	Status string `json:"status"`
	Data   struct {
		Token  string `json:"authToken"`
		UserId string `json:"userId"`
	} `json:"data"`
}

func GetToken(url, username, password string) (Credentical, error) {
	body := "user=" + username + "&password=" + password
	req, err := http.NewRequest("POST", url+"/login", bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r, err := httpClient.Do(req)
	if r != nil {
		defer r.Body.Close()
	}
	if err != nil {
		return Credentical{}, err
	}
	cred := new(LoginResponse)
	err = json.NewDecoder(r.Body).Decode(&cred)
	if err != nil {
		return Credentical{}, err
	}
	return cred.Data, nil
}
