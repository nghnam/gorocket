package gorocket

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type StatusResponse struct {
	Status string `json:"status"`
}

func (c *Client) SendMessage(roomId, msg string) error {
	body, err := json.Marshal(map[string]string{"msg": msg})
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", c.Url+"/rooms/"+roomId+"/send", bytes.NewBuffer(body))
	req.Header.Add("X-Auth-Token", c.Token)
	req.Header.Add("X-User-Id", c.UserId)
	r, err := httpClient.Do(req)
	if r != nil {
		defer r.Body.Close()
	}
	if err != nil {
		return err
	}
	res := new(StatusResponse)
	err = json.NewDecoder(r.Body).Decode(&res)
	if err != nil {
		return err
	}
	return nil

}
