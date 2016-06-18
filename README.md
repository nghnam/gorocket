#gorocket - RocketChat API in Golang


#Installation
```sh
go get -u github.com/nghnam/gorocket
```

#Usage
```sh
import "github.com/nghnam/gorocket"
```

Example can be seen in `bin/rocketchat.go`

#Environment variables

```sh
export RC_URL='http://127.0.0.1:3000/api'
export RC_USER='rocket'
export RC_PASS='rocket'
#export RC_TOKEN=''
#export RC_USERID=''
```

`RC_TOKEN` and `RC_USERID` are optional.


```sh
$ go build rocketchat.go
$ ./rocketchat --send --room-name general --message '```RocketChat API in Golang```'
```

#Author

Nguyen Hoang Nam <nghnam@outlook.com>

#License
This project is licensed under the terms of the MIT license.
