# Simple HTTP(s) Request


### example import package(s)
```go
package main

import (
	httpreq "github.com/wahyuhadi/httpreq"
	"fmt"
	"encoding/json"
)

```

### Example Response Interface for api

```go

// Example Response for url  https://www.googleapis.com/drive/v2/files/folderId/children/childId
type Responses struct {
	Error struct {
		Errors []struct {
			Domain       string `json:"domain"`
			Reason       string `json:"reason"`
			Message      string `json:"message"`
			LocationType string `json:"locationType"`
			Location     string `json:"location"`
		} `json:"errors"`
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
} 
```

### example main
```go
func main(){
	// For example will request to https://www.googleapis.com/drive/v2/files/folderId/children/childId
	r, err:=httpreq.Req("https://www.googleapis.com/drive/v2/files/folderId/children/childId")
	if err!=nil {
		fmt.Println(err)
		return
	}

	defer r.Body.Close()
	Res := new(Responses)
	json.NewDecoder(r.Body).Decode(Res) // Decode for Interface Response
	fmt.Println(Res.Error.Message) // Print message in response
}
```