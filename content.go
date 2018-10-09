package main

import "time"

type Content struct {
	Text      string    `datastore:"text"`
	CreatedAt time.Time `datastore:"createdAt"`
}
