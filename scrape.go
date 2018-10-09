package main

import (
	"io/ioutil"
	"net/http"
	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/urlfetch"
)

func scrape(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	client := urlfetch.Client(ctx)

	resp, err := client.Get("https://www.google.com/")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	text := string(byteArray[1])

	cnt := &Content{
		Text:      text,
		CreatedAt: time.Now(),
	}

	const k = "Content"
	key := datastore.NewIncompleteKey(ctx, k, nil)
	if _, err := datastore.Put(ctx, key, cnt); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
