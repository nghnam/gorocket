package gorocket

import (
	"net/http"
	"time"
)

var httpClient = &http.Client{
	Timeout: time.Second * 10,
}

type Client struct {
	Url    string
	Token  string
	UserId string
}
