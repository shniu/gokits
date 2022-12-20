package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	var batchTotal = 1
	interval := 500

	for batchTotal > 0 {

		for i := 0; i < 10; i++ {
			go getLocalSrv()
		}

		batchTotal--
		time.Sleep(time.Duration(interval) * time.Millisecond)
	}

	// select {}
	time.Sleep(time.Duration(20) * time.Second)
}

func getLocalSrv() {
	tatumSrvUrl := "http://127.0.0.1:7008/v1/tatum/kms/63634ad2ac8f94697cde206a"

	resp, err := http.Get(tatumSrvUrl)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.Status)
}
