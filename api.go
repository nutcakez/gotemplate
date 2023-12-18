package main

import (
	"fmt"
	"io"
	"net/http"
)

func GetAvailableWorldBosses() {
	url := "https://api.guildwars2.com/v2/account/worldbosses"
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+apikey)
	res, _ := client.Do(req)
	body, _ := io.ReadAll(res.Body)
	fmt.Println(res.StatusCode)
	fmt.Println(string(body))
}
