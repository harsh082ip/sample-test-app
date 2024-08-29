package main

import (
	"log"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {

	for i := 1; i <= 15000; i++ {
		wg.Add(1)
		go func() {
			resp, err := http.Get("http://13.233.247.244:8003/data")
			if err != nil {
				log.Println(i, err.Error())
				wg.Done()
				return
			}

			if resp.StatusCode != 200 {
				log.Println(i, resp.StatusCode, resp.Body)
				wg.Done()
				return
			}

			log.Println(i, resp.StatusCode, resp.Body)
			wg.Done()
		}()
	}

	wg.Wait()
}
