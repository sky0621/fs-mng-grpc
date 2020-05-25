package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/sky0621/fs-mng-grpc/samples/common"
)

func main() {
	get("/v1/review/facilities")
	get("/v1/review/records")
}

func get(path string) {
	resp, err := http.Get("http://localhost" + common.ReviewWebServerPort + path)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func() {
		if resp == nil {
			return
		}
		if err := resp.Body.Close(); err != nil {
			log.Fatalf("did not close: %v", err)
		}
	}()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("could not ReadAll: %v", err)
	}
	fmt.Println(string(body))
}
