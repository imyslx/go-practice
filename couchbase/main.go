package main

import (
	"encoding/json"
	"fmt"

	"github.com/couchbase/gocb"
	zlog "github.com/rs/zerolog/log"
)

type Person struct {
	ID        string        `json:"id,omitempty"`
	Firstname string        `json:"firstname,omitempty"`
	Lastname  string        `json:"lastname,omitempty"`
	Social    []SocialMedia `json:"socialmedia.omitempty"`
}

type SocialMedia struct {
	Title string `json:"title"`
	Link  string `json:"link"`
}

func main() {
	cluster, err := gocb.Connect("couchbase://127.0.0.1")
	if err != nil {
		zlog.Fatal().Err(err).Msg("Cannot connect to couchbase.")
	}
	cluster.Authenticate(gocb.PasswordAuthenticator{
		Username: "user",
		Password: "pass",
	})
	bucket, err := cluster.OpenBucket("idp_database", "")
	if err != nil {
		zlog.Fatal().Err(err).Msg("Cannot connect to bucket.")
	}

	var person Person
	bucket.Upsert("imyslx", Person{
		Firstname: "imy",
		Lastname:  "slx",
		Social: []SocialMedia{
			{Title: "Twitter", Link: "https://www.twitter.com/imyslx"},
			{Title: "GitHub", Link: "https://www.github;com/imyslx"},
		},
	}, 0)

	bucket.Get("imyslx", &person)

	jsonBytes, _ := json.Marshal(person)
	fmt.Println(string(jsonBytes))
}
