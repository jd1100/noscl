package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/docopt/docopt-go"
	"github.com/fiatjaf/go-nostr/event"
)

func setMetadata(opts docopt.Opts) {
	initNostr()

	name, _ := opts.String("--name")
	description, _ := opts.String("--description")
	image, _ := opts.String("--image")

	jmetadata, _ := json.Marshal(struct {
		Name        string `json:"name,omitempty"`
		Description string `json:"image,omitempty"`
		Image       string `json:"image,omitempty"`
	}{name, description, image})

	event, statuses, err := pool.PublishEvent(&event.Event{
		PubKey:    getPubKey(config.PrivateKey),
		CreatedAt: uint32(time.Now().Unix()),
		Kind:      event.KindSetMetadata,
		Tags:      make(event.Tags, 0),
		Content:   string(jmetadata),
	})
	if err != nil {
		log.Printf("Error publishing: %s.\n", err.Error())
		return
	}

	printPublishStatus(event, statuses)
}
