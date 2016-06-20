package gorocket

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Room struct {
	ID   string `json:"_id"`
	Name string `json:"name"`
	LM   string `json:"lm"`
	MSGS uint64 `json:"msgs"`
	T    string `json:"t"`
	TS   string `json:"ts"`
	U    struct {
		ID       string `json:"_id"`
		Username string `"json:"username"`
	} `json:"u"`
	Usernames []string `json:"usernames"`
}

type ListRoomsResponse struct {
	Status string `json:"status"`
	Rooms  []Room `json:"rooms"`
}

type CreateRoomResponse struct {
	Status  string `json:"status"`
	Channel Room   `json:"channel"`
}

func (c *Client) ListPublicRooms() ([]Room, error) {
	req, err := http.NewRequest("GET", c.Url+"/publicRooms", nil)
	req.Header.Add("X-Auth-Token", c.Token)
	req.Header.Add("X-User-Id", c.UserId)
	r, err := httpClient.Do(req)
	if r != nil {
		defer r.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	rooms := new(ListRoomsResponse)
	err = json.NewDecoder(r.Body).Decode(&rooms)
	if err != nil {
		return nil, err
	}
	return rooms.Rooms, nil
}

func (c *Client) CreateRoom(name string) error {
	body := "name=" + name
	req, err := http.NewRequest("GET", c.Url+"v1/channels.create", bytes.NewBuffer([]byte(body)))
	r, err := httpClient.Do(req)
	if r != nil {
		defer r.Body.Close()
	}
	if err != nil {
		return err
	}
	room := new(CreateRoomResponse)
	err = json.NewDecoder(r.Body).Decode(&room)
	if err != nil {
		return err
	}
	return nil

}

func (c *Client) GetRoomId(name string) (string, error) {
	rooms, err := c.ListPublicRooms()
	if err != nil {
		return "", err
	}
	for _, room := range rooms {
		if room.Name == name {
			return room.ID, nil
		}
	}
	return "", nil
}
