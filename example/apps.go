package main

import (
	httpreq "github.com/wahyuhadi/httpreq"
	"fmt"
	"io/ioutil"
)

func main(){
	// For example will request to https://www.googleapis.com/drive/v2/files/folderId/children/childId
	r, err:=httpreq.Req("https://www.googleapis.com/drive/v2/files/folderId/children/childId")
	if err!=nil {
		fmt.Println(err)
		return
	}

	defer r.Body.Close()
	responseData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	responseString := string(responseData)
	fmt.Println(responseString)

}