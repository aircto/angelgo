[![build status](https://travis-ci.org/Aircto/angelgo.svg?branch=master)](https://travis-ci.org/Aircto/angelgo.svg?branch=master)

# Simple Go-Lang Library for Angel.co Web API

## Install

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
		api := angelco.New('client-id', 'client-secret')
		api.SetAccessToken('access_token')
		
		user, err: = api.Me()
		
		if err != nil {
			return
		}
		
		fmt.Println(user.Id, user.Email)
	}
```

## Getting Access Token

```go
	package main
	
	import (
		"fmt"
		"github.com/aircto/angelgo/angelco"
	)
	
	
	func main() {
		api := angelco.New('client-id', 'client-secret')
		url := api.AuthUrl('email') # scope email
		
		// take the user to "url" for authentication and it return back to
		// 'callback url' you have registered along with "code"
		
		token, err := api.GetAccessToken(code)
		if err != nil {
			return
		}
		fmt.Println(token.key, token.type) // token.key is actual access_token we will be using for every request
		
	}
```

## API Reference

``` go
	Me()(*UserResponse, error)
    User(userId int64) (*UserResponse, error)
    UserStartupRoles(userId int64) (*StartupRolesReponse, error)

    MyStatusList() (*StatusUpdatesResponse, error)
    UserStatusList(userId int64) (*StatusUpdatesResponse, error)

    // Will give a list of status messages if you are team member of the startup else it will be emptyn
    StartupStatusList(startupId int64) (*StatusUpdatesResponse, error)
    PostMyStatus(message string) (*StatusUpdateResponse, error)
    PostStartupStatus(startupId int64, message string) error
    RemoveStatus(statusId int64) error

    // Jobs api
    JobsList() (*JobsListResponse, error)
    Job(jobId int64) (*JobResponse, error)
    JobsOfStartup(startupId int64) ([]Job , error)
    JobsOfTag(tagId int64) (*JobsListResponse, error)

    //Startups api
    Startup(startupId int64) (*StartupResponse, error)
	
```

## License

MIT License. See LICENSE file
