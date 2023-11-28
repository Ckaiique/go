package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main()  {
	url :=  "https://br1.api.riotgames.com/lol/summoner/v4/summoners/by-name/reienma?api_key=RGAPI-257aa420-e921-45cb-9e84-9b34b086b1c5"
	resp,err :=  http.Get(url)

	if err != nil{
		fmt.Println(err)
		os.Exit(0)	
	}

	Body,erro :=  io.ReadAll(resp.Body)

	if erro != nil{
		fmt.Println(err)
		os.Exit(0)	
	}

	var m json.RawMessage
	errs := json.Unmarshal(Body, &m)
	
	if errs != nil{
		fmt.Println(err)
		os.Exit(0)	
	}


	fmt.Println(string(m) )
}