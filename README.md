[![build status](https://travis-ci.org/Aircto/angelgo.svg?branch=master)](https://travis-ci.org/Aircto/angelgo.svg?branch=master)
# Simple Go-Lang Library for Angel.co Web API

### Install

```bash
$ go get https://github.com/aircto/angelgo/
```

## Simple Usage

```go
	package main
	
	import (
		"fmt"
		"github.com/aircto/angelgo/angelco"
	)
	
	func main() {
		api := New('client-id', 'client-secret')
		api.SetAccessToken('access_token')
		
		user, err: = api.Me()
		
		if err != nil {
			return
		}
		
		fmt.Println(user.Id, user.Email)
	}
```
