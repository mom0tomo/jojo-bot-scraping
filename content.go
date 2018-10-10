package main

import "time"

type Content struct {
	Title     string `datastore:"title"`
	Date      string `datastore:"date"`
	Text      string    `datastore:"text"`
	CreatedAt time.Time `datastore:"createdAt"`
}
