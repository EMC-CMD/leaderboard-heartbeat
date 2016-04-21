package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

func heartbeat() {
	u, err := url.Parse(os.Getenv("LEADERBOARD_URL"))
	if err != nil {
		log.Fatal(err)
	}
	q := u.Query()
	q.Set("token", os.Getenv("LEADERBOARD_TOKEN"))
	u.RawQuery = q.Encode()

	for true {
		resp, err := http.Get(u.String())
		if err != nil {
			log.Printf("error in response: %+v", err)
		}
		log.Printf("resp: %+v", resp)

		time.Sleep(60 * time.Second)
	}
}

func main() {
	go heartbeat()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
