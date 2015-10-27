package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	now := time.Now()
	tz, _ := time.LoadLocation("Europe/Paris")
	parisTime := now.In(tz)
	fmt.Printf("Local time: %s\nParis time: %s\n", now, parisTime)

	_, err := http.Get("https://golang.org/")
	if err == nil {
		fmt.Println("GoLang website is UP")
	} else {
		fmt.Printf("GoLang website is DOWN\nErr: %s\n", err.Error())
	}
}
