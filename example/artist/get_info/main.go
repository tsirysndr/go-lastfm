package main

import (
	"encoding/json"
	"fmt"

	"github.com/tsirynsdr/go-lastfm"
)

func main() {
	client := lastfm.NewClient("975f432af9cbae122695bdd768f40ae8")
	res, _ := client.Artist.GetInfo("Travis Scott")
	r, _ := json.Marshal(res)
	fmt.Println(string(r))
}
